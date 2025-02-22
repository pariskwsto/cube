package linux

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pariskwsto/cube/internal/styles"
)

// CreateNewUser creates a new linux system user
func CreateNewSystemUser(username string) error {
	fmt.Println(styles.SubCommand(fmt.Sprintf("Creating new system user: %s", username)))

	fmt.Println(styles.Command("$ adduser " + username))
	cmd := exec.Command("adduser", username)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

// GrantAdminPrivileges grants sudo privileges to a given user
func GrantAdminPrivileges(username string) error {
	fmt.Println(styles.SubCommand(fmt.Sprintf("Granting administrative privileges to user: %s", username)))

	fmt.Println(styles.Command("$ usermod -aG sudo " + username))
	cmd := exec.Command("usermod", "-aG", "sudo", username)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to grant admin privileges: %w", err)
	}

	return nil
}

// CreateAndConfigureUser wraps both user creation and admin privileges assignment
func CreateAndConfigureUser(username string) error {
	if err := CreateNewSystemUser(username); err != nil {
		return err
	}

	if err := GrantAdminPrivileges(username); err != nil {
		return err
	}

	return nil
}
