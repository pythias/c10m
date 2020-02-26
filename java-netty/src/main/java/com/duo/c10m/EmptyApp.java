package com.duo.c10m;

import com.duo.c10m.core.EmptyServer;
import javafx.util.Pair;

public class EmptyApp {
    public static void main( String[] args ) {
        Pair<String, Integer> server = CliHelper.getServer(args);
        new EmptyServer().run(server.getKey(), server.getValue());
    }
}
