package main

import (
	"encoding/json"
	"fmt"
	"os"
	"param-finder/pkg/scanner"
	"param-finder/pkg/utils"

	"github.com/akamensky/argparse"
)

func saveResults(results []scanner.ParsedResult, filePath string) error {

	// Create or truncate the file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	encoder.SetIndent("", "    ")

	// Encode and write the JSON data to the file
	if err := encoder.Encode(results); err != nil {
		return err
	} else {
	}

	fmt.Println("Saved JSON to", filePath)
	return nil
}

func main() {

	// Get script arguments
	parser := argparse.NewParser("param-finder", "Search for params in source-code")
	dirFile := parser.String("F", "filename", &argparse.Options{Required: true, Help: "File that contains the list of directories"})
	outFile := parser.String("O", "outfile", &argparse.Options{Required: true, Help: "Output file"})
	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	// Read file of directories that will be scanned.
	directories, err := utils.FileToList(*dirFile)

	if err != nil {
		panic(err)
	}

	// Scan over directories.
	for _, dir := range directories {
		results, err := scanner.Scan(dir)

		if err != nil {
			continue
		}

		parsedResult := scanner.Parse(results)

		saveResults(parsedResult, *outFile)

	}

}
