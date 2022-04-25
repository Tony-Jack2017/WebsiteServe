package cmd

import (
	"Serve/app/router"
	"Serve/tools/color_print"
	"github.com/spf13/cobra"
)

var port string

var serveCmd = &cobra.Command{
	Use: "serve",
	Short: "serve command is start the admin serve",
	Run: func(cmd *cobra.Command, args []string) {
		color_print.ColorPrintln(color_print.Success("serve is starting..."))
		router.InitRouter(port)
	},
}


func init() {
	serveCmd.Flags().StringVarP(&port, "port", "p", "8000", "this is default port")
	rootCmd.AddCommand(serveCmd)
}
