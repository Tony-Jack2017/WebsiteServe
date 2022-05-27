package cmd

import (
	"Serve/tools/color_print"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "apis",
	Short: "apis is shortcut of admin serve generator",
	Long: "apis is command tool for admin-serve",
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color_print.ColorPrintln(color_print.Fail(err))
		os.Exit(1)
	}
}