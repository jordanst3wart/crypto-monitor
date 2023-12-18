#!/usr/bin/env sh

# fail and exit if anything fails
set -e

ssh -i ~/.ssh/python-watch-key.pem ec2-user@13.239.8.107 tail -n 150 /home/ec2-user/crypto-monitor/bin/stdout.log
