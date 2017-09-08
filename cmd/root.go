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
	Long: `Alloy is a starter template for creating web applications. It is written using Go language (golang).
Alloy is an omakase of sorts. It brings together many useful libraries and 
tools that you might need to build (CRUD) web applications.`,
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
