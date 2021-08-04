package main

import (
	"strings"
)


func wordcount(phrase string) int {
	return len(strings.Fields(phrase))
}
