package runodischeduler;
/*  OdiInvokeScenario class  */


import oracle.odi.runtime.agent.invocation.ExecutionInfo;
import oracle.odi.runtime.agent.invocation.InvocationException;
import oracle.odi.runtime.agent.invocation.RemoteRuntimeAgentInvoker;
import oracle.odi.runtime.agent.invocation.StartupParams;

public class OdiInvokeScenario {

	private DatabaseManager dbManager;
	private int jobId;
	private int maxRetries;

	public OdiInvokeScenario(DatabaseManager databaseManager,int jobId) {

		this.dbManager = databaseManager;
		this.jobId = jobId;
		

	}

	public ExecutionInfo runScenario(String odiHost,String odiUser,String odiPassword,
			String odiScenName,String odiScenVersion,
			StartupParams odiStartupParams,String odiKeywords,
			String odiContextCode,int odiLogLevel,
			String odiSessionName,boolean odiSynchronous,
			String odiWorkRepName) throws InvocationException {
		
		
		ExecutionInfo exeInfo=null;
		try{

			//Starting the session on the remote agent.
			RemoteRuntimeAgentInvoker remoteRuntimeAgentInvoker = 
					new RemoteRuntimeAgentInvoker(odiHost,odiUser,
							odiPassword.toCharArray());
			//Parameter: 
			//pScenName,pScenVersion,pVariables,pKeywords,
			//pContextCode, pLogLevel, pSessionName, pSynchronous, pWorkRepName
			//pVariables ==> startupParams = new StartupParams();

			 exeInfo = 
					remoteRuntimeAgentInvoker.invokeStartScenario(odiScenName,odiScenVersion,odiStartupParams,
							odiKeywords,odiContextCode,odiLogLevel,
							odiSessionName,odiSynchronous,odiWorkRepName);

			
			
			
			//Retrieve the session ID.
			System.out.println(
					"-----------------------------------" + "\n" +  
							"Scenario Started in Session : " + exeInfo.getSessionId() + "\n" + 
							"Scenario Session Status : " + exeInfo.getSessionStatus() + "\n" +
					"-----------------------------------" );
		
		}
		catch (Exception e)    {
			e.printStackTrace();
		}
		return exeInfo;
	}

}

