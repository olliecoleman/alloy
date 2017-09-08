package cmd

import (
	"crypto/rand"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// VERSION tracks the current version of this application
const VERSION = "1.0.0"

var keySize int

func init() {
	RootCmd.AddCommand(versionCmd)

	genKey.Flags().IntVarP(&keySize, "size", "s", 32, "Key size")
	RootCmd.AddCommand(genKey)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of alloy",
	Long:  `Not the most useful command in the toolbox. It simply prints the current version of the Alloy toolkit.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Alloy version: %s\n", VERSION)
	},
}

var genKey = &cobra.Command{
	Use:   "gen-key",
	Short: "Generate a random key to use with securecookies",
	Run: func(cmd *cobra.Command, args []string) {
		if keySize < 2 {
			fmt.Println("Please enter a keySize > 2")
			return
		}

		k := make([]byte, keySize/2)
		io.ReadFull(rand.Reader, k)
		fmt.Printf("%x\n", k)
	},
}
