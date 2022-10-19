/*
 * Copyright 2012, 2014, Red Hat Middleware LLC, and other contributors
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

import org.jboss.arquillian.container.spi.ConfigurationException;
import org.jboss.arquillian.container.spi.client.container.ContainerConfiguration;

/**
 * <p>
 * WLPremoteContainerConfiguration
 * </p>
 * 
 * <p>
 * Adapted from WLPManagedContainerconfiguration
 * </p>
 *
 * @author <a href="mailto:tayres@gmail.com">Tony Ayres</a>
 * @version $Revision: $
 */
public class WLPRemoteContainerConfiguration implements ContainerConfiguration {

    private String serverName = "defaultServer";
    private int serverStartTimeout = 30;
    private int appDeployTimeout = 20;
    private int appUndeployTimeout = 2;

    private String username;

    private String password;
    private String hostName;
    private int httpPort = 9080;
    private int httpsPort = 9443;
    private boolean outputToConsole = true;

    @Override
    public void validate() throws ConfigurationException {

        if (hostName == null || hostName.equals("")) {
            throw new ConfigurationException("hostName is required for initialization");
        }

        if (username == null) {
            throw new ConfigurationException("username is required for initialization");
        }

        if (password == null) {
            throw new ConfigurationException("password is required for initialization");
        }

        // Validate httpPort
        if (httpPort > 65535 || httpPort < 0)
            throw new ConfigurationException("httpPort provided is not valid: " + httpPort);

        if (httpsPort > 65535 || httpsPort < 0)
            throw new ConfigurationException("httpsPort provided is not valid: " + httpPort);

    }

    public String getServerName() {
        return serverName;
    }

    public void setServerName(String serverName) {
        this.serverName = serverName;
    }

    public int getHttpPort() {
        return httpPort;
    }

    public void setHttpPort(int httpPort) {
        this.httpPort = httpPort;
    }

    public boolean isOutputToConsole() {
        return outputToConsole;
    }

    public void setOutputToConsole(boolean outputToConsole) {
        this.outputToConsole = outputToConsole;
    }

    public int getServerStartTimeout() {
        return serverStartTimeout;
    }

    public void setServerStartTimeout(int serverStartTimeout) {
        this.serverStartTimeout = serverStartTimeout;
    }

    public int getAppDeployTimeout() {
        return appDeployTimeout;
    }

    public void setAppDeployTimeout(int appDeployTimeout) {
        this.appDeployTimeout = appDeployTimeout;
    }

    public int getAppUndeployTimeout() {
        return appUndeployTimeout;
    }

    public void setAppUndeployTimeout(int appUndeployTimeout) {
        this.appUndeployTimeout = appUndeployTimeout;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public String getHostName() {
        return hostName;
    }

    public void setHostName(String hostName) {
        this.hostName = hostName;
    }

    public int getHttpsPort() {
        return httpsPort;
    }

    public void setHttpsPort(int httpsPort) {
        this.httpsPort = httpsPort;
    }

}
