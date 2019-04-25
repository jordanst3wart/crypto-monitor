
# Setup

```
# create from template
# trying without go dep
sls create --template aws-go
# install dependcies
go get github.com/aws/aws-lambda-go/events
go get github.com/aws/aws-lambda-go/lambda


```


Did stuff through this:
https://tutorialedge.net/golang/consuming-restful-api-with-go/


Todo document architecture







### Business Requirements
Send email of buy

buy lower on AU, and Coinfloor, exchange if global market is higher. The market should remain high. Compare to somewhat average global price.

Purge old orders.

harder: buy lower on Global exchange if global market is higher

