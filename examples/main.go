package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use: "gofx",
}

func main() {
	cmd.AddCommand(AppCmd)

	if err := cmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
