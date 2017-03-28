package crush

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

// Web starts our webserver on a given port
func (q *Q) Web(port string) {
	log.WithFields(log.Fields{"port": port}).Info("Crush is starting ...")
	r := mux.NewRouter()

	// setup our routes
	r.HandleFunc("/{topic}/{id}", q.WebTopicID)
	r.HandleFunc("/{topic}", q.WebNewMessage)

	// set some defaults
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	// start serving
	log.Fatal(srv.ListenAndServe())
}

// WebNewMessage handler for web requests with new messages
func (q *Q) WebNewMessage(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	msg := q.Message(vars["topic"])
	if msg != nil {
		// woot woot! message found
		response.WriteHeader(http.StatusOK)
		fmt.Fprintf(response, msg.String())
	} else {
		// boo, couldn't find a message
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(response, "{}")
	}
}

// WebTopicID creates/deletes/gets a message based on a topic and id
func (q *Q) WebTopicID(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	get := request.URL.Query()
	switch {
	// we need to delete the message at the given id
	case request.Method == "DELETE":
		q.Complete(vars["topic"], vars["id"])
		response.WriteHeader(http.StatusOK)
		fmt.Fprintf(response, "")
		break
	// Create a new message(not really restful, but super useful)
	case request.Method == "GET":
		fallthrough
	// create a new message
	case request.Method == "POST":
		body, bodyerr := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		if bodyerr != nil {
			response.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(response, bodyerr.Error())
		} else {
			msg := NewMessage(vars["topic"], vars["id"], string(body))

			if flight, exists := get["flight"]; exists && len(get["flight"]) > 0 {
				// user set flight, so lets set it here
				msg.Flight = flight[0]
			}

			if attempts, exists := get["attempts"]; exists && len(get["attempts"]) > 0 {
				// user set attempts, so lets set it here
				msg.Attempts, _ = strconv.Atoi(attempts[0])
			}

			q.NewRawMessage(msg)
			response.WriteHeader(http.StatusOK)
			fmt.Fprintf(response, msg.String())
		}
		break
	}
}
