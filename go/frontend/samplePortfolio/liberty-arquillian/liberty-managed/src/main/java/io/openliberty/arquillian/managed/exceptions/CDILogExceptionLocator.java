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

import java.util.logging.Logger;

import javax.enterprise.inject.spi.DefinitionException;
import javax.enterprise.inject.spi.DeploymentException;

/**
 * This is old code to determine CDI exceptions from log lines
 */
public class CDILogExceptionLocator implements DeploymentExceptionLocator {
    
    private static Logger log = Logger.getLogger(CDILogExceptionLocator.class.getName());

    @Override
    public Throwable getException(String appName, String logLine, long deploymentTime) {
        if (logLine.contains("DefinitionException")) {
            log.finest("DefinitionException found in line " + logLine);
            return new DefinitionException(logLine);
         } else if (logLine.contains("DeploymentException") ||
                 logLine.contains("InconsistentSpecializationException") ||
                 logLine.contains("UnserializableDependencyException")) {
            /*
             * The CDI specification allows an implementation to throw a subclass of
             * javax.enterprise.inject.spi.DeploymentException. Weld has three types
             * such exceptions:
             *  - org.jboss.weld.exceptions.DeploymentException
             *  - org.jboss.weld.exceptions.InconsistentSpecializationException
             *  - org.jboss.weld.exceptions.UnserializableDependencyException
             */
            log.finest("DeploymentException found in line " + logLine);
            return new DeploymentException(logLine);
         } else {
             return null;
         }
    }

}
