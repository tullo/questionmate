package main

import (
	"fmt"
	"github.com/rwirdemann/questionmate/adapter/repositories/file"
)

func main() {
	r := file.NewQuestionRepository("legacylab")
	fmt.Printf("%s", r)
}
