package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version of the admin serve",
	Long: `This is temporary version of admin serve,
in future we will create more strong command tool, welcome to join us`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve version: 1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
