![crush](https://raw.githubusercontent.com/kcmerrill/queued/master/assets/crush.jpg)

## Installation
`$ go get -u github.com/kcmerrill/crush`

or via docker:

`$ docker run -P -d --name crush kcmerrill/crush`

## Summary
`Crush` is a simple, laid back, rest based and in-memory queue.

## Features
 - Quick and easy setup
 - Submit and retrieve messages
 - (b) Create web based workers to be used with
 - (b) Remember the results for a given message id
 - (b) Submit messages in intervals(daily, weekly, hourly, etc ..)
 - (b) Submit messages with a time delay

(b) - in backlog

## Quick Setup
Create a topic `characters` with a message id of `crush`. The value of given message with id `crush` is `the turtle`:

`$ curl -X POST -d "the turtle" http://localhost:8000/characters/crush`

Grab a message off the topic `characters`(not FIFO):

`$ curl -X GET http://localhost:8000/characters`
