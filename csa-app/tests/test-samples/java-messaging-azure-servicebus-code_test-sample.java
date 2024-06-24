TokenCredential credential = new DefaultAzureCredentialBuilder().build();

// 'fullyQualifiedNamespace' will look similar to "{your-namespace}.servicebus.windows.net"
// Any clients created from this builder will share the underlying connection.
ServiceBusClientBuilder sharedConnectionBuilder = new ServiceBusClientBuilder()
    .credential(fullyQualifiedNamespace, credential);

// Create receiver and sender which will share the connection.
ServiceBusReceiverClient receiver = sharedConnectionBuilder
    .receiver()
    .receiveMode(ServiceBusReceiveMode.PEEK_LOCK)
    .queueName(queueName)
    .buildClient();
ServiceBusSenderClient sender = sharedConnectionBuilder
    .sender()
    .queueName(queueName)
    .buildClient();

// Use the clients and finally close them.
try {
    sender.sendMessage(new ServiceBusMessage("payload"));
    receiver.receiveMessages(1);
} finally {
    // Clients should be long-lived objects as they require resources
    // and time to establish a connection to the service.
    sender.close();
    receiver.close();
}