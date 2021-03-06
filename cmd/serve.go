package cmd

import (
	router2 "Serve/common/router"
	"Serve/tools/color_print"
	"Serve/tools/read_config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


var serveCmd = &cobra.Command{
	Use: "serve",
	Short: "serve command is start the admin serve",
	Run: func(cmd *cobra.Command, args []string) {
		color_print.ColorPrintln(color_print.Success("serve is starting..."))
		config := read_config.ReadConfig()
		router2.InitRouter("admin", config.Application.Port)
	},
}


func init() {
	serveCmd.Flags().String("port", "8000", "this is default port")
	viper.BindPFlag("application.port", serveCmd.Flags().Lookup("port"))
	rootCmd.AddCommand(serveCmd)
}
