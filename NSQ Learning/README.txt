NSQ Implemention



1. $ nsqlookupd &
2. $ nsqd --lookupd-tcp-address=127.0.0.1:4160 &
3. $ nsqadmin --lookupd-http-address=127.0.0.1:4161 &