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
import org.jboss.shrinkwrap.api.exporter.ZipExporter;
import org.jboss.shrinkwrap.api.spec.EnterpriseArchive;
import org.jboss.shrinkwrap.api.spec.WebArchive;
import org.jboss.shrinkwrap.descriptor.api.Descriptors;
import org.jboss.shrinkwrap.descriptor.api.application7.ApplicationDescriptor;
import org.junit.Test;
import org.junit.runner.RunWith;


@RunWith(Arquillian.class)
public class WLPInjectServletContextTest
{
    
    private static final String DEPLOYMENT1 = "app1";
    private static final String DEPLOYMENT2 = "app2";
    private static final String DEPLOYMENT3 = "app3";
    
    @Deployment(testable = false, name = DEPLOYMENT1)
    public static WebArchive app1() {
        return ShrinkWrap.create(WebArchive.class)
                .addClass(FooServlet.class);
    }
    
    @Deployment(testable = false, name = DEPLOYMENT2)
    public static EnterpriseArchive app2() {
        EnterpriseArchive ear = ShrinkWrap.create(EnterpriseArchive.class)
                .addAsModule(ShrinkWrap.create(WebArchive.class, "test1.war")
                             .addClass(BarServlet.class))
                .addAsModule(ShrinkWrap.create(WebArchive.class, "test2.war")
                             .addClass(BazServlet.class));
        
        ApplicationDescriptor appXml = Descriptors.create(ApplicationDescriptor.class)
                                                     .version(ApplicationDescriptor.VERSION)
                                                     .applicationName("testApp")
                                                     .createModule().getOrCreateWeb().contextRoot("/test1").webUri("test1.war").up().up()
                                                     .createModule().getOrCreateWeb().contextRoot("/test2").webUri("test2.war").up().up()
                                                     ;
        ear.setApplicationXML(new StringAsset(appXml.exportAsString()));
        
        return ear;
    }
    
    @Deployment(testable = false, name = DEPLOYMENT3)
    public static EnterpriseArchive app3() {

        EnterpriseArchive ear = ShrinkWrap.create(EnterpriseArchive.class)
                .addAsModule(ShrinkWrap.create(WebArchive.class, "test3.war")
                        .addClass(BuzServlet.class));

        File deploymentFile = new File("target/testAppWithLoadedFromFile.ear");

        ApplicationDescriptor appXml = Descriptors.create(ApplicationDescriptor.class)
                .version(ApplicationDescriptor.VERSION)
                .applicationName("testAppWithLoadedFromFile")
                .createModule().getOrCreateWeb().contextRoot("/test3").webUri("test3.war").up().up();
        ear.setApplicationXML(new StringAsset(appXml.exportAsString()));
        ear.as(ZipExporter.class).exportTo(deploymentFile, true);

        // Create the ShrinkWrap archive from file system, which leads to leading
        // slashes '/' in the
        // web module names.

        return ShrinkWrap.createFromZipFile(EnterpriseArchive.class, deploymentFile);
    }
    
    @ArquillianResource(FooServlet.class)
    @OperateOnDeployment(DEPLOYMENT1)
    private URL fooContextRoot;
    
    @ArquillianResource(BarServlet.class)
    @OperateOnDeployment(DEPLOYMENT2)
    private URL barContextRoot;
    
    @ArquillianResource(BazServlet.class)
    @OperateOnDeployment(DEPLOYMENT2)
    private URL bazContextRoot;
    
    @ArquillianResource(BuzServlet.class)
    @OperateOnDeployment(DEPLOYMENT3)
    private URL buzContextRoot;
    
    @Test
    public void testFoo() throws Exception {
        URL url = new URL(fooContextRoot, "foo");
        String response = readAllAndClose(url.openStream());
        assertEquals("I am foo", response);
    }
    
    @Test
    public void testBar() throws Exception {
        URL url = new URL(barContextRoot, "bar");
        String response = readAllAndClose(url.openStream());
        assertEquals("I am bar", response);
    }
    
    @Test
    public void testBaz() throws Exception {
        URL url = new URL(bazContextRoot, "baz");
        String response = readAllAndClose(url.openStream());
        assertEquals("I am baz", response);
    }
    
    @Test
    public void testBuz() throws Exception {
        URL url = new URL(buzContextRoot, "buz");
        String response = readAllAndClose(url.openStream());
        assertEquals("I am buz", response);
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
