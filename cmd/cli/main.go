package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"dab.io/dice"
)

const (
	defaultCount = 4
	maxCount     = 100
)

func main() {
	count := defaultCount

	flag.Parse()
	if flag.NArg() > 0 {
		_, _ = fmt.Sscan(flag.Arg(0), &count)
	}
	if count > maxCount {
		fmt.Printf("ERROR: More than %d words are not supported.\n", maxCount)
		os.Exit(1)
	}

	fmt.Printf("%s\n", strings.Join(dice.Roll(count), " "))
}
