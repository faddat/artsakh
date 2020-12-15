package main

import (
	"os"

	"github.com/faddat/artsakh/cmd/artsakhd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
