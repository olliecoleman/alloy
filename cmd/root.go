package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "alloy",
	Short: "Boilerplate for creating web apps in Go (golang)",
	Long:  `Alloy is a starter template for creating web applications using Go programming language. It does not aim to be a web framework but is instead a collection of useful libraries and packages that acts a sensible starting point.`,
}

// Execute runs the root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func runCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Start()
}
