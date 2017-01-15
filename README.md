![queued](https://raw.githubusercontent.com/kcmerrill/queued/master/assets/queued.png)

## Installation
`$ go get -u github.com/kcmerrill/queued`

or via docker:

`$ docker run -P -d kcmerrill/queued`

## Summary
Queued is just a simple in memory queue system to be used in simple applications.

## Instructions
Create a topic `heros` with a message id of `superman`. The value of given message is `clark kent`:

`$ curl -X POST -d "clark kent" http://localhost:8000/heros/superman`

Grab a message off the topic `heros`(not FIFO):

`$ curl -X GET http://localhost:8000/heros`
