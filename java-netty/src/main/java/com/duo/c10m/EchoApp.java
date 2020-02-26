package com.duo.c10m;

import com.duo.c10m.core.EchoServer;

public class EchoApp {
    public static void main( String[] args ) {
        CliHelper.Server server = CliHelper.getServer(args);
        new EchoServer().run(server.getHost(), server.getPort());
    }
}
