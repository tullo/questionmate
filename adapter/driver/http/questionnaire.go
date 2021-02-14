package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tullo/questionmate/usecase"
)

func NewGetQuestionnaireHandler(uc usecase.GetQuestionnaire) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["questionnaire"]
		q, ok := uc.Process(name)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return

		}
		data, err := json.Marshal(q)
		if err != nil {
			log.Printf("error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(data)
		if err != nil {
			log.Printf("error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
