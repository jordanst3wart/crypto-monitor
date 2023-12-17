# TODO
I need to change the logging behaviour
```
docker run --log-opt mode=non-blocking --log-opt max-buffer-size=4m $container
```

# Setup
See Makefile

Did stuff through this:
https://tutorialedge.net/golang/consuming-restful-api-with-go/

### Business Requirements
Send email of buy

buy lower on AU, and Coinfloor, exchange if global market is higher. The market should remain high. Compare to somewhat average global price.

Purge old orders.

harder: buy lower on Global exchange if global market is higher


### instance
ECR is free under 500MB
t3.nano using docker containers


https://github.com/golang-standards/project-layout


docker image a tar file
```sh
docker build -t crypto:latest .
docker save crypto:latest > crypto-monitor-container.tar
```

On server (scp server):
```sh
docker load < /path/to/destination/crypto-monitor-container.tar
```
Need to output logs somehow...


# deploy

ssh
```sh
ssh -i ~/.ssh/python-watch-key.pem ec2-user@13.239.8.107
```


```sh
scp -i ~/.ssh/python-watch-key.pem -r bin ec2-user@13.239.8.107:/home/ec2-user/crypto-monitor
```

run with no hanging up, logging and in the background:
```sh
nohup /home/ec2-user/crypto-monitor/bin/main >> /home/ec2-user/crypto-monitor/bin/crypto.log 2>&1 &
```

Check for arbitrage with:
```sh
cat /home/ec2-user/crypto-monitor/bin/crypto.log | grep ARBITRAGE | grep -v Inf
```

Use Renovate bot: https://github.com/renovatebot/renovate

https://go.dev/doc/effective_go

https://medium.com/@kdnotes/golang-naming-rules-and-conventions-8efeecd23b68

Kill existing process:
```sh
ps -ef | grep /home/ec2-user/crypto-monitor/bin/main
# kill 1292727 id
```


# TODO make automated deploy and test
Build deploy:
```sh
# need to stop it on server with ps -ef | grep main
env GOOS=linux go build -ldflags="-s -w" -o bin/main main/functions.go main/main.go
scp -i ~/.ssh/python-watch-key.pem -r bin ec2-user@13.239.8.107:/home/ec2-user/crypto-monitor
```


TODO
- add amount ot arbitrage message
- add more crypto currencies to coinjar exchange
- add an automated deployment
- add automated tests

