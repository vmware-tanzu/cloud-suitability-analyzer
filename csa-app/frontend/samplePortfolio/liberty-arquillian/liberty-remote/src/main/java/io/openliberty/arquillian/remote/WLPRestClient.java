/*
 * Copyright 2014 Red Hat Inc. and/or its affiliates and other contributors
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
package io.openliberty.arquillian.remote;

import java.io.File;
import java.io.IOException;
import java.net.URLEncoder;
import java.util.Map;
import java.util.logging.Level;
import java.util.logging.Logger;

import org.apache.http.HttpHost;
import org.apache.http.HttpResponse;
import org.apache.http.HttpStatus;
import org.apache.http.HttpVersion;
import org.apache.http.client.ClientProtocolException;
import org.apache.http.client.fluent.Executor;
import org.apache.http.client.fluent.Request;
import org.apache.http.entity.ContentType;
import org.jboss.arquillian.container.spi.client.container.DeploymentException;

import com.fasterxml.jackson.core.JsonParseException;
import com.fasterxml.jackson.databind.JsonMappingException;
import com.fasterxml.jackson.databind.ObjectMapper;

/**
 * 
 * This is a wrapper around the WebSphere Liberty JMX Rest API.
 * 
 * @author <a href="mailto:tayres@gmail.com">Tony Ayres</a>
 *
 */
public class WLPRestClient {

    private static final String className = WLPRestClient.class.getName();

    private static Logger log = Logger.getLogger(className);

    private WLPRemoteContainerConfiguration configuration;

    private static final String IBMJMX_CONNECTOR_REST = "/IBMJMXConnectorREST";
    private static final String FILE_ENDPOINT = IBMJMX_CONNECTOR_REST + "/file/";
    private static final String MBEANS_ENDPOINT = IBMJMX_CONNECTOR_REST + "/mbeans/";
    private static final String UTF_8 = "UTF-8";
    private static final String STARTED = "STARTED";

    private final Executor executor;

    public WLPRestClient(WLPRemoteContainerConfiguration configuration) {
        this.configuration = configuration;
        executor = Executor.newInstance().auth(new HttpHost(configuration.getHostName()), configuration.getUsername(),
                configuration.getPassword());
    }

    /**
     * Uses the rest api to upload an application binary to the dropins folder
     * of WLP to allow the server automatically deploy it.
     * 
     * @param archive
     * @throws ClientProtocolException
     * @throws IOException
     */
    public void deploy(File archive) throws ClientProtocolException, IOException {

        if (log.isLoggable(Level.FINER)) {
            log.entering(className, "deploy");
        }

        String deployPath = String.format("${wlp.user.dir}/servers/%s/dropins/%s", configuration.getServerName(),
                archive.getName());

        String serverRestEndpoint = String.format("https://%s:%d%s%s", configuration.getHostName(),
                configuration.getHttpsPort(), FILE_ENDPOINT, URLEncoder.encode(deployPath, UTF_8));

        HttpResponse result = executor.execute(
                Request.Post(serverRestEndpoint).useExpectContinue().version(HttpVersion.HTTP_1_1)
                        .bodyFile(archive, ContentType.DEFAULT_BINARY)).returnResponse();

        if (log.isLoggable(Level.FINE)) {
            log.fine("While deploying file " + archive.getName() + ", server returned response: "
                    + result.getStatusLine().getStatusCode());
        }

        if (!isSuccessful(result)) {
            throw new ClientProtocolException("Could not deploy application to server, server returned response: " + result);
        }

        if (log.isLoggable(Level.FINER)) {
            log.exiting(className, "deploy");
        }

    }

    /**
     * Deletes the specified application from the servers dropins directory. WLP
     * will detect this and then undeploy it.
     * 
     * @param String
     *            - applicationName
     * @throws ClientProtocolException
     * @throws IOException
     */
    public void undeploy(String applicationName) throws ClientProtocolException, IOException {

        if (log.isLoggable(Level.FINER)) {
            log.entering(className, "undeploy");
        }

        String deployPath = String.format("${wlp.user.dir}/servers/%s/dropins/%s", configuration.getServerName(),
                applicationName);

        String serverRestEndpoint = String.format("https://%s:%d%s%s", configuration.getHostName(),
                configuration.getHttpsPort(), FILE_ENDPOINT, URLEncoder.encode(deployPath, UTF_8));

        HttpResponse result = executor.execute(
                Request.Delete(serverRestEndpoint).useExpectContinue().version(HttpVersion.HTTP_1_1)).returnResponse();

        if (isSuccessful(result)) {
            log.fine("File " + applicationName + " was deleted");
        } else {
            throw new ClientProtocolException("Unable to undeploy application " + applicationName
                    + ", server returned response: " + result.getStatusLine());
        }

        if (log.isLoggable(Level.FINER)) {
            log.exiting(className, "undeploy", result);
        }
    }

    /**
     * Calls the rest api to determine if the application server is up and
     * running.
     * 
     * @return boolean - true if the server is running
     */
    public boolean isServerUp() throws ClientProtocolException, IOException {
        if (log.isLoggable(Level.FINER)) {
            log.entering(className, "isServerUp");
        }

        String hostName = String.format("https://%s:%d%s", configuration.getHostName(), configuration.getHttpsPort(),
                IBMJMX_CONNECTOR_REST);

        HttpResponse result = executor.execute(Request.Get(hostName)).returnResponse();
        
        if (!isSuccessful(result)) {
            throw new ClientProtocolException("Could not successfully connect to REST endpoint, server returned response: " + result);
        }

        if (log.isLoggable(Level.FINER)) {
            log.exiting(className, "isServerUp");
        }
        return isSuccessful(result);
    }

    /**
     * Queries the rest api for the servers name.
     * 
     * @return the name of the running server e.g. defaultServer
     * @throws IOException
     * @throws ClientProtocolException
     */
    public String getServerName() throws ClientProtocolException, IOException {
        if (log.isLoggable(Level.FINER)) {
            log.entering(className, "getServerName");
        }

        String restEndpoint = String.format("https://%s:%d%sWebSphere:feature=kernel,name=ServerInfo/attributes/Name",
                configuration.getHostName(), configuration.getHttpsPort(), MBEANS_ENDPOINT);

        String jsonResponse = executor.execute(Request.Get(restEndpoint)).returnContent().asString();
        String serverName = parseJsonResponse(jsonResponse);

        if (log.isLoggable(Level.FINER)) {
            log.exiting(className, "isServerUp");
        }
        return serverName;
    }

    /**
     * Checks the application State using the WLP rest api.
     * 
     * @param applicationName
     * @return true if the application is in STARTED state
     * @throws DeploymentException
     */
    public boolean isApplicationStarted(String applicationName) {

        if (log.isLoggable(Level.FINER)) {
            log.entering(className, "isApplicationStarted");
        }

        String restEndpoint = String.format(
                "https://%s:%d%sWebSphere:service=com.ibm.websphere.application.ApplicationMBean,name=%s/attributes/State",
                configuration.getHostName(), configuration.getHttpsPort(), MBEANS_ENDPOINT, applicationName);

        String status = "";
        try {
            String jsonResponse = executor.execute(Request.Get(restEndpoint)).returnContent().asString();
            status = parseJsonResponse(jsonResponse);
        } catch (ClientProtocolException e) {
            // This exception is expected if the application hasn't been
            // deployed yet as its MBean won't exist.
            // We expect this and can continue, set status to error.
            log.finest("Expected error occurred while checking if application " + applicationName
                    + " is already started, app may not have been deployed yet. Ok to continue. " + e);
            status = "error";
        } catch (IOException e) {
            log.severe("IOException occurred while checking if application " + applicationName + " is already started " + e);
            status = "error";
        }

        boolean applicationState;

        if (STARTED.equals(status)) {
            applicationState = true;
        } else {
            applicationState = false;
        }

        if (log.isLoggable(Level.FINER)) {
            log.exiting(className, "isApplicationStarted", applicationState);
        }

        return applicationState;
    }

    /*
     * Get a value from a json response.
     */
    private String parseJsonResponse(String jsonString) {
        ObjectMapper mapper = new ObjectMapper();
        String value = "";
        try {
            Map result = mapper.readValue(jsonString.getBytes(), Map.class);
            value = (String) result.get("value");
        } catch (JsonParseException e) {
            log.severe("Error parsing Json response " + e);
        } catch (JsonMappingException e) {
            log.severe("Error mapping Json response " + e);
        } catch (IOException e) {
            log.severe("IOException while parsing Json response " + e);
        }
        return value;
    }

    /**
     * Tests if a HttpResponse contains a 2xx response code
     * 
     * @param response
     * @return true if the response code is a 2xx code.
     */
    private boolean isSuccessful(HttpResponse response) {
        if (response.getStatusLine().getStatusCode() >= HttpStatus.SC_OK
                && response.getStatusLine().getStatusCode() <= HttpStatus.SC_NO_CONTENT) {
            return true;
        }
        return false;
    }

}
