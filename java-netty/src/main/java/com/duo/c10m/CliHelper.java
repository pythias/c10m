package com.duo.c10m;

import javafx.util.Pair;
import org.apache.commons.cli.*;

public class CliHelper {
    public static Pair<String, Integer> getServer(String[] args) {
        Options options = new Options();
        options.addRequiredOption("h", "host", true, "Server hostname (default: 0.0.0.0).");
        options.addRequiredOption("p", "port", true, "Server port (default: 100).");

        CommandLineParser parser = new DefaultParser();
        HelpFormatter formatter = new HelpFormatter();

        try {
            CommandLine cmd = parser.parse(options, args);
            String host = cmd.getOptionValue("host");
            Integer port = Integer.parseInt(cmd.getOptionValue("port"));
            return new Pair<String, Integer>(host, port);
        } catch (ParseException e) {
            System.out.println(e.getMessage());
            formatter.printHelp("utility-name", options);
            System.exit(1);
        }

        return null;
    }
}
