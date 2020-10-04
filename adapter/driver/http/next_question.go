package http

import (
	"encoding/json"
	"github.com/rwirdemann/questionmate/domain"
	"github.com/rwirdemann/questionmate/usecase"
	"io/ioutil"
	"log"
	"net/http"
)

func MakeNextQuestionHandler(questionReader usecase.QuestionReader) http.HandlerFunc {
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

		q, hasNext := questionReader.NextQuestion(body.Answers)
		if hasNext {
			data, err := json.Marshal(q)
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
		} else {
			writer.WriteHeader(http.StatusNoContent)
		}
	}
}
