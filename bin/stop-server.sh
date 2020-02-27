#!/bin/bash
TYPE=$1

BIN_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

case $TYPE in

  Netty)
    ${BIN_DIR}/../java-netty/bin/stop.sh
    ;;

  *)
    echo -n "Unknown server type."
    ;;
esac

${BIN_DIR}/../java-netty/bin/stop.sh
