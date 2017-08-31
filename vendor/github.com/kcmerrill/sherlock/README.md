# sherlock

Entity tracker. What you track is up to you. 

[![Build Status](https://travis-ci.org/kcmerrill/sherlock.svg?branch=master)](https://travis-ci.org/kcmerrill/sherlock) [![Join the chat at https://gitter.im/kcmerrill/sherlock](https://badges.gitter.im/kcmerrill/sherlock.svg)](https://gitter.im/kcmerrill/sherlock?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

![sherlock](https://raw.githubusercontent.com/kcmerrill/sherlock/master/assets/sherlock.jpg "sherlock")

## What is it

A lot of front end services enable tracking via curl requests to build up entities. For example, a user joins your site with an email address of `kcmerrill@gmail.com`. A new entity is created and off of that you can start to track things. What you deceide to track is completly up to you. You can increment counters, store strings, store lists, unique lists etc etc ... 

For now, sherlock is only in memory(eventually we will store entities in a database). More to come.

## Usage

```golang

s := sherlock.New() // create a new sherlock

// create NewProperties on the entity(date|string|int|bool|*list) *coming later
s.Entity("kcmerrill@gmail.com").NewProperty("username", "string").Set("themayor")
// lets create another entity property but with shorthand string
s.Entity("doesnotexist").S("str_does_not_exist").Set("some_value")
s.Entity("doesnotexist").I("i_does_not_exist").Set(10)

if name := s.Entity("kcmerrill@gmail.com").Property("username").String(); name != "themayor" {
    t.Fatalf("Expected 'themayor', Actual: '%s'", name)
}

// make sure the entity creation time isn't zero
if s.Entity("kcmerrill@gmail.com").Created().IsZero() {
    t.Fatalf("Created should not be a zero time.Time")
}

// lets play with the counter now
e := s.Entity("kcmerrill@gmail.com")
e.NewProperty("counter", "int").Set(1000)

if i := s.Entity("kcmerrill@gmail.com").Property("counter").Int(); i != 1000 {
    t.Fatalf("Was expecting 'counter' to be 1000")
}

// Add to it
e.Property("counter").Add(100)

if i := s.Entity("kcmerrill@gmail.com").Property("counter").Int(); i != 1100 {
    t.Fatalf("Was expecting 'counter' to be 1100")
}

```

## UDP

You can interact with sherlock via it's UDP server. By default, it will be listening on port `8081`, or you can use the `--udp-port` option when starting. 

To send in a property simply follow this structure: `entity|property|value`, `entity|property:(string|date|bool|int)|(set|reset|add|remove):value(if necessary)`.

A quick note. Once you define a property type once on that entity, you don't have to anytime thereafter.

Here are a few examples:

```bash

$ echo -n "kcmerrill|this.is.an.event" | nc -w 0 -u localhost 8081 # creates an event
$ echo -n "kcmerrill|counter:int|100" | nc -w 0 -u localhost 8081 # sets counter property(by default an int) to 100. Notice the :int
$ echo -n "kcmerrill|email:string|kcmerrill@gmail.com" | nc -w 0 -u localhost 8081 # notice no property type. String is the default
$ echo -n "kcmerrill|counter|add:1" | nc -w 0 -u localhost 8081 # adds 1 to our counter
$ echo -n "kcmerrill|counter|remove:1" | nc -w 0 -u localhost 8081 # removes 1 to our counter
$ echo -n "kcmerrill|counter|reset" | nc -w 0 -u localhost 8081 # resets our counter to 0
$ echo -n "kcmerrill|logged.in:bool|true" | nc -w 0 -u localhost 8081 # sets logged.in property which is a bool to true
$ echo -n "kcmerrill|logged.in|false" | nc -w 0 -u localhost 8081 # sets logged.in property which is a bool to true, notice how we don't need to use bool anymore?
$ echo -n "kcmerrill|logged.in|reset" | nc -w 0 -u localhost 8081 # sets logged.in to it's reset state(false)

```

## HTTP

You can also interact with sherlock via it's HTTP server. By default, it will be listening on port `80`, or you can use the `--web-port` option when starting.

Both UDP and HTTP interactions should look similiar

A quick note. Once you define a property type once on that entity, you don't have to anytime thereafter.

```bash

$ curl -X GET http://localhost/kcmerrill/my.event.here # Will create an event
$ curl -X GET http://localhost/kcmerrill/logged.in:bool/true # Will create a bool property and set it to true
$ curl -X GET http://localhost/kcmerrill/counter/100 # Will create an int property and set it to 100
$ curl -X GET http://localhost/kcmerrill/counter/add:1 # Will add 1 to our counter
$ curl -X GET http://localhost/kcmerrill/counter/remove:1 # Will remove 1 from our counter

```