package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"time"
)

/* Start our webserver */
func Web(port string) {
	log.WithFields(log.Fields{"port": port}).Info("Web interface starting ...")
	r := mux.NewRouter()
	/* Setup our routes */
	r.HandleFunc("/{topic}/{id}", wTopicId)
	r.HandleFunc("/{topic}", wNewMessage)
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func wNewMessage(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	msg := q.Message(vars["topic"])
	if msg != nil {
		/* woot woot */
		response.WriteHeader(http.StatusOK)
		fmt.Fprintf(response, msg.Format("json"))
	} else {
		/* couldn't find a message */
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(response, "{}")
	}
}

func wTopicId(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	switch {
	case request.Method == "DELETE":
		q.Complete(vars["topic"], vars["id"])
		response.WriteHeader(http.StatusOK)
		fmt.Fprintf(response, "")
		break
	case request.Method == "GET":
		fallthrough
	case request.Method == "POST":
		body, bodyerr := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		if bodyerr != nil {
			response.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(response, bodyerr.Error())
		} else {
			msg := q.NewMessage(vars["topic"], vars["id"], string(body))
			response.WriteHeader(http.StatusOK)
			fmt.Fprintf(response, msg.Format("json"))
		}
		break
	}
}
