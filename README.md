# Golly

Golly is a simple tool used to analyze endpoints availability worldwide.

[![Build Status](https://travis-ci.org/fernandonogueira/golly.svg?branch=master)](https://travis-ci.org/fernandonogueira/golly)

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/fernandonogueira/golly)

See the docs (http://docs.gollyapi.apiary.io/)
 
### Project motivation
 Golly is being developed to analyze availability of services worldwide enabling API consumers to identify service failures and route issues through the internet.  

### Sample requests

##### Sync analysis 
Path: _/syncAnalysis_

Headers: _Content-Type: application/json_

Body:
```
{
    "token":"anyGeneratedUniqueStringYouWantToSend",
    "httpMethod":"GET",
    "url":"http://google.com",
    "headers":{"Content-Type":"application/json"},
    "alwaysReturnBody":false
}
```

A request will be sent to google.com and the result will be returned synchronously.


##### Async Analysis
Path : _/analysis_

Headers: _Content-Type: application/json_

Body: 
```
{
    "token":"anyGeneratedUniqueStringYouWantToSend",
    "httpMethod":"GET",
    "url":"http://google.com",
    "alwaysReturnBody":false,
    "body":"",
    "headers":{"Content-Type":"application/json"},
    "webhookAddress":"http://someAddressToSendResponse/somePath"
}
```

A request will be sent to google.com and the result will be returned asynchronously through a webhook.