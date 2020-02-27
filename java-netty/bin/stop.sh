#!/bin/bash
#
BIN_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
for pid in `cat ${BIN_DIR}/.pids`; do
    kill $pid && echo "Killing ${pid}"
done

: > ${BIN_DIR}/.pids
