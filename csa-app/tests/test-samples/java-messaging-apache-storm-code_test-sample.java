package storm;

import backtype.storm.Config;
import backtype.storm.LocalCluster;
import backtype.storm.StormSubmitter;
import backtype.storm.generated.AlreadyAliveException;
import backtype.storm.generated.InvalidTopologyException;
import backtype.storm.testing.TestWordSpout;
import backtype.storm.topology.BasicOutputCollector;
import backtype.storm.topology.OutputFieldsDeclarer;
import backtype.storm.topology.TopologyBuilder;
import backtype.storm.topology.base.BaseBasicBolt;
import backtype.storm.tuple.Fields;
import backtype.storm.tuple.Tuple;
import backtype.storm.tuple.Values;

/**
 * This topology is a simple topology with one spout and one bolt. The spout sends random words to
 * the bolt and the bolt appends a string to the tuple.
 *   
 * @author Zachary Radtka
 *
 */
public class BasicAppendTopology {
	
	/** The name of the topology */
	public static final String TOPOLOGY_NAME = "appendTopology";
	
	/** The name of the spout */
	public static final String SPOUT_NAME = "spout";
	
	/** The name of the bolt */
	public static final String BOLT_NAME_APPEND = "append";

	/**
	 * The <code>AppendBolt</code> class is a simple bolt that takes an incoming tuple, appends a 
	 * value to the tuple, and outputs the tuple.
	 *  
	 * @author Zachary Radtka
	 *
	 */
	public static class AppendBolt extends BaseBasicBolt {

		/** use serialVersionUID for interoperability */
		private static final long serialVersionUID = 6283801867001033058L;

		/** The name of the output field */
		private static final String OUTPUT_FIELD_NAME = "word";

		/** The value to append to the incoming tuple */
		private final String appendValue;
		
		/**
		 * Initializes an <code>AppendBolt</code>
		 * 
		 * @param value	
		 * 		A <code>String</code> to append to the tuple
		 */
		public AppendBolt(String value) {
			this.appendValue = value;
		}
		
		@Override
		public void execute(Tuple input, BasicOutputCollector collector) {
			collector.emit(new Values(input.getString(0) + this.appendValue));
		}

		@Override
		public void declareOutputFields(OutputFieldsDeclarer declarer) {
			declarer.declare(new Fields(OUTPUT_FIELD_NAME));
		}

	}
	
	/**
	 * Display the application usage to STDERR
	 */
	public static void displayUsage(){
		System.err.println("[options]");
		System.err.println("OPTIONS");
		System.err.println("\t-n=NAME\n\t\tset the topology name to NAME");
		System.err.println("\t-r\n\t\texecute on a remote Storm cluster");
		System.err.println("\t-v=VALUE\n\t\tappend VALUE to the end of words");
	}
	
	/**
	 * Submits (runs) the topology.
     *
     * Usage: "BasicAppendTopology [options]"
     *
     * By default, the topology is run locally under the name "appendTopology".
     *
     * Examples:
	 * 
     * <pre>
     *    <code>
     * # Runs in local mode (LocalCluster), with topology name "appendTopology"
     * $ storm jar storm-examples-jar-with-dependencies.jar storm.BasicAppendTopology
     *
     * # Runs in local mode (LocalCluster), with topology append value "?"
     * $ storm jar storm-examples-jar-with-dependencies.jar storm.BasicAppendTopology -v=?
     *
     * # Runs in remote/cluster mode, with topology name "production-topology"
     * $ storm jar storm-examples-jar-with-dependencies.jar storm.BasicAppendTopology -v=? -n=production-topology -r 
     *  </code>
     * </pre>
	 * 
	 * @param args
	 * 			Three arguments are possible in any order: -n=TOPOLOGY_NAME, -r, -v=VALUE
	 * @throws AlreadyAliveException
	 * @throws InvalidTopologyException
	 * @throws InterruptedException
	 */
	public static void main(String[] args) throws AlreadyAliveException, InvalidTopologyException,
			InterruptedException {

		boolean runLocally = true;
		String topologyName = TOPOLOGY_NAME;
		String appendValue = "!";
		
		// Ensure that the correct arguments were specified at runtime
		if (args.length > 3) {
			displayUsage();
			System.exit(1);
		}
		
		// Get the values from the command line
		for (int i = 0; i < args.length; i++) {
			if (args[i].startsWith("-n=")) {
				topologyName = topologyName.concat(args[i].substring(args[i].indexOf('=') + 1, args[i].length()));
			} else if (args[i].equalsIgnoreCase("-r")){
				runLocally = false;
			} else if (args[i].startsWith("-v=")) {
				appendValue = args[i].substring(args[i].indexOf('=') + 1, args[i].length());			
			} else {
				displayUsage();
				System.exit(1);
			}
		}
		
		// Create a new topology
		TopologyBuilder builder = new TopologyBuilder();

		// Set the spout
		builder.setSpout(SPOUT_NAME, new TestWordSpout(), 10);

		// Set the bolt to our Appender
		builder.setBolt(BOLT_NAME_APPEND, new AppendBolt(appendValue), 5).shuffleGrouping(SPOUT_NAME);


		// Turn debug on/off to see output from all bolts
		Config conf = new Config();
		conf.setDebug(true);

		// The topology can be run locally or remotely 
		if (runLocally) {
			conf.setMaxTaskParallelism(3);
			LocalCluster cluster = new LocalCluster();
			cluster.submitTopology(topologyName, conf, builder.createTopology());
			Thread.sleep(10000);
			cluster.shutdown();			
		} else {
			conf.setNumWorkers(3);
			StormSubmitter.submitTopologyWithProgressBar(topologyName, conf, builder.createTopology());
		}

	}
}