#!/bin/bash
TYPE=$1

BIN_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

ARGS=()
for ((i=2;i<=$#;i++)); do
    ARGS+=( "${!i}" )
done

case $TYPE in

  Netty)
    ${BIN_DIR}/../java-netty/bin/start.sh "${ARGS[@]}"
    ;;

  *)
    echo -n "unknown"
    ;;
esac


