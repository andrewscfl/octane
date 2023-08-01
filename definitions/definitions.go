package definitions

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/xyproto/unzip"
)

func HandleDefinitionsLoading() error {
	fmt.Println("Checking for malware definitions... 🔍")
	if !fileExists("./.octane-definitions/full.csv") {
		fmt.Println("Definitions not found! Downloading... 📥")
		err := downloadDefinitionsZip()
		if err != nil {
			return err
		}
		err = unzip.Extract(".octane-definitions.zip", ".octane-definitions")
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Definitions already exist! 🎉")
	}
	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func downloadDefinitionsZip() error {

	// Download definitions from https://bazaar.abuse.ch/export/csv/full/
	resp, err := http.Get("https://bazaar.abuse.ch/export/csv/full/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// create file to store definitions
	out, err := os.Create(".octane-definitions.zip")
	if err != nil {
		return err
	}
	defer out.Close()

	// write the body to file
	io.Copy(out, resp.Body)

	return nil
}
