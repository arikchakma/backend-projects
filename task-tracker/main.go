package main

import (
	"github.com/arikchakma/backend/task-tracker/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
