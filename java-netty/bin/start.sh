#!/bin/bash

HOST=$1
BEGINS=$2
ENDS=$3
SIZE=$4

BIN_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
OPTS="-Xms${SIZE}G -Xmx${SIZE}G"
JAR=${BIN_DIR}/../target/netty-server-1.0-SNAPSHOT.jar
MAIN_CLASS=com.duo.c10m.EchoApp

ulimit -n 1048576

for port in $(seq "$BEGINS" "$ENDS") ; do
    nohup java ${OPTS} -cp ${JAR} ${MAIN_CLASS} -h $HOST -p $port 2>&1 &
    echo "Run $HOST:$port at $!";
done
