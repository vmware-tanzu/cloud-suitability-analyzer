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
package io.openliberty.arquillian.support;

import java.util.concurrent.atomic.AtomicReference;

import com.ibm.ws.container.service.app.deploy.ApplicationInfo;
import com.ibm.ws.container.service.state.ApplicationStateListener;
import com.ibm.ws.container.service.state.StateChangeException;
import com.ibm.wsspi.logging.Incident;
import com.ibm.wsspi.logging.IncidentForwarder;

public class IncidentListener implements IncidentForwarder, ApplicationStateListener {
    
    private AtomicReference<ExceptionInfo> lastException = new AtomicReference<>();
    private AtomicReference<String> appName = new AtomicReference<String>(null);

    @Override
    public void process(Incident incident, Throwable exception) {
        // We always expect a state change exception if there's a failure during deployment
        if (exception instanceof StateChangeException) {
            ExceptionInfo info = new ExceptionInfo(exception, appName.get());
            lastException.set(info);
        }
    }
   
    public ExceptionInfo getLastException() {
        return lastException.get();
    }
    
    public void clear() {
        lastException.set(null);
    }

    @Override
    public void applicationStarting(ApplicationInfo appInfo) throws StateChangeException {
        appName.set(appInfo.getName());
    }

    @Override
    public void applicationStarted(ApplicationInfo arg0) throws StateChangeException {
        // Do nothing
    }

    @Override
    public void applicationStopped(ApplicationInfo arg0) {
        // Do nothing
    }

    @Override
    public void applicationStopping(ApplicationInfo arg0) {
        // Do nothing
    }
    
    public static class ExceptionInfo {
        private Throwable exception;
        private String appName;
        
        public ExceptionInfo(Throwable exception, String appName) {
            this.exception = exception;
            this.appName = appName;
        }

        public Throwable getException() {
            return exception;
        }

        public String getAppName() {
            return appName;
        }
    }

}
