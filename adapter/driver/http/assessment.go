package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tullo/questionmate/domain"
	"github.com/tullo/questionmate/usecase"
)

func NewAssessmentHandler(uc usecase.Assessment) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var body struct {
			Answers []domain.Answer `json:"answers"`
		}
		if err := json.Unmarshal(b, &body); err != nil {
			log.Printf("error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ass := uc.GetAssessment(body.Answers)
		data, err := json.Marshal(ass)
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
			return
		}
	}
}
