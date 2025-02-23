package ubuntu

import (
	"errors"
	"fmt"

	"github.com/pariskwsto/cube/internal/linux"
	"github.com/pariskwsto/cube/internal/utils"
)

// InitialSetupAsRoot runs the initial VPS setup for Ubuntu as root
func InitialSetupAsRoot() error {
	// Step 1: Update Linux system
	utils.PrintSubCommand("Starting system update...")

	if err := linux.UpdateLinux(); err != nil {
		return fmt.Errorf("error updating system: %w", err)
	}

	// Step 2: Create new system user
	utils.PrintSubCommand("Starting new system user creation...")

	username, err := linux.StdinUsernamePrompt("Enter new username for system: ")
	if err != nil || username == "" {
		errMsg := "error reading username or no username provided"
		utils.PrintError(errMsg)
		return errors.New(errMsg)
	}

	if err := linux.CreateAndConfigureUser(username); err != nil {
		errMsg := fmt.Sprintf("error creating new system user: %v", err)
		utils.PrintError(errMsg)
		return errors.New(errMsg)
	}

	return nil
}

// StandardSetup runs the standard VPS setup for Ubuntu
func StandardSetup() error {
	// Step 1: Update Linux system
	utils.PrintSubCommand("Starting system update...")

	if err := linux.UpdateLinux(); err != nil {
		return fmt.Errorf("error updating system: %w", err)
	}

	return nil
}