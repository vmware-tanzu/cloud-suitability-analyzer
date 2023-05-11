delegate int WorkDelegate(int arg);
...
WorkDelegate del = DoWork;

// Calling del.BeginInvoke starts executing the delegate on a
// separate ThreadPool thread
Console.WriteLine("Starting with BeginInvoke");
var result = del.BeginInvoke(11, WorkCallback, null);

// This writes output to the console while DoWork is running in the background
Console.WriteLine("Waiting on work...");

// del.EndInvoke waits for the delegate to finish executing and 
// gets its return value
var ret = del.EndInvoke(result);