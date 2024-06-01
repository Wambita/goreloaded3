package goreloaded

import "testing"

func TestBinToDec(t *testing.T) {
	tests := []struct {
		name     string
		binstr   string
		expected string
	}{
		{name: "normal input", binstr: "10", expected: "2"},
		{name: "double digits", binstr: "1010", expected: "10"},
		{name: "tripple digits", binstr: "10000", expected: "16"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := BinToDec(test.binstr)
			if output != test.expected {
				t.Errorf("BinToDec(%v) \nExpected: %v \nGot:%v", test.binstr, test.expected, output)
			}
		})
	}
}
