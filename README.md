# Crush

A simple, laid back, RESTful in-memory message queue.

[![Build Status](https://travis-ci.org/kcmerrill/crush.svg?branch=master)](https://travis-ci.org/kcmerrill/crush)

![crush](https://raw.githubusercontent.com/kcmerrill/queued/master/assets/crush.jpg)

## Binaries || Installation

[![MacOSX](https://raw.githubusercontent.com/kcmerrill/go-dist/master/assets/apple_logo.png "Mac OSX")](http://go-dist.kcmerrill.com/kcmerrill/crush/mac/amd64) [![Linux](https://raw.githubusercontent.com/kcmerrill/go-dist/master/assets/linux_logo.png "Linux")](http://go-dist.kcmerrill.com/kcmerrill/crush/linux/amd64)

via golang:

`$ go get -u github.com/kcmerrill/crush`

or via docker:

`$ docker run -P -d --name crush kcmerrill/crush`

## Quick Setup

Submitting a message is as simple as GET/POST request. If you only need the id, there is no need to post a body. However, posting a body will be the messages `value`. Retreiving a message is as simple as `GET /topicname`.

In the following example, we'll create a message(and topic too!). The topic is `nemo-characters` with a new message. `message.id = "crush"` and `message.value = "the turtle"`. Easy PZ.

### Submit message

`$ curl -X POST -d "the turtle" http://localhost:8000/nemo-characters/crush`

or

`$ curl -X GET http://localhost:8000/nemo-characters/nemo`

### Grab message

`$ curl -X GET http://localhost:8000/nemo-characters`

### Delete message

`$ curl -X DELETE http://localhost:8000/nemo-characters/nemo`
