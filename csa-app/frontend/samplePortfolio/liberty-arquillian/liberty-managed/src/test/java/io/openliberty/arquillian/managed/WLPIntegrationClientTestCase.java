/*
 * Copyright 2010-2012, Red Hat Middleware LLC, and individual contributors
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
package io.openliberty.arquillian.managed;

import java.io.ByteArrayOutputStream;
import java.io.InputStream;
import java.net.URL;

import org.jboss.arquillian.container.test.api.Deployment;
import org.jboss.arquillian.junit.Arquillian;
import org.jboss.arquillian.test.api.ArquillianResource;
import org.jboss.shrinkwrap.api.ShrinkWrap;
import org.jboss.shrinkwrap.api.asset.EmptyAsset;
import org.jboss.shrinkwrap.api.spec.EnterpriseArchive;
import org.jboss.shrinkwrap.api.spec.JavaArchive;
import org.jboss.shrinkwrap.api.spec.WebArchive;
import org.junit.Assert;
import org.junit.Test;
import org.junit.runner.RunWith;

// This is based on the JettyEmbeddedClientTestCase.java
// written by Dan Allen and Aslak Knutsen.

@RunWith(Arquillian.class)
public class WLPIntegrationClientTestCase
{
   @Deployment(testable=false)
   public static EnterpriseArchive createDeployment() 
   {
      return ShrinkWrap.create(EnterpriseArchive.class, "test.ear")
            .addAsModule(ShrinkWrap.create(WebArchive.class, "test.war")
                              .addClass(HelloServlet.class))
            .addAsModule(ShrinkWrap.create(JavaArchive.class, "test.jar")
                              .addClass(WLPIntegrationClientTestCase.class)
                              .addAsManifestResource(EmptyAsset.INSTANCE, "beans.xml"));
   }
   
   @Test
   public void shouldBeAbleToInvokeServletInDeployedWebApp(@ArquillianResource URL url) throws Exception
   {
      String body = readAllAndClose(
            new URL(url, HelloServlet.URL_PATTERN).openStream());
      
      Assert.assertEquals(
            "Verify that the servlet was deployed and returns expected result",
            HelloServlet.MESSAGE,
            body);
   }
   
   private String readAllAndClose(InputStream is) throws Exception 
   {
      ByteArrayOutputStream out = new ByteArrayOutputStream();
      try
      {
         int read;
         while( (read = is.read()) != -1)
         {
            out.write(read);
         }
      }
      finally 
      {
         try { is.close(); } catch (Exception e) { }
      }
      return out.toString();
   }
}
