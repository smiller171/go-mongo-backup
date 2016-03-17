package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	var output string
	output = "MongoDumpServer v0.1"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(output); err != nil {
		log.Println("Failed", err)
	}
}

func dumpCreate(w http.ResponseWriter, r *http.Request) {
	var target dumpTarget

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Println("Failed", err)
	}
	if err := r.Body.Close(); err != nil {
		log.Println("Failed", err)
	}
	if err := json.Unmarshal(body, &target); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(422) //unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Println("Failed", err)
		}
	}

	d := dumpStart(target)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if d.Result == "success" {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(500)
	}
	if err := json.NewEncoder(w).Encode(d); err != nil {
		log.Println("Failed to encode json", err)
	}

}
