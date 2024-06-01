package goreloaded

import (
	"fmt"
	"os"
	"strconv"
)

// Hexadecimal to decimal string conversion
func HexToDec(hexstr string) string {
	decStr, err := strconv.ParseInt(hexstr, 16, 64)
	if err != nil {
		fmt.Println("Error converting hexadecimal to decimal: ", err)
		os.Exit(0)
	}
	return strconv.FormatInt(decStr, 10)
}
