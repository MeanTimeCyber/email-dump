package main

import (
	"flag"
	"fmt"

	"github.com/MeanTimeCyber/email-dump/input"
)

func main() {
	// Args
	var filePath string

	flag.StringVar(&filePath, "i", "", "Input email file")
	flag.Parse()

	if filePath == "" {
		fmt.Println("No input file provided. Use -i to specify the email file.")
		flag.Usage()
		return
	}

	// Process the email file
	fmt.Printf("Processing email file: %s\n", filePath)

	if !input.FileExists(filePath) {
		fmt.Printf("File does not exist: %s\n", filePath)
		return
	}

	// Get file details
	details, err := input.GetFileDetails(filePath)

	if err != nil {
		fmt.Printf("Error getting file details: %v\n", err)
		return
	}

	fmt.Printf("\nFile details:\n")
	details.PrettyPrint()
}
