package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpadapter "github.com/rwirdemann/questionmate/adapter/driver/http"
	"github.com/rwirdemann/questionmate/adapter/repositories/file"
	"github.com/rwirdemann/questionmate/usecase"
	"log"
	"net/http"
	"os"
)

func main() {
	fn := fmt.Sprintf("%s/src/github.com/rwirdemann/questionmate/config/legacylab-short", os.Getenv("GOPATH"))
	filenamePtr := flag.String("filename", fn, "the questions file")
	flag.Parse()

	// 1. Instantiate the "I need to go out httpadapter"
	repositoryAdapter := file.NewQuestionRepository(*filenamePtr)

	// 2. Instantiate the hexagon
	hexagon := usecase.NextQuestion{QuestionRepository: repositoryAdapter}

	// 3. Instantiate the "I need to go in adapter"
	httpAdapter := httpadapter.MakeNextQuestionHandler(hexagon)

	r := mux.NewRouter()
	r.HandleFunc("/questions", httpAdapter).Methods("POST")
	r.HandleFunc("/evaluations", httpadapter.MakeEvaluationsHandler()).Methods("POST")
	log.Printf("Service listening on http://localhost:8080...")
	handler := cors.AllowAll().Handler(r)
	_ = http.ListenAndServe(":8080", handler)
}
