package com.github.axet.desktop.os.linux.handle;

import java.util.Set;
import java.util.TreeSet;

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
            MessageLoop.setDaemon(false);
            MessageLoop.start();
            try {
                lock.wait();
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        }
    }

    public static GtkMessageLoop loop = null;
    public static Set<Integer> count = new TreeSet<Integer>();

    synchronized public static void inc(int id) {
        if (count.size() == 0) {
            loop = new GtkMessageLoop();
        }
        count.add(id);
    }

    synchronized public static void dec(int id) {
        count.remove(id);

        if (count.size() == 0) {
            close();
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

    public static void invokeLater(final GSourceFunc r, Pointer data) {
        LibGtk.INSTANCE.g_main_context_invoke(context, r, data);
    }

}
