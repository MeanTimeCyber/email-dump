package dumper

import (
	"fmt"
	"os"

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

	return nil
}

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
	tab.Header("Property").SetAlign(tabulate.MR)
	tab.Header("Value").SetAlign(tabulate.MR)

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

func dumpAllProperties(props map[string]string) {
	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Property").SetAlign(tabulate.MR)
	tab.Header("Value").SetAlign(tabulate.MR)

	for key, value := range props {
		row := tab.Row()
		row.Column(key)
		row.Column(value)
	}

	tab.Print(os.Stdout)
}
