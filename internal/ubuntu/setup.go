package ubuntu

import (
	"fmt"

	"github.com/pariskwsto/cube/internal/linux"
	"github.com/pariskwsto/cube/internal/styles"
)

// InitialSetupAsRoot runs the initial VPS setup for Ubuntu as root
func InitialSetupAsRoot() error {
	// Step 1: Update Linux system
	fmt.Println(styles.Command("> Starting system update..."))
	if err := linux.UpdateLinux(); err != nil {
		return fmt.Errorf("error updating system: %w", err)
	}
	fmt.Println(styles.Success("System update completed successfully!"))

	return nil
}

// StandardSetup runs the standard VPS setup for Ubuntu
func StandardSetup() error {
	// Step 1: Update Linux system
	fmt.Println(styles.Command("> Starting system update..."))
	if err := linux.UpdateLinux(); err != nil {
		return fmt.Errorf("error updating system: %w", err)
	}
	fmt.Println(styles.Success("System update completed successfully!"))

	return nil
}