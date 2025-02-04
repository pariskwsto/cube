/*
Copyright © 2025 Paris Kostopoulos <pariskwsto@gmail.com>
Licensed under the MIT License. See LICENSE file for details.
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// updateLinuxSystemCmd represents the updateLinuxSystem command
var updateLinuxSystemCmd = &cobra.Command{
	Use:   "update-linux-system",
	Short: "Updates the Linux system packages",
	Long: `Runs system updates using apt or other package managers, depending on the OS.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := updateLinux()
		if err != nil {
			fmt.Println("Error updating system:", err)
			os.Exit(1)
		}
		fmt.Println("System update completed successfully!")
	},
}

// updateLinux executes the necessary commands to update the Linux system
func updateLinux() error {
	fmt.Println("Updating package lists...")
	if err := aptUpdateCmd(); err != nil {
		return err
	}

	fmt.Println("Upgrading installed packages...")
	if err := aptUpgradeCmd(); err != nil {
		return err
	}

	fmt.Println("Removing unnecessary packages...")
	return aptAutoremoveCmd()
}

// aptUpdateCmd executes `sudo apt update`
func aptUpdateCmd() error {
	cmd := exec.Command("sudo", "apt", "update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// aptUpgradeCmd executes `sudo apt upgrade -y`
func aptUpgradeCmd() error {
	cmd := exec.Command("sudo", "apt", "upgrade", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// aptAutoremoveCmd executes `sudo apt autoremove -y`
func aptAutoremoveCmd() error {
	cmd := exec.Command("sudo", "apt", "autoremove", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// init initializes the cobra updateLinuxSystem command
func init() {
	rootCmd.AddCommand(updateLinuxSystemCmd)
}
