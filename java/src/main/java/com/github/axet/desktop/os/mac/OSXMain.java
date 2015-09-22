package com.github.axet.desktop.os.mac;

public class OSXMain {

    static int count = 0;

    static final Object lock = new Object();

    static Thread t;

    public static void inc() {
        if (t == null) {
            t = new Thread(new Runnable() {
                public void run() {
                    synchronized (lock) {
                        try {
                            lock.wait();
                        } catch (InterruptedException e) {
                            Thread.currentThread().interrupt();
                        }
                    }
                }
            });
            t.setDaemon(false);
            t.start();
        }

        count++;
    }

    public static void dec() {
        count--;

        if (count == 0) {
            synchronized (lock) {
                lock.notifyAll();
            }
        }
    }

}
