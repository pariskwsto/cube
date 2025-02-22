package cmd

import (
	"github.com/pariskwsto/cube/internal/linux"
	"github.com/spf13/cobra"
)

// linuxTipsCmd represents the linuxTips command
var linuxTipsCmd = &cobra.Command{
	Use:   "linux-tips",
	Short: "Display useful Linux command tips",
	Run: func(cmd *cobra.Command, args []string) {
		linux.CommandsSuggestions()
	},
}

func init() {
	rootCmd.AddCommand(linuxTipsCmd)
}
