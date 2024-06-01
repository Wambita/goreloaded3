package goreloaded

import (
	"fmt"
	"os"
	"strconv"
)

func BinToDec(binstr string) string {
	decStr, err := strconv.ParseInt(binstr, 2, 64)
	if err != nil {
		fmt.Println("error converting binary string to decimal string:", err)
		os.Exit(0)
	}
	return strconv.FormatInt(decStr, 10)
}
