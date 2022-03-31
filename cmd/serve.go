package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Short: "serve command is start the admin serve",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve is starting ....")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
