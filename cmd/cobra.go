package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "wind-admin",
	Short: "wind-admin is shortcut of admin serve generator",
	Long: "wind-admin is command tool for admin-serve",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}