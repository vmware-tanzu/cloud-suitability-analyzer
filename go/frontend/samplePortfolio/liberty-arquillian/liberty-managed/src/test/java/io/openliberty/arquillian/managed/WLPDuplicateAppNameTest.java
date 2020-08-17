package io.openliberty.arquillian.managed;

import org.junit.runner.RunWith;
import org.junit.runners.Suite;
import org.junit.runners.Suite.SuiteClasses;

@RunWith(Suite.class)
@SuiteClasses({
              WLPDuplicateAppNamePart1.class,
              WLPDuplicateAppNamePart2.class
})
/**
 * Test that we can deploy an app successfully, and then deploy an invalid app
 * with the same name and process the deployment exception correctly
 * <p>
 * Arquillian checks for two deployments in the same test having the same
 * archive name, so this situation is only possible if the two apps are deployed
 * in different tests. Therefore we have to use a suite to ensure we execute
 * them in the correct order.
 */
public class WLPDuplicateAppNameTest {}
