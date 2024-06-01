package goreloaded

import (
	"fmt"
	"regexp"
	"strings"
)

func WordModifications(line string) string {
	// Hexadecimal modifications
	hexMatches := regexp.MustCompile(`([0-9A-Fa-f]+)\s?\(hex\)`).FindAllStringSubmatch(line, -1)
	for _, match := range hexMatches {
		hexStr := match[1]
		decStr := HexToDec(hexStr)
		line = strings.Replace(line, hexStr+" (hex)", decStr, 1)

	}

	// Binary modifications
	binMatches := regexp.MustCompile(`([01]+)\s?\(bin\)`).FindAllStringSubmatch(line, -1)
	for _, match := range binMatches {
		binStr := match[1]
		decStr := BinToDec(binStr)
		line = strings.Replace(line, binStr+" (bin)", decStr, 1)
		fmt.Println(line)
	}
	return line
}
