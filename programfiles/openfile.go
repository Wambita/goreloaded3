package goreloaded

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// Function to open input file, create output file, read input file and store contents as a string
func OpenFile(inputfilePath string, outputfilePath string) (string, error) {
	// possible errors: broken file, wrong extension, empty, doesn't exist
	// store fileinfo to use to check stat about the file
	fileinfo, err := os.Stat(inputfilePath)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist, ensure sample.txt exists and is not empty")
	}
	// check on file extensions to be .txt only
	if filepath.Ext(inputfilePath) != ".txt" {
		return "", fmt.Errorf("wrong input file extension use .txt files only")
	}
	// name of input file to be "sample.txt"
	if fileinfo.Name() != "sample.txt" {
		return "", fmt.Errorf(`wrong input file name, please input "sample.txt" as the input file name`)
	}
	// size of the file not to be 0
	if fileinfo.Size() == 0 {
		return "", fmt.Errorf("file exists, but is empty. Please ensure sample.txt has text to be formatted")
	}
	// open the inputfile and store it in a variable , handle errors
	inputFile, err := os.Open(inputfilePath)
	if err != nil {
		return "", fmt.Errorf("error opening file")
	}
	defer inputFile.Close() // saves on memory

	// create output file, handle possible errors
	// file extension to be .txt
	if filepath.Ext(outputfilePath) != ".txt" {
		return "", fmt.Errorf("wrong output file extension, use .txt files only")
	}
	// name of file to be result.txt
	if outputfilePath != "result.txt" {
		return "", fmt.Errorf(`wrong output file name, please input "result.txt" as the output file name`)
	}
	ouputFile, err := os.Create(outputfilePath)
	if err != nil {
		return "", fmt.Errorf("error creating output file")
	}
	defer ouputFile.Close()

	// read the input file and store the input in a string variable
	var line string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line = scanner.Text()
	}
	return line, nil
}
