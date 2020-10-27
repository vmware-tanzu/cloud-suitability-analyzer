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

import java.util.Dictionary;
import java.util.Hashtable;

import javax.servlet.ServletContext;
import javax.servlet.ServletContextEvent;
import javax.servlet.ServletContextListener;
import javax.servlet.annotation.WebListener;

import org.osgi.framework.BundleContext;
import org.osgi.framework.Constants;
import org.osgi.framework.ServiceRegistration;

import com.ibm.ws.container.service.state.ApplicationStateListener;
import com.ibm.ws.ffdc.FFDC;

@WebListener
public class Initializer implements ServletContextListener {
    
    public static final String INCIDENT_LISTENER_ATTRIBUTE = "incident-listener";
    public static final String BUNDLE_CONTEXT_ATTRIBUTE = "osgi-bundlecontext"; //OSGi 6.0 enterprise, 128.6.1
    
    private ServiceRegistration<ApplicationStateListener> appStateListenerRegistration;

    @Override
    public void contextInitialized(ServletContextEvent sce) {
        // Register the listener as an incident forwarder
        ServletContext context = sce.getServletContext();
        IncidentListener listener = new IncidentListener();
        context.setAttribute(INCIDENT_LISTENER_ATTRIBUTE, listener);
        FFDC.registerIncidentForwarder(listener);
        
        // Register the listener as an ApplicationStateListener
        BundleContext bContext = (BundleContext) context.getAttribute(BUNDLE_CONTEXT_ATTRIBUTE);
        Dictionary<String, Object> properties = new Hashtable<>();
        // Big number for the ranking so we get called early.
        // Other components use this event to initialize things for the app and may throw validation exceptions,
        // we need to know the app that's starting before an exception is thrown
        properties.put(Constants.SERVICE_RANKING, 5000); 
        appStateListenerRegistration = bContext.registerService(ApplicationStateListener.class, listener, properties);
    }

    @Override
    public void contextDestroyed(ServletContextEvent sce) {
        ServletContext context = sce.getServletContext();
        IncidentListener listener = (IncidentListener) context.getAttribute(INCIDENT_LISTENER_ATTRIBUTE);
        FFDC.deregisterIncidentForwarder(listener);
        appStateListenerRegistration.unregister();
    }

}
