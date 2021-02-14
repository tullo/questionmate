package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpadapter "github.com/tullo/questionmate/adapter/driver/http"
	"github.com/tullo/questionmate/adapter/repositories/file"
	"github.com/tullo/questionmate/adapter/repositories/parser"
	"github.com/tullo/questionmate/usecase"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fn := fmt.Sprintf("%s/config", wd)
	dir := flag.String("directory", fn, "the directory")
	flag.Parse()

	// 1. Instantiate repository.
	repo := file.NewQuestionnaireRepository(*dir+"/coma", parser.YAMLParser{})

	// 2. Instantiate hexagons.
	questionnaire := usecase.NewGetQuestionnaire()
	questionnaire.Repositories["coma"] = repo

	hexagon := usecase.NextQuestion{QR: repo}
	evaluator := usecase.Assessment{QR: repo}

	// 3. Instantiate handlers.
	qh := httpadapter.NewGetQuestionnaireHandler(questionnaire)
	next := httpadapter.NewNextQuestionHandler(hexagon)
	ah := httpadapter.NewAssessmentHandler(evaluator)

	r := mux.NewRouter()
	r.HandleFunc("/{questionnaire}", qh).Methods("GET")
	r.HandleFunc("/{questionnaire}/questions", next).Methods("POST")
	r.HandleFunc("/{questionnaire}/assessment", ah).Methods("POST")
	log.Printf("Service listening on http://localhost:8080...")
	handler := cors.AllowAll().Handler(r)
	_ = http.ListenAndServe(":8080", handler)
}
