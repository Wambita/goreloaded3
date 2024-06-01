package goreloaded

import "testing"

func TestHexToDec(t *testing.T) {
	tests := []struct {
		name     string
		hexstr   string
		expected string
	}{
		{name: "normal input", hexstr: "A", expected: "10"},
		{name: "double digits", hexstr: "1A", expected: "26"},
		{name: "tripple digits", hexstr: "E8B", expected: "3723"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := HexToDec(test.hexstr)
			if output != test.expected {
				t.Errorf("HexToDec(%v) \nExpected: %v \nGot:%v", test.hexstr, test.expected, output)
			}
		})
	}
}
