package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rwirdemann/questionmate/adapter/qm"
	"github.com/rwirdemann/questionmate/domain"
	"github.com/rwirdemann/questionmate/usecase/nextquestion"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var nq nextquestion.UseCase

func main() {
	fn := fmt.Sprintf("%s/src/github.com/rwirdemann/questionmate/config/%s", os.Getenv("GOPATH"), "legacylab.qm")
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	nq = nextquestion.NewUseCase(qm.QuestionStore{}, data)

	r := mux.NewRouter()
	r.HandleFunc("/questions", nextQuestionHandler).Methods("POST")
	log.Printf("Service listening on http://localhost:8080...")
	handler := cors.AllowAll().Handler(r)
	_ = http.ListenAndServe(":8080", handler)
}

func nextQuestionHandler(writer http.ResponseWriter, request *http.Request) {
	b, err := ioutil.ReadAll(request.Body)
	log.Printf("body: %s", b)
	if err != nil {
		log.Printf("error: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var body struct {
		Answers []domain.Answer `json:"answers"`
	}
	if err := json.Unmarshal(b, &body); err != nil {
		log.Printf("error: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	q, _ := nq.NextQuestion(body.Answers)
	data, err := json.Marshal(q)
	if err != nil {
		log.Printf("error: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = writer.Write(data)
	if err != nil {
		log.Printf("error: %s", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
