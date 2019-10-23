package main

import (
	"os"

	"github.com/vhikd/mxget/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
