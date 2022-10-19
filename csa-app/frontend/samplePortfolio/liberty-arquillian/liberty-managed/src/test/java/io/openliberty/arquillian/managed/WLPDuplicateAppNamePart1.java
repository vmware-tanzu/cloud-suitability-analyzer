package io.openliberty.arquillian.managed;

import org.jboss.arquillian.container.test.api.Deployment;
import org.jboss.arquillian.junit.Arquillian;
import org.jboss.shrinkwrap.api.ShrinkWrap;
import org.jboss.shrinkwrap.api.spec.WebArchive;
import org.junit.Test;
import org.junit.runner.RunWith;

/**
 * Run from the {@link WLPDuplicateAppNameTest} suite because ordering is important.
 * <p>
 * Note: This test doesn't get run by surefire directly because its name doesn't start or end in "Test"
 */
@RunWith(Arquillian.class)
public class WLPDuplicateAppNamePart1 {

    @Deployment
    public static WebArchive createDeployment() 
    {
       return ShrinkWrap.create(WebArchive.class, "duplicateName.war")
               .addClass(HelloServlet.class);
    }
    
    @Test
    public void duplicateName1Deploys() {
        // Just want to check it deploys correctly
    }
}
