package goreloaded

import "testing"

// possible test cases

// This is redundant,
func TestOpenFile(t *testing.T) {
	inputfilePath := "sample.txt"
	outputfilePath := "result.txt"

	// test opening the file, return error if file cannot be opened
	t.Run("Opening existing file", func(t *testing.T) {
		_, err := OpenFile(inputfilePath, outputfilePath)
		if err != nil {
			t.Errorf("failed to open existing file: %v", err)
		}
	})

	// test opening non existent file
	inputfilePath = "nonexistentfile.txt"
	outputfilePath = "notfoun.txt"
	t.Run("Opening non existent file", func(t *testing.T) {
		_, err := OpenFile(inputfilePath, outputfilePath)
		if err == nil {
			t.Errorf("Expected error for a non existent file but found none")
		}
	})
	// test for empty file, return error if no empty file error is returned
	inputfilePath = "empty.txt"
	t.Run("Opening an empty file", func(t *testing.T) {
		_, err := OpenFile(inputfilePath, outputfilePath)
		if err == nil {
			t.Errorf("expected error for empty file but found none")
		}
	})
	// test for wrong extension if wrong extension does not return error, return error
	inputfilePath = "image.png"
	t.Run("Opening file with wrong extension", func(t *testing.T) {
		_, err := OpenFile(inputfilePath, outputfilePath)
		if err == nil {
			t.Errorf("expected error for file with wron extension, but found none")
		}
	})

	// test for wrong file name , return error if wrong file names are used
	inputfilePath = "empty.txt"
	t.Run("Opening file with wrong name", func(t *testing.T) {
		_, err := OpenFile(inputfilePath, outputfilePath)
		if err == nil {
			t.Errorf("expected error for file with wrong name , but found none")
		}
	})
}

// optimized function that eliminates redundancy
// create a struct and loop through each possible test case
func TestOpenFile2(t *testing.T) {
	testcases := []struct {
		name            string
		inputfilePath   string
		outputfilePath  string
		expectedsuccess bool
	}{
		{"Opening existing file", "sample.txt", "result.txt", true},
		{"Opening non existent file", "nonexistentfile.txt", "notfound.txt", false},
		{"Opening an empty file", "empty.txt", "result.txt", false},
		{"Opening file with wrong extension", "image.png", "image.png", false},
		{"Opening file with wrong name", "empty.txt", "empty.txt", false},
	}

	for _, test := range testcases {
		t.Run(test.name, func(t *testing.T) {
			_, err := OpenFile(test.inputfilePath, test.outputfilePath)
			if (err == nil) != test.expectedsuccess {
				t.Errorf("Test case %q failed: expected success=%v, got error: %v", test.name, test.expectedsuccess, err)
			}
		})
	}
}
