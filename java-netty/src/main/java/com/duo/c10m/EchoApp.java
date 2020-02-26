package com.duo.c10m;

import com.duo.c10m.core.EchoServer;
import javafx.util.Pair;

public class EchoApp {
    public static void main( String[] args ) {
        Pair<String, Integer> server = CliHelper.getServer(args);
        new EchoServer().run(server.getKey(), server.getValue());
    }
}
