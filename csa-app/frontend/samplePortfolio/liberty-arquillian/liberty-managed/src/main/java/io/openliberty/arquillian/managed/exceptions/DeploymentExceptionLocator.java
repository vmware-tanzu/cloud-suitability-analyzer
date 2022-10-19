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

public interface DeploymentExceptionLocator {
    
    /**
     * Return the exception that caused the most recently deployed application not to start
     * 
     * @param appName the name of the application
     * @param logLine the deployment failure line from the log
     * @param deploymentTime the time in ms since the epoch that the application was deployed, can be used to limit the time to search
     * @return the exception that caused the deployment failure
     */
    public Throwable getException(String appName, String logLine, long deploymentTime);

}
