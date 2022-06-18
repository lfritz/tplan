package main

import (
	"fmt"
	"os"

	"github.com/lfritz/tplan/internal"
)

func main() {
	err := internal.Process(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
