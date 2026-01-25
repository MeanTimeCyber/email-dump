package dumper

import (
	"fmt"
	"net/mail"
	"os"

	"github.com/hexiosec/email-parse/emlparse"
	"github.com/markkurossi/tabulate"
)

func DumpEML(filePath string) error {
	eml, err := emlparse.ReadFromFile(filePath)

	if err != nil {
		return fmt.Errorf("Error reading MSG file: %s", err.Error())
	}

	fmt.Println("Message processed successfully.")

	err = PrintHeaders(eml.Message.Header)

	if err != nil {
		return fmt.Errorf("Error printing EML summary: %s", err.Error())
	}

	return nil
}

func PrintHeaders(header mail.Header) error {
	mainKeys := []string{
		"From",
		"Return-Path",
		"To",
		"Subject",
		"Date",
		"Message-Id",
		//"Authentication-Results",
	}

	tab := tabulate.New(tabulate.Unicode)

	for _, key := range mainKeys {
		value, ok := header[key]

		if ok {
			row := tab.Row()
			row.Column(key)
			row.Column(value[0])
		}
	}

	tab.Print(os.Stdout)

	return nil
}
