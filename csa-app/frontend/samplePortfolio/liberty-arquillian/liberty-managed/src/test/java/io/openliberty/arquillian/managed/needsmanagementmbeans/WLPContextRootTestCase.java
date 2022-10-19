/*
 * Copyright 2018, IBM Corporation, and other contributors
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
package io.openliberty.arquillian.managed.needsmanagementmbeans;

import static org.junit.Assert.assertEquals;

import java.io.ByteArrayOutputStream;
import java.io.File;
import java.io.InputStream;
import java.net.URL;

import org.jboss.arquillian.container.test.api.Deployment;
import org.jboss.arquillian.container.test.api.OperateOnDeployment;
import org.jboss.arquillian.junit.Arquillian;
import org.jboss.arquillian.test.api.ArquillianResource;
import org.jboss.shrinkwrap.api.ShrinkWrap;
import org.jboss.shrinkwrap.api.asset.StringAsset;
import org.jboss.shrinkwrap.api.spec.EnterpriseArchive;
import org.jboss.shrinkwrap.api.spec.WebArchive;
import org.jboss.shrinkwrap.descriptor.api.Descriptors;
import org.jboss.shrinkwrap.descriptor.api.application7.ApplicationDescriptor;
import org.junit.Test;
import org.junit.runner.RunWith;

@RunWith(Arquillian.class)
public class WLPContextRootTestCase
{
    private static final String DEPLOYMENT1 = "app1";
    private static final String DEPLOYMENT2 = "app2";
    private static final String DEPLOYMENT3 = "app3";
    private static final String DEPLOYMENT4 = "app4";
    
    @Deployment(testable = false, name = DEPLOYMENT1)
    public static WebArchive app1()
    {
        return ShrinkWrap.create(WebArchive.class, "testApp1.war")
                .addClass(FooServlet.class);
    }
    
    @Deployment(testable = false, name = DEPLOYMENT2)
    public static WebArchive app2()
    {
        return ShrinkWrap.create(WebArchive.class, "testApp2.war")
            .addAsWebInfResource(new File("src/test/resources/ibm-web-ext_1.xml"), "ibm-web-ext.xml")
            .addClass(OtherFooServlet.class);
    }
    
    @Deployment(testable = false, name = DEPLOYMENT3)
    public static EnterpriseArchive app3() 
    {
        EnterpriseArchive ear = ShrinkWrap.create(EnterpriseArchive.class, "testApp3.ear")
            .addAsModule(ShrinkWrap.create(WebArchive.class, "test1.war")
                .addClass(BarServlet.class)
                .addAsWebInfResource(new File("src/test/resources/ibm-web-ext_2.xml"), "ibm-web-ext.xml"));
        
        ApplicationDescriptor appXml = Descriptors.create(ApplicationDescriptor.class)
           .version(ApplicationDescriptor.VERSION)
           .applicationName("testApp3")
           .createModule().getOrCreateWeb().contextRoot("/test1").webUri("test1.war").up().up();
        ear.setApplicationXML(new StringAsset(appXml.exportAsString()));
        
        return ear;
   }
    
    @Deployment(testable = false, name = DEPLOYMENT4)
    public static EnterpriseArchive app4() 
    {
        EnterpriseArchive ear = ShrinkWrap.create(EnterpriseArchive.class, "testApp4.ear")
            .addAsModule(ShrinkWrap.create(WebArchive.class, "test2.war")
                .addClass(BazServlet.class)
                .addAsWebInfResource(new File("src/test/resources/ibm-web-ext_3.xml"), "ibm-web-ext.xml"));
        return ear;
   }
    
   @ArquillianResource(FooServlet.class)
   @OperateOnDeployment(DEPLOYMENT1)
   private URL fooContextRoot;
   
   @ArquillianResource(OtherFooServlet.class)
   @OperateOnDeployment(DEPLOYMENT2)
   private URL otherFooContextRoot;
   
   @ArquillianResource(BarServlet.class)
   @OperateOnDeployment(DEPLOYMENT3)
   private URL barContextRoot;
   
   @ArquillianResource(BazServlet.class)
   @OperateOnDeployment(DEPLOYMENT4)
   private URL bazContextRoot;
   
   @Test
   public void testIfFooContextRootMatchesWarName() throws Exception 
   {
       URL url = new URL(fooContextRoot, "foo");
       assertEquals("/testApp1/", fooContextRoot.getPath());
       String response = readAllAndClose(url.openStream());
       assertEquals("I am foo", response);
   }
   
   @Test
   public void testIfOtherFooContextRootMatchesWebExtValue() throws Exception 
   {
       URL url = new URL(otherFooContextRoot, "otherFoo");
       assertEquals("/contextRootFromWebExt1/", otherFooContextRoot.getPath());
       String response = readAllAndClose(url.openStream());
       assertEquals("I am other foo", response);
   }
   
   @Test
   public void testIfBarContextRootMatchesApplicationXmlValue() throws Exception 
   {
       URL url = new URL(barContextRoot, "bar");
       assertEquals("/test1/", barContextRoot.getPath());
       String response = readAllAndClose(url.openStream());
       assertEquals("I am bar", response);
   }
   
   @Test
   public void testIfBazContextRootMatchesWebExtValue() throws Exception {
       URL url = new URL(bazContextRoot, "baz");
       assertEquals("/contextRootFromWebExt3/", bazContextRoot.getPath());
       String response = readAllAndClose(url.openStream());
       assertEquals("I am baz", response);
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