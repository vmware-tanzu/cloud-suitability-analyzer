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

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.logging.Logger;

/**
 * Builds an exception cause chain from a list of exception class names and messages
 * <p>
 * This is used to turn a stack trace from the server into a chain of exceptions on the client.
 * <p>
 * If an exception class thrown by the server can't be loaded on the client, then it is replaced by an {@link UnloadableLogException}
 */
public class NestedExceptionBuilder {
    
    private static Logger log = Logger.getLogger(NestedExceptionBuilder.class.getName());
    
    /**
     * Mapping from some known implementation classes to the API classes they extend
     * <p>
     * In general, it's much more likely the client has the API classes available to load
     */
    private final static Map<String, String> exceptionMappings = new HashMap<>();
    static {
        exceptionMappings.put("org.jboss.weld.exceptions.DeploymentException", "javax.enterprise.inject.spi.DeploymentException");
        exceptionMappings.put("org.jboss.weld.exceptions.InconsistentSpecializationException", "javax.enterprise.inject.spi.DeploymentException");
        exceptionMappings.put("org.jboss.weld.exceptions.UnserializableDependencyException", "javax.enterprise.inject.spi.DeploymentException");
        exceptionMappings.put("org.jboss.weld.exceptions.DefinitionException", "javax.enterprise.inject.spi.DefinitionException");
    }
    
    /**
     * Small class the hold an potential nested exception
     */
    public static class ExMsg {
        String exName;
        String logMsg;
        List<String> superclasses = new ArrayList<>();
        public String toString(){
            return exName + "(" + logMsg + ")";
        }
    }
    
    /**
     * When we can't find/load a class from what we think is the
     * exception name we use a generic wrapper to allow us to
     * see the log line later
     */
    public static class UnloadableLogException extends RuntimeException {
        private static final long serialVersionUID = 1L;
        
        public UnloadableLogException(String logLine){
            super(logLine);
        }
        
        public UnloadableLogException(String logLine, Throwable cause){
            super(logLine, cause);
        }
        
    }
    
    
    /**
     * Create and return the nested FFDC exception chain as a Throwable with
     * cause(s)
     * 
     * @param exs
     *            - an ordered list of the Exception log messages
     * @return a Throwable with usable causal chain
     */
    public static Throwable buildNestedException(List<ExMsg> exs) {

        Throwable causeSoFar = null;

        for (int nestLevel = exs.size() - 1; nestLevel >= 0; nestLevel--) {
            causeSoFar = createException(exs.get(nestLevel), causeSoFar);
        }

        return causeSoFar;
    }

    /**
     * Turn a FFDC log line into a Throwable object. This depends on Class.forname
     * loading so maven pom dependencies need to include classes that need to be
     * loaded on the classpath
     * 
     * @param x - the exception message object
     * @param cause - something to nest inside the Throwable we create - can be null
     * @return - a non-null Throwable that reflects the log Line as well as
     *           possible. Note that fully 'private' classes are unlikely to be in
     *           tests' @ShouldThrow tests anyway - so in that case the important 
     *           thing is to not break the causal chain with this link. 
     *           We may have private subclasses of public Exceptions - we don't handle, 
     *           but these should be treated as special cases, handled elsewhere - for
     *           example with a service loaded DeploymentExceptionTransformer.
     */
    private static Throwable createException(ExMsg x, Throwable cause) {
        
        Throwable t = attemptCreation(x.exName, x.logMsg, cause);
        
        if (t == null) {
            // We couldn't load the actual class, try any superclasses if we know them
            for (String superclass : x.superclasses) {
                t = attemptCreation(superclass, x.logMsg, cause);
                if (t != null) {
                    break;
                }
            }
        }
        
        if (t == null) {
            // If we couldn't load the actual class, check if it's a known name that we can map to an API exception class
            String mappedExName = exceptionMappings.get(x.exName);
            if (mappedExName != null) {
                t = attemptCreation(mappedExName, x.logMsg, cause);
            }
        }
        
        if (t == null) {
            // If we still can't load the class, fall back to an UnloadableLogException
            log.warning("Unable to create an object for " + x.exName + ", falling back to UnloadableLogException");
            t = new UnloadableLogException("<" + x.exName + "> " + x.logMsg, cause);
        }
        
        log.finest("Actually created class= " + t.getClass().getName() + " instance msg="
                + t.getMessage());
        
        return t;
    }
    
    /**
     * Attempt to create an exception instance
     * 
     * @param className the exception class name
     * @param message the exception message
     * @param cause the exception cause
     * @return the exception, or null if it could not be created
     */
    private static Throwable attemptCreation(String className, String message, Throwable cause) {
        Class<? extends Throwable> clazz = null;

        // First, attempt to load the exception class thrown on the server
        // There will be lots of cases where this isn't possible
        try {
            clazz = loadThrowable(className);
        } catch (ClassNotFoundException e) {
            log.fine("Unable to load class " + className + ": " + e);
            return null;
        }
        
        Throwable t = null;
        
        if (cause == null) {
            // No cause, just a message
            t = attemptConstruction(clazz, message);
            
            if (t == null) {
                // Can also try with a message and a null cause
                t = attemptConstruction(clazz, message, null);
            }
            
            if (t == null) {
                // Final option, try the no-arg constructor
                t = attemptConstruction(clazz);
            }
        } else {
            // Cause and message
            t = attemptConstruction(clazz, message, cause);
            
            if (t == null) {
                // Some old exceptions construct with just a string and then initialise the cause
                t = attemptConstructionAndInit(clazz, message, cause);
            }
            
            if (t == null) {
                // Worst case, drop the message and just include the cause
                // Arquillian only matches on the exception class anyway
                t = attemptConstruction(clazz, cause);
            }
        }
        
        if (t == null) {
            log.fine("Unable to construct an instance of " + className);
        }
        
        return t;
    }
    
    /**
     * Attempt to get the Class object for a named Throwable
     * 
     * @param name the class name
     * @return the Class
     * @throws ClassNotFoundException if the class could not be loaded, or if it was loaded but was not a throwable
     */
    private static Class<? extends Throwable> loadThrowable(String name) throws ClassNotFoundException {
        try {
            Class<?> clazz = Class.forName(name);
            return clazz.asSubclass(Throwable.class);
        } catch (ClassNotFoundException e) {
            throw e;
        } catch (ClassCastException e) {
            throw new ClassNotFoundException("Loaded " + name + " successfully but it isn't a Throwable", e);
        } catch (Exception e) {
            throw new ClassNotFoundException("Unable to load " + name + " due to " + e, e);
        }
    }
    
    /**
     * Attempt to create an instance by calling its (String) constructor
     * @return the created instance, or null if it could not be created
     */
    private static Throwable attemptConstruction(Class<? extends Throwable> clazz, String msg) {
        try {
            return clazz.getConstructor(String.class).newInstance(new Object[] { msg });
        } catch (Exception e) {
            log.finer("Arquillian container could not construct " + clazz.getName() + "(String): " + e);
            return null;
        }
    }
    
    /**
     * Attempt to create an instance by calling its (String, Throwable) constructor
     * @return the created instance, or null if it could not be created
     */
    private static Throwable attemptConstruction(Class<? extends Throwable> clazz, String msg, Throwable cause) {
        try {
            return clazz.getConstructor(String.class, Throwable.class).newInstance(new Object[] { msg, cause });
        } catch (Exception e) {
            log.finer("Arquillian container could not construct " + clazz.getName() + "(String, Throwable): " + e);
            return null;
        }
    }
    
    /**
     * Attempt to create an instance by calling its (String) constructor and then calling initCause(cause)
     * @return the created instance, or null if it could not be created
     */
    private static Throwable attemptConstructionAndInit(Class<? extends Throwable> clazz, String msg, Throwable cause) {
        try {
            Throwable t = clazz.getConstructor(String.class).newInstance(new Object[] { msg });
            t.initCause(cause);
            return t;
        } catch (Exception e) {
            log.finer("Arquillian container could not construct " + clazz.getName() + "(String) and invoke initCause(): " + e);
            return null;
        }
    }
    
    /**
     * Attempt to create an instance by calling its (Throwable) constructor
     * @return the created instance, or null if it could not be created
     */
    private static Throwable attemptConstruction(Class<? extends Throwable> clazz, Throwable cause) {
        try {
            return clazz.getConstructor(Throwable.class).newInstance(new Object[] { cause });
        } catch (Exception e) {
            log.finer("Arquillian container could not construct " + clazz.getName() + "(Throwable): " + e);
            return null;
        }
    }
    
    private static Throwable attemptConstruction(Class<? extends Throwable> clazz) {
        try {
            return clazz.getConstructor().newInstance();
        } catch (Exception e) {
            log.finer("Arquillian container could not construct " + clazz.getName() + "(): " + e);
            return null;
        }
    }
}
