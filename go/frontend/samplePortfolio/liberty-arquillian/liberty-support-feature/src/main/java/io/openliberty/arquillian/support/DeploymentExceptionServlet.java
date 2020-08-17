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

import java.io.IOException;
import java.io.ObjectOutputStream;
import java.io.PrintWriter;
import java.util.HashSet;
import java.util.Set;

import javax.servlet.ServletContext;
import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import io.openliberty.arquillian.support.IncidentListener.ExceptionInfo;

@WebServlet("/deployment-exception")
public class DeploymentExceptionServlet extends HttpServlet {

    private static final long serialVersionUID = 1L;
    private static final Set<Class<?>> topClasses = new HashSet<>();
    static {
        topClasses.add(Object.class);
        topClasses.add(Throwable.class);
        topClasses.add(Exception.class);
        topClasses.add(Error.class);
        topClasses.add(RuntimeException.class);
    }
    

    @Override
    protected void doGet(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
        IncidentListener listener = getIncidentListener(req);
        
        resp.setContentType("text/plain;charset=UTF-8");
        ExceptionInfo ex = listener.getLastException();
        
        String appName = req.getParameter("appName");
        
        if (appName == null) {
            resp.setStatus(400);
            resp.getWriter().println("No appName given");
            return;
        }
        
        if (ex == null) {
            resp.setStatus(400);
            resp.getWriter().println("No exception logged");
            return;
        }
        
        if (!appName.equals(ex.getAppName())) {
            resp.setStatus(400);;
            resp.getWriter().println("Last exception was not thrown by the requested app");
            return;
        }
        
        String format = req.getParameter("format");
        if (format == null) {
            resp.setStatus(400);
            resp.getWriter().println("Format parameter not set");
        } else if (format.equals("text")) {
            printText(resp.getWriter(), ex.getException());
        } else if (format.equals("stack")) {
            ex.getException().printStackTrace(resp.getWriter());
        } else if (format.equals("serialize")) {
            resp.setContentType("application/java-serialized-object");
            try (ObjectOutputStream oos = new ObjectOutputStream(resp.getOutputStream())) {
                oos.writeObject(ex.getException());
            }
        } else {
            resp.setStatus(400);
            resp.getWriter().println("Invalid format requested");
        }
    }

    @Override
    protected void doPost(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
        if ("true".equalsIgnoreCase(req.getParameter("clear"))) {
            getIncidentListener(req).clear();
        } else {
            throw new ServletException("Invalid command");
        }
    }

    private IncidentListener getIncidentListener(HttpServletRequest req) {
        ServletContext context = req.getServletContext();
        IncidentListener listener = (IncidentListener) context.getAttribute(Initializer.INCIDENT_LISTENER_ATTRIBUTE);
        if (listener == null) {
            throw new RuntimeException("Incident listener is not set");
        }
        return listener;
    }
    
    private void printText(PrintWriter resp, Throwable exception) {
        Throwable t = exception;
        HashSet<Throwable> processed = new HashSet<>();
        while (t != null && !processed.contains(t)) {
            printClass(resp, t);
            processed.add(t);
            t = t.getCause();
        }
    }
    
    private void printClass(PrintWriter resp, Throwable exception) {
        resp.print("exClass ");
        resp.print(exception.getClass().getName());
        resp.println();
        
        Class<?> clazz = exception.getClass().getSuperclass();
        while (clazz != null && !topClasses.contains(clazz)) {
            resp.print("exSuperclass ");
            resp.print(clazz.getName());
            resp.println();
            clazz = clazz.getSuperclass();
        }
        
        resp.println(exception.getMessage());
    }

}
