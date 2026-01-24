package input

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/gabriel-vasile/mimetype"
	"github.com/markkurossi/tabulate"
)

const (
	OutlookMsgMime = "application/vnd.ms-outlook"
	EMLMime        = "message/rfc822"
)

// FileExists checks if a file exists at the given path.
func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

type FileDetails struct {
	Size      int64
	Name      string
	SHA256    string
	MimeType  string
	Extension string
}

// GetFileDetails retrieves details about the file at the given path.
func GetFileDetails(filePath string) (*FileDetails, error) {
	var details FileDetails

	stats, err := os.Stat(filePath)

	if err != nil {
		return nil, err
	}

	details.Size = stats.Size()
	details.Name = stats.Name()
	details.Extension = path.Ext(stats.Name())

	// get SHA256
	hash, err := getSHA256(filePath)

	if err != nil {
		return nil, err
	}

	details.SHA256 = hash

	// get MimeType
	mtype, err := GetFileMimeType(filePath)

	if err != nil {
		return nil, err
	}

	details.MimeType = mtype

	return &details, nil
}

// getSHA256 computes the SHA256 hash of the file at the given path.
func getSHA256(filePath string) (string, error) {
	f, err := os.Open(filePath)

	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()

	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// PrettyPrint prints the file details in a tabular format.
func (fd FileDetails) PrettyPrint() {
	tab := tabulate.New(tabulate.Unicode)
	err := tabulate.Reflect(tab, 0, nil, fd)

	if err != nil {
		fmt.Printf("Error creating table: %v\n", err)
		return
	}

	tab.Print(os.Stdout)
}

func GetFileMimeType(filePath string) (string, error) {
	mtype, err := mimetype.DetectFile(filePath)
	return mtype.String(), err

}
