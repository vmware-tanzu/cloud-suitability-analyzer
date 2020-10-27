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
import java.util.logging.Level;
import java.util.logging.Logger;

import org.apache.http.client.ClientProtocolException;
import org.jboss.arquillian.container.spi.client.container.DeployableContainer;
import org.jboss.arquillian.container.spi.client.container.DeploymentException;
import org.jboss.arquillian.container.spi.client.container.LifecycleException;
import org.jboss.arquillian.container.spi.client.protocol.ProtocolDescription;
import org.jboss.arquillian.container.spi.client.protocol.metadata.HTTPContext;
import org.jboss.arquillian.container.spi.client.protocol.metadata.ProtocolMetaData;
import org.jboss.arquillian.container.spi.client.protocol.metadata.Servlet;
import org.jboss.shrinkwrap.api.Archive;
import org.jboss.shrinkwrap.api.exporter.ZipExporter;
import org.jboss.shrinkwrap.descriptor.api.Descriptor;

public class WLPRemoteContainer implements DeployableContainer<WLPRemoteContainerConfiguration> {

    private static final String className = WLPRemoteContainer.class.getName();

    private static Logger log = Logger.getLogger(className);

    private WLPRemoteContainerConfiguration containerConfiguration;

    private WLPRestClient restClient;

    @Override
    public Class<WLPRemoteContainerConfiguration> getConfigurationClass() {
        return WLPRemoteContainerConfiguration.class;
    }

    @Override
    public void setup(WLPRemoteContainerConfiguration configuration) {
        if (log.isLoggable(Level.FINER)) {
            log.entering(className, "setup");
        }

        this.containerConfiguration = configuration;
        restClient = new WLPRestClient(containerConfiguration);

        if (log.isLoggable(Level.FINER)) {
            log.exiting(className, "setup");
        }
    }

    @Override
    public void start() throws LifecycleException {

        boolean ready;

        try {
            ready = restClient.isServerUp();

            if (!ready) {
                throw new LifecycleException("Remote server is not started");
            }

            String arqServerName = containerConfiguration.getServerName();
            String remoteServerName = restClient.getServerName();

            if (!arqServerName.equals(remoteServerName)) {
                throw new LifecycleException("The serverName (" + arqServerName
                        + ") specified in arquillian.xml does not match the server name of the remote server ("
                        + remoteServerName + ").");
            }

        } catch (ClientProtocolException e) {
            throw new LifecycleException("Could not determine remote server status : " + e.getMessage(), e);
        } catch (IOException e) {
            throw new LifecycleException("Could not determine remote server status : " + e.getMessage(), e);
        }
    }

    @Override
    public void stop() throws LifecycleException {
        if (log.isLoggable(Level.FINER)) {
            log.entering(className, "stop");
        }

        if (log.isLoggable(Level.FINER)) {
            log.exiting(className, "stop");
        }

    }

    @Override
    public ProtocolDescription getDefaultProtocol() {
        if (log.isLoggable(Level.FINER)) {
            log.entering(className, "getDefaultProtocol");
        }

        String defaultProtocol = "Servlet 3.0";

        if (log.isLoggable(Level.FINER)) {
            log.exiting(className, "getDefaultProtocol", defaultProtocol);
        }

        return new ProtocolDescription(defaultProtocol);
    }

    @Override
    public ProtocolMetaData deploy(Archive<?> archive) throws DeploymentException {
        if (log.isLoggable(Level.FINER)) {
            log.entering(className, "deploy");
        }

        String archiveName = archive.getName();
        String archiveType = createDeploymentType(archiveName);
        String deployName = createDeploymentName(archiveName);

        if (!archiveType.equalsIgnoreCase("ear") && !archiveType.equalsIgnoreCase("war")
                && !archiveType.equalsIgnoreCase("eba")) {

            throw new DeploymentException("Invalid archive type: " + archiveType
                    + ".  Valid archive types are ear, war, and eba.");
        }

        try {
            File exportedArchiveLocation = new File(System.getProperty("java.io.tmpdir"), archiveName);
            archive.as(ZipExporter.class).exportTo(exportedArchiveLocation, true);
            restClient.deploy(exportedArchiveLocation);
            exportedArchiveLocation.deleteOnExit();
        } catch (Exception e) {
            throw new DeploymentException(e.getMessage());
        }

        // Wait until the application is deployed and available
        waitForApplicationTargetState(deployName, true, containerConfiguration.getAppDeployTimeout());

        // Return metadata on how to contact the deployed application
        ProtocolMetaData metaData = new ProtocolMetaData();
        HTTPContext httpContext = new HTTPContext(containerConfiguration.getHostName(), containerConfiguration.getHttpPort());
        httpContext.add(new Servlet("ArquillianServletRunner", deployName));
        metaData.addContext(httpContext);

        if (log.isLoggable(Level.FINER)) {
            log.exiting(className, "deploy");
        }

        return metaData;
    }

    @Override
    public void undeploy(Archive<?> archive) throws DeploymentException {
        if (log.isLoggable(Level.FINER)) {
            log.entering(className, "undeploy");
        }

        String archiveName = archive.getName();
        String deployName = createDeploymentName(archiveName);

        try {
            restClient.undeploy(archiveName);
        } catch (Exception e) {
            throw new DeploymentException("Error undeploying application " + archiveName + " " + e);
        }

        // Wait until the application is undeployed
        waitForApplicationTargetState(deployName, false, containerConfiguration.getAppUndeployTimeout());

        if (log.isLoggable(Level.FINER)) {
            log.exiting(className, "undeploy");
        }
    }

    @Override
    public void deploy(Descriptor descriptor) throws DeploymentException {
    }

    @Override
    public void undeploy(Descriptor descriptor) throws DeploymentException {
    }

    private String createDeploymentName(String archiveName) {
        return archiveName.substring(0, archiveName.lastIndexOf("."));
    }

    private String createDeploymentType(String archiveName) {
        return archiveName.substring(archiveName.lastIndexOf(".") + 1);
    }

    private void waitForApplicationTargetState(String applicationName, boolean targetState, int timeout) throws DeploymentException {
        if (log.isLoggable(Level.FINER)) {
            log.entering(className, "waitForApplicationTargetState", new Object[]{applicationName, targetState, timeout});
        }

        try {
            int timeleft = timeout * 1000;
            boolean curState = !targetState;
            while (curState != targetState) {
                Thread.sleep(100);
                curState = restClient.isApplicationStarted(applicationName);
                if (timeleft <= 0)
                    throw new DeploymentException("Timeout while waiting for ApplicationState to reach 'STARTED'");
                timeleft -= 100;
            }
        } catch (InterruptedException e) {
            throw new DeploymentException("Error occurred while while waiting for ApplicationState to reach STARTED " + e);
        }

        if (log.isLoggable(Level.FINER)) {
            log.exiting(className, "waitForApplicationTargetState");
        }
    }

}
