# Arquillian support Liberty user feature

A liberty user feature which allows deployment exceptions to be reported more reliably when using the liberty-managed container

The arquillian support feature adds an additional http endpoint which the arquillian container can query to determine the cause when an application fails to start.

It is only for supporting the running of arquillian tests and must not be installed on a production system.

## Usage

1. Extract the arquillian-liberty-support-x.x.x-feature.zip into the `usr` directory of your liberty runtime
1. Add `<feature>usr:arquillian-support-1.0</feature>` to the `<featureManager>` section of your `server.xml`

## Maven usage

You can install the arquillian-liberty-support feature as part of a maven build using the [maven-dependency-plugin:unpack goal](https://maven.apache.org/plugins/maven-dependency-plugin/unpack-mojo.html)

Example:

    <plugin>
      <groupId>org.apache.maven.plugins</groupId>
      <artifactId>maven-dependency-plugin</artifactId>
      <version>3.1.1</version>
      <executions>
        <execution>
          <id>extract-support-feature</id>
          <phase>pre-integration-test</phase>
          <goals>
            <goal>unpack</goal>
          </goals>
        </execution>
      </executions>
      <configuration>
        <artifactItems>
          <artifactItem>
            <groupId>io.openliberty.arquillian</groupId>
            <artifactId>arquillian-liberty-support</artifactId>
            <version>1.0.4</version>
            <type>zip</type>
            <classifier>feature</classifier>
            <overWrite>false</overWrite>
            <outputDirectory>${project.build.directory}/liberty/wlp/usr</outputDirectory>
          </artifactItem>
        </artifactItems>
      </configuration>
    </plugin>
