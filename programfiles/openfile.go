package goreloaded

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// Function to open input file, create output file, read input file and store contents as a string
func OpenFile(inputfilePath string, outputfilePath string) (string, error) {
	// input file validation
	fileInfo, err := os.Stat(inputfilePath)
	if err != nil {
		return "", fmt.Errorf("error opening input file: %v", err)
	}

	if fileInfo.Size() == 0 {
		return "", fmt.Errorf("error: file exists but is empty")
	}

	if filepath.Ext(inputfilePath) != ".txt" {
		return "", fmt.Errorf("wrong file extension, use .txt files: %v", err)
	}

	// input file open
	inputfile, err := os.Open(inputfilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("error:input/sample file does not exist \n %v", err)
		} else {
			return "", fmt.Errorf("error: %v", err)
		}
	}
	defer inputfile.Close()

	// output file creation
	outputfile, err := os.Create(outputfilePath)
	if err != nil {
		return "", fmt.Errorf("error creating output file: \n%v", err)
	}
	defer outputfile.Close()

	// read contents of input file
	var line string
	scanner := bufio.NewScanner(inputfile)
	for scanner.Scan() {
		line = scanner.Text()
	}
	return line, nil
}
