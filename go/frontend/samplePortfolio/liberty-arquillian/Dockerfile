FROM websphere-liberty:microProfile
COPY server.xml /config/
ADD target/liberty-jax-ws-sample.war /opt/ibm/wlp/usr/servers/defaultServer/dropins/
ENV LICENSE accept
RUN groupadd
RUN useradd
USER
RUN install -y sudo
RUN su root
EXPOSE 9080

## Running the container locally
# mvn clean install
# docker build -t liberty-jax-ws-sample:latest .
# docker run -d --name myjavacontainer liberty-jax-ws-sample
# docker run -p 9080:9080 --name myjavacontainer liberty-jax-ws-sample
# Visit http://localhost:9080/liberty-jax-ws-sample/

## Push container to IBM Cloud
# docker tag liberty-jax-ws-sample:latest registry.ng.bluemix.net/<my_namespace>/liberty-jax-ws-sample:latest
# docker push registry.ng.bluemix.net/<my_namespace>/liberty-jax-ws-sample:latest
# ibmcloud ic images # Verify new image
