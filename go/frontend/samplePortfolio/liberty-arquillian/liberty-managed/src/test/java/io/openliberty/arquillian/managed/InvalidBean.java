package io.openliberty.arquillian.managed;

import javax.enterprise.context.ApplicationScoped;

/**
 * This bean should result in a definition error because it has a public non-static field and is not Dependent scoped
 */
@ApplicationScoped
public class InvalidBean {
    
    public String foo = "test";
    
    public String getFoo() {
        return foo;
    }

}
