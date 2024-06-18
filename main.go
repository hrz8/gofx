package main

import (
	"log"
	"os"

	CMD "hrz8/gofx/cmd"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use: "gofx",
}

func main() {
	cmd.AddCommand(CMD.AppCmd)

	if err := cmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
