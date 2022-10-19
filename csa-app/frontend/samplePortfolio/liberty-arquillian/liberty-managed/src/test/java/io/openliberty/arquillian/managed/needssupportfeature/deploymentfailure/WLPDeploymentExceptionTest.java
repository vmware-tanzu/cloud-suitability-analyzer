/*
 * Copyright 2019, IBM Corporation, Red Hat Middleware LLC, and individual contributors
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
package io.openliberty.arquillian.managed.needssupportfeature.deploymentfailure;

import static org.junit.Assert.fail;

import javax.enterprise.inject.spi.Extension;

import org.jboss.arquillian.container.test.api.Deployer;
import org.jboss.arquillian.container.test.api.Deployment;
import org.jboss.arquillian.junit.Arquillian;
import org.jboss.arquillian.test.api.ArquillianResource;
import org.jboss.shrinkwrap.api.ShrinkWrap;
import org.jboss.shrinkwrap.api.asset.EmptyAsset;
import org.jboss.shrinkwrap.api.spec.WebArchive;
import org.junit.Test;
import org.junit.runner.RunWith;

@RunWith(Arquillian.class)
public class WLPDeploymentExceptionTest {
    
    public static final String INVALID_APP = "invalidApp";

    @Deployment(name = INVALID_APP, managed = false)
    public static WebArchive buildInvalidApp() {
        
        WebArchive war = ShrinkWrap.create(WebArchive.class)
                                   .addPackage(WLPDeploymentExceptionTest.class.getPackage())
                                   .addAsServiceProvider(Extension.class, StartupFailureExtension.class)
                                   .addAsWebInfResource(EmptyAsset.INSTANCE, "beans.xml");

        return war;
    }
    
    @ArquillianResource
    Deployer deployer;
    
    /**
     * Deploy the invalid app repeatedly to ensure that it throws the correct deployment error every time
     */
    @Test
    public void testDeploymentErrors() {
        for (int i = 0; i < 15; i++) {
            try {
                deployer.deploy(INVALID_APP);
                fail("App deployed successfully");
            } catch (Exception ex) {
                assertExceptionInCauseChain("Incorrect exception on deployment " + i, ex, TestAppException.class);
            } finally {
                deployer.undeploy(INVALID_APP);
            }
        }
    }

    private void assertExceptionInCauseChain(String message, Throwable ex, Class<? extends Throwable> expectedExceptionClass) {
        Throwable currentCause = ex;
        while (currentCause != null) {
            if (expectedExceptionClass.isInstance(currentCause)) {
                return;
            }
            currentCause = currentCause.getCause();
        }
        throw new AssertionError(message + ": Exception cause chain did not include an instance of " + expectedExceptionClass, ex);
    }

}
