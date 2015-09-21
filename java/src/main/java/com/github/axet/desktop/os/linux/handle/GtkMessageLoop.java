package com.github.axet.desktop.os.linux.handle;

import com.github.axet.desktop.os.linux.libs.LibGtk;
import com.sun.jna.Pointer;

public class GtkMessageLoop {
    Object lock = new Object();

    static GMainLoop mainloop;
    static GMainContext context;

    Thread MessageLoop = new Thread(new Runnable() {
        @Override
        public void run() {
            LibGtk.INSTANCE.gtk_init(null, null);

            synchronized (lock) {
                lock.notifyAll();
            }

            main();
        }
    }, GtkMessageLoop.class.getSimpleName());

    public GtkMessageLoop() {
        synchronized (lock) {
            MessageLoop.start();
            try {
                lock.wait();
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        }
    }

    public static GtkMessageLoop loop = null;
    public static int count = 0;

    synchronized public static void inc() {
        if (count == 0) {
            loop = new GtkMessageLoop();
        }
        count++;
    }

    synchronized public static void dec() {
        count--;

        if (count == 0) {
            loop = null;
        }
    }

    synchronized public static void close() {
        LibGtk.INSTANCE.g_main_loop_quit(mainloop);
    }

    public static void main() {
        mainloop = LibGtk.INSTANCE.g_main_loop_new(null, false);
        context = LibGtk.INSTANCE.g_main_loop_get_context(mainloop);
        LibGtk.INSTANCE.g_main_loop_run(mainloop);
    }

    public static void invokeLater(final Runnable r) {
        LibGtk.INSTANCE.g_main_context_invoke(context, new GSourceFunc() {
            @Override
            public boolean invoke(Pointer data) {
                r.run();
                return false;
            }
        }, null);
    }

}
