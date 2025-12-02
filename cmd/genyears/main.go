package main

import (
	"log"
	"os"

	yearsgen "advent-of-code-go/internal/yearsgen"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to determine working directory: %v", err)
	}
	if err := yearsgen.Generate(root); err != nil {
		log.Fatalf("failed to generate year registry: %v", err)
	}
}
