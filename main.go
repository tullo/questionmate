package main

import (
	"fmt"
	"github.com/rwirdemann/questionmate/adapter/qm"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fn := fmt.Sprintf("%s/src/github.com/rwirdemann/questionmate/config/%s", os.Getenv("GOPATH"), "legacylab.qm")
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	var store qm.QuestionStore
	q := store.LoadQuestions(data)
	fmt.Printf(q.String())
}
