package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/husobee/vestigo"
)

func mainREST(addr string) {
	r := vestigo.NewRouter()
	r.Post("/info", SetInfo)

	log.Fatal(http.ListenAndServeTLS(addr, certFile, keyFile, r))

}

// SetInfo - Rest HTTP Handler
func SetInfo(w http.ResponseWriter, r *http.Request) {
	var (
		input    apiInput
		response apiResponse
	)

	// decode input
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&input)
	r.Body.Close()

	// validate input
	if err := validate(input); err != nil {
		response.Success = false
		response.Reason = err.Error()
		respBytes, _ := json.Marshal(response)

		w.WriteHeader(400)
		w.Write(respBytes)
		return
	}
	response.Success = true
	respBytes, _ := json.Marshal(response)

	w.WriteHeader(200)
	w.Write(respBytes)
}

type apiResponse struct {
	Success bool   `json:"success"`
	Reason  string `json:"reason,omitempty"`
}

type apiInput struct {
	Name   string `json:"name"`
	Age    int    `json:"int"`
	Height int    `json:"height"`
}

// Validate - implementation of Validatable
func (ai apiInput) Validate() error {
	var err validationErrors
	if ai.Name == "" {
		err = append(err, errors.New("Name must be present"))
	}
	if ai.Age <= 0 {
		err = append(err, errors.New("Age must be real"))
	}
	if ai.Height <= 0 {
		err = append(err, errors.New("Height must be real"))
	}
	if len(err) == 0 {
		return nil
	}
	return err
}
