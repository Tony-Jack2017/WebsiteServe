package cmd

import (
	"Serve/app/admin/router"
	"Serve/tools/color_print"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Short: "serve command is start the admin serve",
	Run: func(cmd *cobra.Command, args []string) {
		color_print.ColorPrintln(color_print.Success("serve is starting..."))
		router.InitRouter()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
