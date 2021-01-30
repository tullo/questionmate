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
	fn := fmt.Sprintf("%s/config/coma", wd)
	directoryPtr := flag.String("directory", fn, "the questions directory")
	flag.Parse()

	// 1. Instantiate the "I need to go out httpadapter"
	repositoryAdapter := file.NewQuestionRepository(*directoryPtr, parser.YAMLParser{})

	// 2. Instantiate the hexagons
	getQuestionnaire := usecase.NewGetQuestionnaire()
	getQuestionnaire.Repositories["coma"] = repositoryAdapter

	hexagon := usecase.NextQuestion{QuestionRepository: repositoryAdapter}
	evaluator := usecase.Assessment{QuestionRepository: repositoryAdapter}

	// 3. Instantiate the "I need to go in adapter"
	getQuestionnaireHttpAdapter := httpadapter.MakeGetQuestionnaireHandler(getQuestionnaire)
	nextQuestionHttpAdapter := httpadapter.MakeNextQuestionHandler(hexagon)
	evaluatorHttpAdapter := httpadapter.MakeAssessmentHandler(evaluator)

	r := mux.NewRouter()
	r.HandleFunc("/{questionnaire}", getQuestionnaireHttpAdapter).Methods("GET")
	r.HandleFunc("/{questionnaire}/questions", nextQuestionHttpAdapter).Methods("POST")
	r.HandleFunc("/{questionnaire}/assessment", evaluatorHttpAdapter).Methods("POST")
	log.Printf("Service listening on http://localhost:8080...")
	handler := cors.AllowAll().Handler(r)
	_ = http.ListenAndServe(":8080", handler)
}
