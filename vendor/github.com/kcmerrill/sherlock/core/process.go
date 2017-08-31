package core

import (
	"strconv"
	"strings"
)

// Process takes in a string to process
func (s *Sherlock) Process(msg, del string) {
	bits := strings.Split(msg, del)

	if len(bits) != 3 {
		// no need to go on ... dumb dumb
		return
	}

	if len(bits) == 2 {
		// oh snap! we have an entity and an event!
		s.E(bits[0]).Event(bits[1])
		return
	}

	// lets setup our vars
	entity := bits[0]
	property := bits[1]
	pType := "string"
	action := bits[2]
	value := ""

	if pBits := strings.Split(property, ":"); len(pBits) == 2 {
		property = pBits[0]
		pType = pBits[1]
	}

	if avBits := strings.Split(action, ":"); len(avBits) == 2 {
		action = avBits[0]
		value = avBits[1]
	} else {
		if action != "reset" {
			value = action
			action = "set"
		}
	}

	p := s.E(entity).NewProperty(property, pType)
	// if it already existed ... lets use it!
	pType = p.Type()

	switch action {
	case "reset":
		p.Reset()
	case "set":
		switch pType {
		case "bool":
			if strings.ToLower(value) == "true" || strings.ToLower(value) == "1" {
				p.Set(true)
			} else {
				p.Set(false)
			}
		case "int":
			if i, iErr := strconv.Atoi(value); iErr == nil {
				p.Set(i)
			}
		default:
			p.Set(value)
		}
	case "remove":
		switch pType {
		case "int":
			if i, iErr := strconv.Atoi(value); iErr == nil {
				p.Remove(i)
			}
		default:
			p.Remove(value)
		}
	case "add":
		switch pType {
		case "int":
			if i, iErr := strconv.Atoi(value); iErr == nil {
				p.Add(i)
			}
		default:
			p.Add(value)
		}
	}
}
