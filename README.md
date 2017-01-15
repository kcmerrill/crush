# Crush
A simple, laid back, rest based in-memory queue.

[![Build Status](https://travis-ci.org/kcmerrill/crush.svg?branch=master)](https://travis-ci.org/kcmerrill/crush)

![crush](https://raw.githubusercontent.com/kcmerrill/queued/master/assets/crush.jpg)

## Installation
`$ go get -u github.com/kcmerrill/crush`

or via docker:

`$ docker run -P -d --name crush kcmerrill/crush`

or via alfred:

`$ alfred kcmerrill/crush start`

## Features
 - Quick and easy setup
 - Submit and retrieve messages
 - (b) Create web based workers to be used with
 - (b) Remember the results for a given message id
 - (b) Submit messages in intervals(daily, weekly, hourly, etc ..)
 - (b) Submit messages with a time delay

(b) - in backlog

## Quick Setup
Submitting a message is as simple as GET/POST request. If you only need the id, there is no need to post a body. However, posting a body will be the messages `value`. Retreiving a message is as simple as `GET /topicname`.

In the following example, we'll create a message(and topic too) called `nemo-characters` with a new message. `message.id = "crush"` and `message.value = "the turtle"`. Easy PZ.

### Submit message
`$ curl -X POST -d "the turtle" http://localhost:8000/nemo-characters/crush`

### Grab message
`$ curl -X GET http://localhost:8000/nemo-characters`
