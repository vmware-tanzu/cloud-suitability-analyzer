package io.openliberty.arquillian.managed;

/**
 * Interrupts a target thread after a specified length of time
 * <p>
 * Used to apply a timeout around waiting for process to end
 * <p>
 * Example usage:
 * <p><pre><code>
 * Interruptor interruptor = new Interruptor(Thread.currentThread(500));
 * try {
 *    interruptor.start();
 *    Thread.sleep(1000);
 *    interruptor.stop();
 * } catch (InterruptedException e) {
 *    // Was interrupted!
 * }
 * </code></pre>
 */
public class Interruptor {
   
   private final Thread interruptTarget;
   private final long timeToWait;
   private Thread taskThread;
   private boolean running = false;
   
   /**
    * Create an interruptor to interrupt the given target thread after the specified time
    * 
    * @param interruptTarget the thread this interruptor will interrupt
    * @param timeToWait the time it should wait before interrupting it
    */
   public Interruptor(Thread interruptTarget, long timeToWait) {
      this.interruptTarget = interruptTarget;
      this.timeToWait = timeToWait;
   }

   /**
    * Starts the countdown to the interruption of the target thread
    */
   public synchronized void start() {
      if (running) {
         throw new IllegalArgumentException("Interruptor has already been started");
      }
      taskThread = new Thread() {
         @Override
         public void run() {
            try {
               sleep(timeToWait);
               doInterrupt();
            } catch (InterruptedException e) {
               // stop() has been called
               // nothing to do, just return
            }
         }
      };
      running = true;
      taskThread.start();
   }

   /**
    * Stop the interruptor, cancelling the scheduled interruption
    * <p>
    * If an interruption has already occurred, this method will clear the thread interrupt flag
    * <p>
    * If an interruption has not occurred yet, this method will cancel the pending interruption 
    */
   public synchronized void stop() {
      running = false;
      taskThread.interrupt();
   }
   
   // Do the actual interruption
   // Synchronized to ensure that interruption will not occur after stop is called
   private synchronized void doInterrupt() {
      if (running) {
         interruptTarget.interrupt();
         running = false;
      }
   }

}
