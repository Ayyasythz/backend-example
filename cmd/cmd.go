package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"sagara_backend_test/cmd/server"
	"sagara_backend_test/lib/log"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Wardrobe",
		Short: "Wardrobe - Backend",
	}
)

func Execute() {
	log.SetFormatter("json")

	rootCmd.AddCommand(server.ServeHTTPCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal("Error: ", err.Error())
		os.Exit(-1)
	}
}
