package main

import (
	"os"

	"github.com/boeing-ai-gateway/nah/pkg/deepcopy"
)

func main() {
	deepcopy.Deepcopy(os.Args[1:]...)
}
