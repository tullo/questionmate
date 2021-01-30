package main

import (
	"fmt"
	"os"

	"github.com/tullo/questionmate/adapter/repositories/file"
	"github.com/tullo/questionmate/adapter/repositories/parser"
)

func main() {
	fn := fmt.Sprintf("%s/src/github.com/tullo/questionmate/config/legacylab", os.Getenv("GOPATH"))
	r := file.NewQuestionRepository(fn, parser.QMParser{})
	fmt.Printf("%s", r)
}
