package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rwirdemann/questionmate/usecase"
	"log"
	"net/http"
)

func MakeGetQuestionnaireHandler(u usecase.GetQuestionnaire) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		name := vars["questionnaire"]
		questionnaire, found := u.Process(name)
		if !found {
			writer.WriteHeader(http.StatusNotFound)
			return

		}
		data, err := json.Marshal(questionnaire)
		if err != nil {
			log.Printf("error: %s", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(data)
		if err != nil {
			log.Printf("error: %s", err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
	}
}
