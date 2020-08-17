package io.openliberty.arquillian.managed;

import org.jboss.arquillian.container.test.api.Deployment;
import org.jboss.arquillian.junit.Arquillian;
import org.jboss.shrinkwrap.api.ShrinkWrap;
import org.jboss.shrinkwrap.api.asset.EmptyAsset;
import org.jboss.shrinkwrap.api.spec.JavaArchive;
import org.junit.Test;
import org.junit.runner.RunWith;
import javax.annotation.Resource;
import org.junit.Assert;

@RunWith(Arquillian.class)
public class WLPResourceTestCase {

    @Deployment
    public static JavaArchive createDeployment() {
        return ShrinkWrap.create(JavaArchive.class)
                .addAsManifestResource(EmptyAsset.INSTANCE, "beans.xml");
    }

    @Resource(lookup = "env/foo")
    String foo;

    @Test
    public void serverXMLResourceInjectionShouldHaveCorrectValue(){
        Assert.assertEquals("bar" ,foo);
    }

    @Test
    public void serverEnvironmentVariablesShouldBeSet(){
        Assert.assertNotNull(System.getenv("foo"));
        Assert.assertNotNull(System.getenv("keystore_password"));
    }

}
