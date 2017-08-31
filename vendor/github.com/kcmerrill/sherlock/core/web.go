package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Web starts our webserver on a given port
func (s *Sherlock) Web(port, auth string) {
	// set our auth token
	s.AuthToken = auth

	r := mux.NewRouter()

	// setup our routes
	r.HandleFunc("/_all", s.basicauth(s.WebEntityAll))
	r.HandleFunc("/{entity}", s.basicauth(s.WebEntity))
	r.HandleFunc("/{entity}/{property}/{action}", s.basicauth(s.WebProcess))
	r.HandleFunc("/{entity}/{event}", s.basicauth(s.WebEvent))

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

// WebEntityAll handler for web requests with new messages
func (s *Sherlock) WebEntityAll(response http.ResponseWriter, request *http.Request) {
	s.lock.Lock()
	j, err := json.Marshal(s.Entities)
	s.lock.Unlock()
	if err == nil {
		response.WriteHeader(http.StatusOK)
		fmt.Fprintf(response, string(j))
	} else {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "{}")
	}
}

// WebEntity handler for web requests with new messages
func (s *Sherlock) WebEntity(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	ignore := []string{"favicon.ico", "robot.txt"}

	for _, i := range ignore {
		if i == vars["entity"] {
			// no need to go on, just return
			return
		}
	}

	j, err := s.E(vars["entity"]).String()
	if err == nil {
		response.WriteHeader(http.StatusOK)
		fmt.Fprintf(response, j)
	} else {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "{}")
	}
}

// WebEvent handler for web events
func (s *Sherlock) WebEvent(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	s.E(vars["entity"]).Event(vars["event"])
	j, err := s.E(vars["entity"]).String()
	if err == nil {
		response.WriteHeader(http.StatusOK)
		fmt.Fprintf(response, j)
	} else {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "{}")
	}
}

// WebProcess handler for web requests with new messages
func (s *Sherlock) WebProcess(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	s.Process(vars["entity"]+"|"+vars["property"]+"|"+vars["action"], "|")

	j, err := s.E(vars["entity"]).String()
	if err == nil {
		response.WriteHeader(http.StatusOK)
		fmt.Fprintf(response, j)
	} else {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(response, "{}")
	}
}

func (s *Sherlock) basicauth(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		token, _, _ := r.BasicAuth()
		if s.AuthToken != "" && s.AuthToken != token {
			http.Error(w, `{"error": "unauthorized"}`, http.StatusUnauthorized)
			return
		}
		fn(w, r)
	}
}
