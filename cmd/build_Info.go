package cmd

import (
	"fmt"

	"github.com/pariskwsto/cube/build"
	"github.com/spf13/cobra"
)

// buildInfoCmd represents the buildInfo command
var buildInfoCmd = &cobra.Command{
	Use:   "build-info",
	Short: "Print detailed build information of cube",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Build Version:\t", build.Version)
		fmt.Println("Build User:\t", build.User)
		fmt.Println("Build Time:\t", build.Time)
		fmt.Println("Git Commit:\t", build.GitCommit)
	},
}

func init() {
	rootCmd.AddCommand(buildInfoCmd)
}
