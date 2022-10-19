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

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.URI;
import java.net.URISyntaxException;
import java.nio.charset.StandardCharsets;
import java.util.ArrayList;
import java.util.List;
import java.util.logging.Logger;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import org.apache.http.HttpResponse;
import org.apache.http.client.fluent.Request;
import org.apache.http.client.utils.URIBuilder;

import io.openliberty.arquillian.managed.exceptions.NestedExceptionBuilder.ExMsg;

/**
 * Tries to receive information about an exception from the server in text form
 * <p>
 * This relies on the liberty-support-feature being installed and running
 * <p>
 * The text format expected from the server includes information about the
 * exception and its cause chain.
 * <p>
 * For each exception in the cause chain, the following information is returned:
 * <ul>
 * <li>the class name</li>
 * <li>the names of all superclasses (if any)</li>
 * <li>the exception message</li>
 * </ul>
 * <p>
 * Example:
 * 
 * <pre>
 * <code>
 * exClass com.example.Exception
 * exSuperclass com.example.SuperclassOfException
 * exSuperclass com.example.SuperclassOfSuperclassOfException
 * This is the exception message
 * which can have multiple lines
 * exClass com.example.CauseOfException
 * Exception cause message
 * ...
 * </code>
 * </pre>
 * <p>
 * This strategy for retrieving exceptions does not retain as much information
 * as retrieving a serialized object (in particular the stack trace is lost) but
 * it's much more robust against missing classes on the client. If one class in
 * the cause chain can't be loaded, other classes in the chain can still be
 * loaded. In addition, if a class in can't be loaded, we also attempt to load
 * classes in its type hierarchy instead. This should catch most cases where a
 * test requires that a spec exception is thrown but liberty throws an exception
 * which subclasses it. The subclass may not be available to the client but the
 * spec exception must be (since it was referenced from the test).
 * 
 */
public class SupportFeatureTextExceptionLocator implements DeploymentExceptionLocator {

    private final static Logger log = Logger.getLogger(SupportFeatureTextExceptionLocator.class.getName());
    
    private final URI uri;
    
    private static final Pattern CLASS_PATTERN = Pattern.compile("exClass (.*)");
    private static final Pattern SUPERCLASS_PATTERN = Pattern.compile("exSuperclass (.*)");
    
    public SupportFeatureTextExceptionLocator(String host, int port) {
        try {
            uri = new URI("http", null, host, port, "/arquillian-support/deployment-exception", "format=text", null);
        } catch (URISyntaxException e) {
            // Shouldn't happen as most of the URI parts are hard coded
            throw new IllegalArgumentException("Invalid URI: " + e, e);
        }
    }

    
    @Override
    public Throwable getException(String appName, String logLine, long deploymentTime) {
        try {
            URIBuilder uriBuilder = new URIBuilder(uri);
            uriBuilder.addParameter("appName", appName);
            
            HttpResponse resp = Request.Get(uriBuilder.build()).execute().returnResponse();
            if (resp.getStatusLine().getStatusCode() == 400) {
                log.warning("After " + appName + " failed to start, the server did not report an exception for that app");
            } else if (resp.getStatusLine().getStatusCode() != 200) {
                log.info("Unable to recieve text format exception from server, is usr:arquillian-support-1.0?");
            } else {
                log.finer("Reading exception returned from server");
                try (BufferedReader reader = new BufferedReader(new InputStreamReader(resp.getEntity().getContent(), StandardCharsets.UTF_8))) {
                    return readResponse(reader);
                }
            }
        } catch (IOException ex) {
            log.warning("IO Exception trying to get text exception: " + ex);
        } catch (Exception ex) {
            log.warning("Unexpected exception thrown while trying to get text exception: " + ex);
        }
        
        return null;
    }
    
    private Throwable readResponse(BufferedReader reader) throws IOException {
        String line;
        ResponseReader responseReader = new ResponseReader();
        while ((line = reader.readLine()) != null) {
            Matcher classMatcher = CLASS_PATTERN.matcher(line);
            if (classMatcher.matches()) {
                responseReader.readClass(classMatcher.group(1));
                continue;
            }
            
            Matcher superclassMatcher = SUPERCLASS_PATTERN.matcher(line);
            if (superclassMatcher.matches()) {
                responseReader.readSuperclass(superclassMatcher.group(1));
                continue;
            }
            
            responseReader.readMessage(line);
        }
        responseReader.finishMsg();
        List<ExMsg> msgs = responseReader.finish();
        
        return NestedExceptionBuilder.buildNestedException(msgs);
    }
    
    private static class ResponseReader {
        List<ExMsg> msgs = new ArrayList<>();
        ExMsg currentMsg = null;
        StringBuilder messageBuilder = null;
        
        private void readClass(String className) {
            log.finer("Read class " + className);
            finishMsg();
            currentMsg.exName = className;
        }
        
        private void readSuperclass(String className) {
            log.finer("Read superclass" + className);
            currentMsg.superclasses.add(className);
        }
        
        private void readMessage(String message) {
            log.finer("Read message " + message);
            if (messageBuilder == null) {
                messageBuilder = new StringBuilder();
            } else {
                messageBuilder.append("\n");
            }
            messageBuilder.append(message);
        }
        
        private void finishMsg() {
            if (currentMsg != null) {
                msgs.add(currentMsg);
                if (messageBuilder != null) {
                    currentMsg.logMsg = messageBuilder.toString();
                }
            }
            currentMsg = new ExMsg();
            messageBuilder = null;
        }
        
        private List<ExMsg> finish() {
            return msgs;
        }
    }

}
