package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type contactData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func submitContact(w http.ResponseWriter, r *http.Request) {
	data, err := marshalContactData(r)
	if err != nil {
		log.Println(err)
	}
	log.Println(data)
	ajaxResponse(w, map[string]string{"success": "true"})
}

func marshalContactData(r *http.Request) (*contactData, error) {
	t := &contactData{}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(t)
	if err != nil {
		return t, err
	}
	return t, nil
}
