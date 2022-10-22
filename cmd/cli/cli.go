package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"dab.io/dice"
)

func main() {
	count := dice.DefaultCount

	flag.Parse()
	if flag.NArg() > 0 {
		_, _ = fmt.Sscan(flag.Arg(0), &count)
	}
	if count > dice.MaxCount {
		fmt.Printf("ERROR: More than %d words are not supported.\n", dice.MaxCount)
		os.Exit(1)
	}

	fmt.Printf("%s\n", strings.Join(dice.Roll(count), " "))
}
