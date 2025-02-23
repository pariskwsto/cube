package linux

import (
	"os"
	"os/exec"

	"github.com/pariskwsto/cube/internal/utils"
)

// UpdateLinux executes the necessary commands to update the Linux system
func UpdateLinux() error {
	if err := aptUpdateCmd(); err != nil {
		return err
	}

	if err := aptUpgradeCmd(); err != nil {
		return err
	}

	return aptAutoremoveCmd()
}

// aptUpdateCmd executes `sudo apt update`
func aptUpdateCmd() error {
	utils.PrintCommand("$ sudo apt update")
	cmd := exec.Command("sudo", "apt", "update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// aptUpgradeCmd executes `sudo apt upgrade -y`
func aptUpgradeCmd() error {
	utils.PrintCommand("$ sudo apt upgrade -y")
	cmd := exec.Command("sudo", "apt", "upgrade", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// aptAutoremoveCmd executes `sudo apt autoremove -y`
func aptAutoremoveCmd() error {
	utils.PrintCommand("$ sudo apt autoremove -y")
	cmd := exec.Command("sudo", "apt", "autoremove", "-y")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}