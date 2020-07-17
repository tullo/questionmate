package main

import (
	"fmt"
	"github.com/rwirdemann/questionmate/adapter/repositories/file"
	"os"
)

func main() {
	fn := fmt.Sprintf("%s/src/github.com/rwirdemann/questionmate/config/legacylab", os.Getenv("GOPATH"))
	r := file.NewQuestionRepository(fn)
	fmt.Printf("%s", r)
}
