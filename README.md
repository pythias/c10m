# C10M

[![Build Status](https://travis-ci.org/pythias/c10m.svg?branch=master)](https://travis-ci.org/pythias/c10m)

## 系统参数调整

```bash
# 端口范围设定
sysctl -w net.ipv4.ip_local_port_range='1024 65535'

# 端口快速回收
sysctl -w net.ipv4.tcp_tw_recycle=1
sysctl -w net.ipv4.tcp_tw_reuse=1
sysctl -w net.ipv4.tcp_timestamps=1

# 文件打开数量的处理
sysctl -w fs.file-max=10485760
ulimit -n 1048576
echo 'ulimit -n 1048576' >> ~/.bash_profile
```

## 脚本

```bash

bin/start-server.sh Netty 10.10.10.10 9000 9003 1G
bin/benchmark.sh Netty -h 10.10.10.10 -b 9000 -e 9003 -n 10000 -c 10 -m "Echo Message"

```