package com.duo.c10m;

import com.duo.c10m.core.EchoServer;

public class EchoApp {
    public static void main( String[] args ) {
        new EchoServer().run("0.0.0.0", 9003);
    }
}
