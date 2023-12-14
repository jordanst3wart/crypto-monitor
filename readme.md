# TODO
I need to change the logging behaviour
```
docker run --log-opt mode=non-blocking --log-opt max-buffer-size=4m $container
```

# Setup
Install go version 18, or latest. 
See here:
https://golang.org/doc/install

See Makefile

Did stuff through this:
https://tutorialedge.net/golang/consuming-restful-api-with-go/

Todo document architecture

### Business Requirements
Send email of buy

buy lower on AU, and Coinfloor, exchange if global market is higher. The market should remain high. Compare to somewhat average global price.

Purge old orders.

harder: buy lower on Global exchange if global market is higher



### instance
ECR is free under 500MB
t3.nano using docker containers


https://github.com/golang-standards/project-layout

Latest errors (these are hopefully fixed as of 14th of Dec):
```
2023/12/07 19:08:17 Name: CoinCorner_BTC Error json: cannot unmarshal number into Go struct field CoinfloorTickerAndBitstamp.volume of type string Coin {      }
2023/12/07 19:08:17 Name: CoinCorner_ETH Error json: cannot unmarshal number into Go struct field CoinfloorTickerAndBitstamp.volume of type string Coin {      }
2023/12/07 19:08:17 ARBITRAGE!!! on bid: IndependentReserve_BCH, ask:BTCMarket_AUD_BCH at +Inf
2023/12/07 19:08:17 ARBITRAGE!!! on bid: Bitstamp_BCH, ask:BTCMarket_AUD_BCH at +Inf
```


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
nohup /home/ec2-user/crypto-monitor/bin/main >> /home/ec2-user/crypto-monitor/bin/main.log 2>&1 &
```

Check for arbitrage with:
```sh
cat /home/ec2-user/crypto-monitor/bin/main.log | grep ARBITRAGE | grep -v Inf
```

Use Renovate bot: https://github.com/renovatebot/renovate

https://go.dev/doc/effective_go

https://medium.com/@kdnotes/golang-naming-rules-and-conventions-8efeecd23b68

Kill existing process:
```sh
ps -ef | grep /home/ec2-user/crypto-monitor/bin/main
# kill 1292727 id
```
