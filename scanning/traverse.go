package scanning

import (
	"encoding/csv"
	"os"
)

func HandleRun() ([][]string, error) {
	var results [][]string

	return results, nil
}

func loadDefinitions() ([][]string, error) {

	var matrix [][]string
	// load file
	definitionsFile, err := os.Open("./.octane-definitions/full.csv")
	if err != nil {
		return matrix, err
	}
	defer definitionsFile.Close()
	// load definitions
	csvReader := csv.NewReader(definitionsFile)

	//load definitions
	matrix, err = csvReader.ReadAll()

	if err != nil {
		return matrix, err
	}

	return matrix, nil

}

func mapDirectories(cwd string) {

}
