package core

import "testing"
import "log"

func TestStringProcessor(t *testing.T) {
	s := New(100)
	s.Process("kcmerrill|username:string|kcmerrill", "|")

	if s.E("kcmerrill").S("username").String() != "kcmerrill" {
		log.Fatalf("kcmerrill should be set as the username")
	}

	s.Process("kcmerrill|username:string|reset", "|")

	if s.E("kcmerrill").S("username").String() != "" {
		log.Fatalf("The username once reset, should be empty")
	}

	s.Process("kcmerrill|email|kcmerrill@gmail.com", "|")

	if s.E("kcmerrill").S("email").String() != "kcmerrill@gmail.com" {
		log.Fatalf("email should be a string by default")
	}
}

func TestIntProcessor(t *testing.T) {
	s := New(100)
	s.Process("kcmerrill|counter:int|add:1", "|")

	if s.E("kcmerrill").I("counter").Int() != 1 {
		log.Fatalf("counter should = 1")
	}

	s.Process("kcmerrill|counter|add:100", "|")

	if s.E("kcmerrill").I("counter").Int() != 101 {
		log.Fatalf("counter should = 101")
	}

	s.Process("kcmerrill|counter|remove:1", "|")

	if s.E("kcmerrill").I("counter").Int() != 100 {
		log.Fatalf("counter should = 100")
	}

	s.Process("kcmerrill|counter|42", "|")

	if s.E("kcmerrill").I("counter").Int() != 42 {
		log.Fatalf("counter should = 42")
	}

	s.Process("kcmerrill|counter|reset", "|")

	if s.E("kcmerrill").I("counter").Int() != 0 {
		log.Fatalf("counter should = 0")
	}
}

func TestBoolProcessor(t *testing.T) {
	s := New(100)

	s.Process("kcmerrill|logged.in:bool|true", "|")

	if s.E("kcmerrill").I("logged.in").Int() != 1 {
		log.Fatalf("logged.in was set to true :(")
	}

	s.Process("kcmerrill|logged.in:bool|false", "|")

	if s.E("kcmerrill").I("logged.in").Int() != 0 {
		log.Fatalf("logged.in should be false")
	}
}
