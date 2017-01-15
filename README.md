[![Build Status](https://travis-ci.org/kcmerrill/crush.svg?branch=master)](https://travis-ci.org/kcmerrill/crush)

![crush](https://raw.githubusercontent.com/kcmerrill/queued/master/assets/crush.jpg)

## Installation
`$ go get -u github.com/kcmerrill/crush`

or via docker:

`$ docker run -P -d --name crush kcmerrill/crush`

or via alfred:

`$ alfred kcmerrill/crush start`

## Summary
`Crush` is a simple, laid back, rest based in-memory queue.

## Features
 - Quick and easy setup
 - Submit and retrieve messages
 - (b) Create web based workers to be used with
 - (b) Remember the results for a given message id
 - (b) Submit messages in intervals(daily, weekly, hourly, etc ..)
 - (b) Submit messages with a time delay

(b) - in backlog

## Quick Setup

`$ curl -X POST -d "the turtle" http://localhost:8000/characters/crush`

Creates a topic(if it doesn't exist) `characters` with a message id of `crush`. The value of given message with id `crush` is `the turtle`. The body of the request is the value of the message. This can be anything you want.

`$ curl -X GET http://localhost:8000/characters`

Grabs a message off the topic `characters`(not FIFO).
