package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	adapter "github.com/rwirdemann/questionmate/adapter/driver/http"
	"github.com/rwirdemann/questionmate/adapter/repositories/file"
	"github.com/rwirdemann/questionmate/usecase"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	nextQuestion := usecase.NextQuestion{QuestionRepository: file.NewQuestionRepository("legacylab.qm")}
	r.HandleFunc("/questions", adapter.MakeNextQuestionHandler(nextQuestion)).Methods("POST")
	r.HandleFunc("/evaluations", adapter.MakeEvaluationsHandler()).Methods("POST")
	log.Printf("Service listening on http://localhost:8080...")
	handler := cors.AllowAll().Handler(r)
	_ = http.ListenAndServe(":8080", handler)
}
