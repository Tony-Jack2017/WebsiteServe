package cmd

import (
	"Serve/tools/color_print"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version of the serve",
	Long: `This is temporary version of serve,
in future we will create more strong command tool, welcome to join us`,
	Run: func(cmd *cobra.Command, args []string) {
		color_print.ColorPrint(color_print.Info("serve version: "), color_print.Success("1.0.0"))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
