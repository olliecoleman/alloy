package cmd

import (
	"log"
	"os"

	"github.com/olliecoleman/alloy/app/router"
	"github.com/spf13/cobra"
)

var production bool

func init() {
	RootCmd.AddCommand(serverCmd)
	serverCmd.Flags().BoolVarP(&production, "production", "p", false, "Run in production env?")
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		if production {
			log.Println("Running webpack -p...")
			if _, err := os.Stat("webpack.config.js"); err == nil {
				err := runCmd("npm", "run", "dist")
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		log.Println("Starting server...")
		err := router.New()
		if err != nil {
			log.Fatal(err)
		}
	},
}
