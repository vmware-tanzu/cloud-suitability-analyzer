package io.openliberty.arquillian.managed;

import javax.enterprise.inject.spi.DefinitionException;

import org.jboss.arquillian.container.test.api.Deployment;
import org.jboss.arquillian.container.test.api.ShouldThrowException;
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
public class WLPDuplicateAppNamePart2 {

    @ShouldThrowException(DefinitionException.class)
    @Deployment
    public static WebArchive createInvalidDeployment()
    {
        return ShrinkWrap.create(WebArchive.class, "duplicateName.war")
                .addClass(InvalidBean.class);
    }
    
    @Test
    public void deploymentExceptionThrown() {
        // No assertion needed here. Just want to check that the @ShouldThrowException above is honoured
    }
}
