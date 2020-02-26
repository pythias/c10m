package com.duo.c10m;

import com.duo.c10m.core.EmptyServer;

public class EmptyApp {
    public static void main( String[] args ) {
        CliHelper.Server server = CliHelper.getServer(args);
        new EmptyServer().run(server.getHost(), server.getPort());
    }
}
