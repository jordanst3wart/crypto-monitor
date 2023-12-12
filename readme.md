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

Arbitage message:
```
2019/07/06 15:02:55 ARBITAGE!!! on bid: ACX_AUD_BTC, ask:Coinjar_AUD_BTC at 1.0107867240319606
2019/07/06 15:27:30 ARBITAGE!!! on bid: Bitstamp_BTC, ask:Coinjar_AUD_BTC at 1.0102502596745062
2019/07/06 15:27:30 ARBITAGE!!! on bid: GEMINI_BTC, ask:Coinjar_AUD_BTC at 1.0106540654643947
2019/07/06 15:27:30 ARBITAGE!!! on bid: Coinfloor_BTC, ask:Coinjar_AUD_BTC at 1.0103487084978577
2019/07/06 15:27:30 ARBITAGE!!! on bid: Coinjar_AUD_ETH, ask:ACX_AUD_ETH at 1.0104761904761905
2019/07/06 21:23:01 ARBITAGE!!! on bid: ACX_AUD_LTC, ask:BTCMarket_AUD_LTC at 1.0105044510385757
2019/07/06 21:24:42 ARBITAGE!!! on bid: ACX_AUD_LTC, ask:BTCMarket_AUD_LTC at 1.0110444747936584
2019/07/06 21:24:42 ARBITAGE!!! on bid: ACX_AUD_LTC, ask:Bitstamp_LTC at 1.0110257609377127
2019/07/07 08:07:20 ARBITAGE!!! on bid: Coinjar_AUD_XRP, ask:Bitstamp_XRP at 1.011372811322571
2019/07/07 08:09:01 ARBITAGE!!! on bid: Coinjar_AUD_XRP, ask:Bitstamp_XRP at 1.0106270559886354
2019/07/07 08:10:43 ARBITAGE!!! on bid: Coinjar_AUD_XRP, ask:Bitstamp_XRP at 1.011372811322571
2019/07/07 08:18:23 ARBITAGE!!! on bid: Coinjar_AUD_XRP, ask:Bitstamp_XRP at 1.0117847338605295
```

Errors getting:
```
2019/07/06 04:50:12 Name: Bitstamp_ETH Error unexpected EOF
2019/07/06 05:19:49 Name: Bitstamp_BTC Error unexpected EOF
2019/07/06 06:13:04 Name: Bitstamp_BTC Error unexpected EOF
2019/07/06 06:24:07 Name: Bitstamp_BCH Error unexpected EOF
2019/07/06 06:35:10 Name: Bitstamp_LTC Error unexpected EOF
2019/07/06 06:51:17 Name: Bitstamp_LTC Error unexpected EOF
2019/07/06 07:36:08 Name: Bitstamp_BCH Error unexpected EOF
2019/07/06 08:17:34 Name: Bitstamp_LTC Error unexpected EOF
2019/07/06 08:35:22 Name: Bitstamp_XRP Error unexpected EOF
2019/07/06 09:48:51 Name: Bitstamp_BTC Error unexpected EOF
2019/07/06 10:33:40 Name: Bitstamp_BCH Error unexpected EOF
2019/07/06 10:33:40 Name: Bitstamp_ETH Error unexpected EOF
2019/07/06 13:09:50 Name: Bitstamp_BCH Error unexpected EOF
2019/07/06 15:24:07 Name: Bitstamp_XRP Error unexpected EOF
2019/07/06 17:02:56 Name: Bitstamp_LTC Error unexpected EOF
2019/07/06 18:51:53 Name: Bitstamp_BTC Error unexpected EOF
2019/07/06 19:41:44 Name: Bitstamp_BCH Error unexpected EOF
2019/07/07 00:26:15 Name: Bitstamp_ETH Error unexpected EOF
2019/07/07 03:04:09 Name: Bitstamp_ETH Error unexpected EOF
2019/07/07 04:24:23 Name: Bitstamp_BTC Error unexpected EOF
2019/07/07 04:24:23 Name: Bitstamp_XRP Error unexpected EOF
2019/07/07 04:24:23 Name: Bitstamp_ETH Error unexpected EOF
2019/07/07 05:14:15 Name: Bitstamp_BTC Error unexpected EOF
2019/07/07 07:23:28 Name: Bitstamp_BCH Error unexpected EOF
2019/07/07 08:18:23 Name: Bitstamp_BTC Error unexpected EOF
2019/07/07 09:43:42 Name: Bitstamp_ETH Error unexpected EOF
```


https://github.com/golang-standards/project-layout

Latest errors:
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


ask:BTCMarket_AUD_BCH seems to be zero and result in +Int when dividing
This errors as well:
2023/12/10 04:45:33 Name: CoinCorner_ETH Error json: cannot unmarshal number into Go struct field CoinfloorTickerAndBitstamp.volume of type string Coin {      }
2023/12/10 04:45:33 Name: CoinCorner_BTC Error json: cannot unmarshal number into Go struct field CoinfloorTickerAndBitstamp.volume of type string Coin {      }


Check for arbitrage with:
```sh
cat /home/ec2-user/crypto-monitor/bin/main.log | grep ARBITRAGE | grep -v Inf
```


Use Renovate bot: https://github.com/renovatebot/renovate
