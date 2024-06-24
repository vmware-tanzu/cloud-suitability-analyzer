import msmq.*;
// Generated using 'com2java' from \WINNT\System32\mqoa.dll

import com.linar.jintegra.AuthInfo;

public class MsmqExample {
  public static void main(String[] args) throws Exception {
    try {
      String hostName = "localhost";
      String pathName = ".\\PRIVATE$\\Greeting2";
      // If running this Java client under MS windows then make sure the J-Integra® 'bin' directory
      // is in your PATH, otherwise run DCOMCNFG on the machine running MSMQ (Start|Run|DCOMCNFG),
      // grant a specific user default launch and access permissions, and uncomment the following
      // line, specifying the appropriate domain/user/password
      // com.linar.jintegra.AuthInfo.setDefault("NT DOMAIN", "NT USER", "NT PASSWORD");
      // Create the queue on the remote Windows machine using its IP address
      // MSMQQueueInfo qinfo = new MSMQQueueInfo("123.456.78.9");

      // Create the queue on the local Windows machine
      MSMQQueueInfo qinfo = new MSMQQueueInfo(hostName);
      qinfo.setPathName(pathName);
      System.out.println("pathName = " + pathName);
      qinfo.create(null, null);
      // Open the queue, and send a message
      IMSMQQueue q = qinfo.open(MQACCESS.MQ_SEND_ACCESS, MQSHARE.MQ_DENY_NONE);

      // Create the message
      IMSMQMessage msg = new MSMQMessage(hostName);

      msg.setLabel("Test Message");
      msg.setBody("This is a test message with a string Body.");
      msg.send(q, null);
      q.close();
      // Open the queue, and read the message
      q = qinfo.open(MQACCESS.MQ_RECEIVE_ACCESS, MQSHARE.MQ_DENY_NONE);
      msg = q.receive(null, null, null, null);
      System.out.println("Received: " + msg.getBody());
      q.close();
    } catch (Exception e) {
      e.printStackTrace();
    } finally {
      com.linar.jintegra.Cleaner.releaseAll();
    }
  }
}