package cmd

import (
	"fmt"

	"github.com/pariskwsto/cube/internal/linux"
	"github.com/pariskwsto/cube/internal/styles"
	"github.com/spf13/cobra"
)

// createNewUserCmd represents the createNewUser command
var createNewUserCmd = &cobra.Command{
	Use:   "create-new-user",
	Short: "Create a new system user and grant admin privileges",
	Run: func(cmd *cobra.Command, args []string) {
		username, err := linux.StdinUsernamePrompt("Enter new username for system: ")
		if err != nil || username == "" {
			fmt.Println("Error reading username or no username provided")
			return
		}

		if err := linux.CreateAndConfigureUser(username); err != nil {
			fmt.Println(styles.Error(fmt.Sprintf("Error creating new system user: %v", err)))
			return
		}

		fmt.Println(styles.Success(fmt.Sprintf("User %s created successfully with admin privileges", username)))
	},
}

func init() {
	rootCmd.AddCommand(createNewUserCmd)
}
