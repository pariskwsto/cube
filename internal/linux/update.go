package linux

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pariskwsto/cube/internal/styles"
)

// updateLinux executes the necessary commands to update the Linux system
func UpdateLinux() error {
	fmt.Println(styles.Command("Updating package lists..."))
	if err := aptUpdateCmd(); err != nil {
		return err
	}

	fmt.Println(styles.Command("Upgrading installed packages..."))
	if err := aptUpgradeCmd(); err != nil {
		return err
	}

	fmt.Println(styles.Command("Removing unnecessary packages..."))
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