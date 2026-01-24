package dumper

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/hexiosec/email-parse/msgparse"
	"github.com/markkurossi/tabulate"
)

func DumpMsg(filePath string) error {

	msg, err := msgparse.ReadMsgFile(filePath, false)

	if err != nil {
		return fmt.Errorf("Error reading MSG file: %s", err.Error())
	}

	fmt.Println("Message processed successfully.")

	//dumpAllProperties(msg.Properties)

	printMainProperties(msg.Properties)

	if len(msg.Attachments) > 0 {
		fmt.Printf("\nAttachments (%d):\n", len(msg.Attachments))

		err = summariseAttachments(msg.Attachments)
	}

	return nil
}

// printMainProperties prints the interesting properties of the MSG file in a tabular format.
func printMainProperties(props map[string]string) {
	mainKeys := []string{
		"Display name",
		"Email address",
		"Received by name",
		"Received by email",
		"Received Representing name",
		"Representing email",

		"Sender name",
		"Sent Representing email",
		"Sent Representing name",
		"Sender email",
		"Reply Recipient Names",

		"Subject",
		"Subject Normalized",
		"Topic",
		"MessageID",

		// "Message Headers",
	}

	tab := tabulate.New(tabulate.Unicode)

	for _, key := range mainKeys {
		value, ok := props[key]

		if ok {
			row := tab.Row()
			row.Column(key)
			row.Column(value)
		}
	}

	tab.Print(os.Stdout)
}

// TODO option to dump out attachments
// summariseAttachments prints a summary of the attachments in the MSG file.
func summariseAttachments(attachment []msgparse.Attachment) error {
	for i, att := range attachment {
		fmt.Printf("\nAttachment %d:\n", i+1)

		tab := tabulate.New(tabulate.Unicode)

		row := tab.Row()
		row.Column("Filename")
		row.Column(att.Filename)

		row = tab.Row()
		row.Column("Long Filename")
		row.Column(att.LongFilename)

		row = tab.Row()
		row.Column("Extension")
		row.Column(att.UnicodeExtension)

		row = tab.Row()
		row.Column("Mime Tag")
		row.Column(att.MimeTag)

		row = tab.Row()
		row.Column("Size")
		row.Column(humanize.Bytes(uint64(len(att.Bytes))))

		row = tab.Row()
		row.Column("SHA256")

		h := sha256.New()
		h.Write(att.Bytes)
		row.Column(fmt.Sprintf("%x", h.Sum(nil)))

		tab.Print(os.Stdout)
	}

	return nil
}

func dumpAllProperties(props map[string]string) {
	tab := tabulate.New(tabulate.Unicode)

	for key, value := range props {
		row := tab.Row()
		row.Column(key)
		row.Column(value)
	}

	tab.Print(os.Stdout)
}
