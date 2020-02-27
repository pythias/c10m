#!/bin/bash

echo $@

if [[ -z "$4" ]]; then
    echo "start: wrong number of arguments"
    echo "Usage: start [host] [port-begin] [port-end] [xms]"
    echo "Options are:"
    echo "    host        Server ip address"
    echo "    port-begin  Server port begins"
    echo "    port-end    Server port ends"
    echo "    xms         JVM Xms, Xmx"
    echo "Examples:"
    echo "    \"./start 10.10.10.10 9000 9003 1G\" Server bind from 10.10.10.10:9000 to 10.10.10.10:9003, each has 1G Xms/Xmx"
    exit 1
fi

HOST=$1
BEGINS=$2
ENDS=$3
SIZE=$4

BIN_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
OPTS="-Xms${SIZE} -Xmx${SIZE}"
JAR=${BIN_DIR}/../target/netty-server-1.0-SNAPSHOT.jar
MAIN_CLASS=com.duo.c10m.EchoApp

ulimit -n 1048576

for port in $(seq "$BEGINS" "$ENDS") ; do
    nohup java ${OPTS} -cp ${JAR} ${MAIN_CLASS} -h $HOST -p $port 2>&1 &
    echo $! >> ${BIN_DIR}/.pids
    echo "Run $HOST:$port at $!";
done
