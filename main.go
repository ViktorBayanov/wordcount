// wordcount tool counts number of words in a string.
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	phrase := readInput()
	fmt.Println(wordcount(phrase))
}

// readInput reads phrase
// from command line arguments and returns them.
func readInput() (phrase string) {
	flag.Parse()
	phrase = flag.Args()[0]
	return phrase
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("wordcount:", err)
	os.Exit(1)
}
