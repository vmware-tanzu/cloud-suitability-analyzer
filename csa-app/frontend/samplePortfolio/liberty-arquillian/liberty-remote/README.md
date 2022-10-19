# Arquillian Liberty Remote Documentation

An Arquillian container adapter (`DeployableContainer` implementation) that can connect and run against a remote (different JVM, different machine) Liberty server andrun tests on it over a remote protocol (effectively in a different JVM).

## Prerequisites

**Prerequisite Version**

This `DeployableContainer` has been tested with the latest two releases of Open Liberty and WebSphere Liberty. 

**Prerequisite Configuration**

The following features are required in the `server.xml` of the Liberty server.

```
<!-- Enable features -->
<featureManager>
    <feature>jsp-2.2</feature>
    <feature>restConnector-1.0</feature>
</featureManager>
```

You will also need to enable security, one example would be:

```
<quickStartSecurity userName="admin" userPassword="admin" />

<keyStore id="defaultKeyStore" password="password" />
```

You need to have those keys trusted by your client as well, otherwise you'll see SSL certificate trust errors, and you need to give permissions for the container adapter to write to the dropins directory:

````
<!-- This section is needed to allow upload of files to the dropins directory,
the remote container adapter relies on this configuration -->
<remoteFileAccess>
    <writeDir>${server.config.dir}/dropins</writeDir>
</remoteFileAccess>
````

If you need a sample `server.xml`, please refer to the [one in our source repository](https://github.com/OpenLiberty/liberty-arquillian/blob/master/liberty-remote/src/test/resources/server.xml).

## Configuration

Default Protocol: Servlet 3.0

**Container Configuration Options**

| Name | Type | Default | Description |
| ---- | ---- | ------- | ----------- |
| serverName | String | defaultServer | Name of the liberty server instance to connect to. |
| hostName | String | None | Hostname for the target machine where the WebSphere Liberty Profile server is running. |
| username | String | None | The username to use to connect to the target. |
| password | String | None | The password to use to connect to the target. |
| httpPort | Integer | 9080 | HTTP Port of the server. |
| httpsPort | Integer | 9443 | HTTPS Port of the server. |
| serverStartTimeout | Integer | 30 | Time in seconds to wait for the application server to start |
| appDeployTimeout | Integer | 20 | Time in seconds to wait for the application deployment to complete and the application to start |
| appUndeployTimeout | Integer | 2 | Time in seconds to wait for the application undeployment to complete |
| outputToConsole | Boolean | true | When enabled output from the application server process will be emitted to stdout |

## Examples

### Example of Container Configuration (arquillian.xml)

```
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<arquillian xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
xmlns="http://jboss.org/schema/arquillian"
xsi:schemaLocation="http://jboss.org/schema/arquillian http://jboss.org/schema/arquillian/arquillian_1_0.xsd">

	<engine>
		<property name="deploymentExportPath">target/</property>
	</engine>
	
	<container qualifier="liberty-remote">
		<configuration>
			<property name="hostName">localhost</property>
			<property name="serverName">defaultServer</property>
			<property name="username">admin</property>
			<property name="password">admin</property>
			<property name="httpPort">9080</property>
			<property name="httpsPort">9443</property>
		</configuration>
	</container>
</arquillian>
```

### Example of Maven profile setup

```
<profile>
	<id>liberty-remote</id>
	<dependencies>
		<dependency>
			<groupId>io.openliberty.arquillian</groupId>
			<artifactId>arquillian-liberty-remote</artifactId>
			<version>1.0.0</version>
		</dependency>
	</dependencies>
</profile>
```
