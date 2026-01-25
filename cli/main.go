package main

import (
	"flag"
	"fmt"

	"github.com/MeanTimeCyber/email-dump/dumper"
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

	if details.MimeType == input.OutlookMsgMime {
		fmt.Printf("\nDetected Outlook MSG file.\n")
		err := dumper.DumpMsg(filePath)

		if err != nil {
			fmt.Printf("Error dumping MSG file: %s\n", err.Error())
		}

	} else if details.MimeType == input.EMLMime {
		fmt.Printf("\nDetected EML file.\n")
		err := dumper.DumpEML(filePath)

		if err != nil {
			fmt.Printf("Error dumping EML file: %s\n", err.Error())
		}
	} else {
		fmt.Printf("\nUnsupported file type: %s\n", details.MimeType)
	}

	fmt.Println("Fin.")
}
