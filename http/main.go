package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	adapter "github.com/rwirdemann/questionmate/adapter/driver/http"
	"github.com/rwirdemann/questionmate/adapter/repositories/file"
	"github.com/rwirdemann/questionmate/domain"
	"github.com/rwirdemann/questionmate/usecase"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	nextQuestion := usecase.NextQuestion{QuestionRepository: file.NewQuestionRepository("legacylab.qm")}
	nextQuestionHandler := adapter.MakeNextQuestionHandler(nextQuestion)
	r.HandleFunc("/questions", nextQuestionHandler).Methods("POST")
	r.HandleFunc("/evaluations", evaluationsHandler).Methods("POST")
	log.Printf("Service listening on http://localhost:8080...")
	handler := cors.AllowAll().Handler(r)
	_ = http.ListenAndServe(":8080", handler)
}

func evaluationsHandler(writer http.ResponseWriter, _ *http.Request) {
	evaluation := domain.Evaluation{}
	evaluation.Targets = append(evaluation.Targets, domain.Target{Text: "changeability", Score: 190})
	data, err := json.Marshal(evaluation)
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
