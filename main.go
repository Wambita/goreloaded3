package main

import (
	"fmt"
	"os"

	goreloaded "goreloaded/programfiles"
)

func main() {
	// Declare variables to store input and output file paths
	var inputfilePath string
	var outputfilePath string

	// check length of os args if none give the usage
	if len(os.Args) == 0 {
		fmt.Println("Usage go run . sample.txt result.txt")
		return
	}
	// if length is just one ie the program name, give the variable the name of the input file and output file ane ensure that the input file exists
	if len(os.Args) == 1 {
		inputfilePath = "sample.txt"
		outputfilePath = "result.txt"
	}
	// if both the input and output files have been specified, take the 2nd and 3rd args as the input and output files
	if len(os.Args) == 3 {
		inputfilePath = os.Args[1]
		outputfilePath = os.Args[2]
	}
	// if just the input file is given create the output file
	if len(os.Args) == 2 {
		inputfilePath = os.Args[1]
		outputfilePath = "result.txt"
		// if theres more than 3 arguments return an error message
	}
	if len(os.Args) > 3 {
		fmt.Println("Invalid number of arguments: \nusage: go run . sample.txt result.txt")
		return
	}
	inputfile, err := goreloaded.OpenFile(inputfilePath, outputfilePath)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Input File Path:", inputfilePath)
	fmt.Println("Output File Path:", outputfilePath)
	fmt.Println("inputFile:", inputfile)
}
