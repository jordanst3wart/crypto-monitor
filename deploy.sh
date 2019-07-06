#!/usr/bin/env bash



#ssh -i ~/.ssh/jordan-personal-aws.pem ec2-user@13.239.29.238
# stop existing process with kill
scp -i ~/.ssh/jordan-personal-aws.pem  bin/main ec2-user@13.239.29.238:/home/ec2-user
# ./main & # on server to run in background
