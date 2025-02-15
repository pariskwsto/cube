package cmd

import (
	"fmt"
	"os"

	"github.com/pariskwsto/cube/internal/linux"
	"github.com/pariskwsto/cube/internal/styles"
	"github.com/spf13/cobra"
)

// updateLinuxCmd represents the updateLinuxCmd command
var updateLinuxCmd = &cobra.Command{
	Use:   "update-linux",
	Short: "Updates the Linux system packages",
	Long: `Runs system updates using apt or other package managers, depending on the OS.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := linux.UpdateLinux()
		if err != nil {
			fmt.Println(styles.Error(fmt.Sprintf("Error updating system: %v", err)))
			os.Exit(1)
		}
		fmt.Println(styles.Success("System update completed successfully!"))
	},
}

func init() {
	rootCmd.AddCommand(updateLinuxCmd)
}
