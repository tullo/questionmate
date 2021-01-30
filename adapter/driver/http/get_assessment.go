package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tullo/questionmate/domain"
	"github.com/tullo/questionmate/usecase"
)

func MakeAssessmentHandler(evaluator usecase.Assessment) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		b, err := ioutil.ReadAll(request.Body)
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

		evaluation := evaluator.GetAssessment(body.Answers)
		data, err := json.Marshal(evaluation)
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
			return
		}
	}
}
