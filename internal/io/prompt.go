package io

import (
	"fmt"
	"strings"
)

func StringPrompt(prompt string) string {
	fmt.Print(prompt, " ")
	var scanline string
	fmt.Scanln(&scanline)

	return scanline
}

func YNPrompt(prompt string) bool {
	return strings.ToLower(StringPrompt(prompt)) == "y"
}
