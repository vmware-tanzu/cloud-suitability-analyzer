/*
 * Copyright 2018, IBM Corporation, Red Hat Middleware LLC, and individual contributors
 * identified by the Git commit log. 
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package io.openliberty.arquillian.managed.exceptions;

import java.io.IOException;
import java.io.InputStream;
import java.io.ObjectInputStream;
import java.net.URI;
import java.net.URISyntaxException;
import java.util.logging.Logger;

import org.apache.http.HttpResponse;
import org.apache.http.client.fluent.Request;
import org.apache.http.client.utils.URIBuilder;

/**
 * Tries to receive a serialized exception from the server
 * <p>
 * This relies on the liberty-support-feature being installed and running
 * <p>
 * This is the best way to get the exception as it retains all of the exception
 * information from the server, including the stack trace, but will fail if any
 * of the serialized classes can't be loaded on the client
 */
public class SupportFeatureSerializedExceptionLocator implements DeploymentExceptionLocator {
    
    private final static Logger log = Logger.getLogger(SupportFeatureSerializedExceptionLocator.class.getName());
    
    private final URI uri;
    
    public SupportFeatureSerializedExceptionLocator(String host, int port) {
        try {
            uri = new URI("http", null, host, port, "/arquillian-support/deployment-exception", "format=serialize", null);
        } catch (URISyntaxException e) {
            // Shouldn't happen as most of the URI parts are hard coded
            throw new IllegalArgumentException("Invalid URI: " + e, e);
        }
    }

    @Override
    public Throwable getException(String appName, String logLine, long deploymentTime) {
        Throwable result = null;
        try {
            URIBuilder uriBuilder = new URIBuilder(uri);
            uriBuilder.addParameter("appName", appName);
            
            HttpResponse resp = Request.Get(uriBuilder.build()).execute().returnResponse();
            if (resp.getStatusLine().getStatusCode() == 400) {
                log.warning("After " + appName + " failed to start, the server did not report an exception for that app");
            } else if (resp.getStatusLine().getStatusCode() != 200) {
                log.info("Unable to recieve serialized exception from server, is usr:arquillian-support-1.0 installed?");
            } else {
                try (InputStream inStream = resp.getEntity().getContent()) {
                    ObjectInputStream objStream = new ObjectInputStream(inStream);
                    Object readObject = objStream.readObject();
                    if (readObject instanceof Throwable) {
                        result = (Throwable) readObject;
                    }
                }
            }
        } catch (IOException ex) {
            log.warning("IO Exception trying to get serialized exception: " + ex);
        } catch (ClassNotFoundException ex) {
            log.finer("Unable to find class for serialized exception: " + ex);
        } catch (Exception ex) {
            log.warning("Unexpected exception while trying to recieve serialized exception: " + ex);
        }
        
        return result;
    }

}
