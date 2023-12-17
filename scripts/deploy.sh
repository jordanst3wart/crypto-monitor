#!/usr/bin/env sh

# fail and exit if anything fails
set -e

# copy scripts
scp -i ~/.ssh/python-watch-key.pem -r scripts/run-on-server ec2-user@13.239.8.107:/tmp

# kill old process
ssh -i ~/.ssh/python-watch-key.pem ec2-user@13.239.8.107 chmod +x /tmp/run-on-server/kill-process.sh
ssh -i ~/.ssh/python-watch-key.pem ec2-user@13.239.8.107 /tmp/run-on-server/kill-process.sh

# copy binary
scp -i ~/.ssh/python-watch-key.pem -r bin ec2-user@13.239.8.107:/home/ec2-user/crypto-monitor

# start new process
ssh -i ~/.ssh/python-watch-key.pem ec2-user@13.239.8.107 chmod +x /tmp/run-on-server/start-process.sh
ssh -i ~/.ssh/python-watch-key.pem ec2-user@13.239.8.107 /tmp/run-on-server/start-process.sh
