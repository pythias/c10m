package com.duo.c10m;

import com.duo.c10m.core.EmptyServer;

public class EmptyApp {
    public static void main( String[] args ) {
        new EmptyServer().run("0.0.0.0", 9003);
    }
}
