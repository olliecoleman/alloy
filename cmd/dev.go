package cmd

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/markbates/refresh/refresh"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(devCmd)
}

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Start the development server",
	Long:  `This command will watch for changes to your .go & .html files and re-run your app as soon as a file is updated.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Thanks to github.com/gobuffalo/buffalo
		ctx := context.Background()
		ctx, cancelFunc := context.WithCancel(ctx)
		go func() {
			err := startDevServer(ctx)
			if err != nil {
				cancelFunc()
				log.Fatal(err)
			}
		}()

		go func() {
			err := startWebpack()
			if err != nil {
				cancelFunc()
				log.Fatal(err)
			}
		}()
		// wait for the ctx to finish
		<-ctx.Done()
	},
}

func startDevServer(ctx context.Context) error {
	log.Println("Starting dev server...")
	configFile := "refresh.yml"
	_, err := os.Stat(configFile)
	if err != nil {
		c := refresh.Configuration{
			AppRoot:            ".",
			IgnoredFolders:     []string{"vendor", "log", "logs", "tmp", "node_modules", "bin", "templates"},
			IncludedExtensions: []string{".go", ".html"},
			BuildPath:          os.TempDir(),
			BuildDelay:         time.Second,
			BinaryName:         "alloy-build",
			CommandFlags:       []string{"server"},
			CommandEnv:         []string{},
			EnableColors:       true,
		}
		c.Dump(configFile)
	}

	c := &refresh.Configuration{}
	err = c.Load(configFile)
	if err != nil {
		return err
	}

	r := refresh.NewWithContext(c, ctx)
	return r.Start()
}

func startWebpack() error {
	log.Println("Starting webpack...")
	if _, err := os.Stat("webpack.config.js"); err != nil {
		return nil
	}

	return runCmd("npm", "run", "watch")
}
