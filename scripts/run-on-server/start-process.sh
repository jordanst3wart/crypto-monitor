#!/usr/bin/env sh

nohup /home/ec2-user/crypto-monitor/bin/main >> /home/ec2-user/crypto-monitor/bin/stdout.log 2>> /home/ec2-user/crypto-monitor/bin/stderr.log &
echo $! > /home/ec2-user/crypto-monitor/bin/crypto.pid
