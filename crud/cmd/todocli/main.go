package main

import (
	"fmt"
	"os"

	"sadlil.com/samples/crud/pkg/todocli"
)

func main() {
	cmd := todocli.NewRootCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stdout, "Failed to run todocli cmd, reason: %v", err)
		os.Exit(1)
	}
}
