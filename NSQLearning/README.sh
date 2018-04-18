#!/usr/bin/env bash
# NSQ Implemention:

# Run below three commands before running NSQ Implementation.

# 1. nsqlookupd - a demon which manages NSQ topology information.
nohup nsqlookupd &

# 2. nsqd - a daemon that receives, queues, and delivers messages to clients.
# consumer and producers must be in connect with this tcp connection listing on port 4160.
nohup nsqd --lookupd-tcp-address=127.0.0.1:4160 &

# 3. nsqadmin -  Web UI to view aggregated cluster stats in realtime and perform various administrative tasks.
nohup nsqadmin --lookupd-http-address=127.0.0.1:4161 &