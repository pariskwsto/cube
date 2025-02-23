package cmd

import (
	"fmt"

	"github.com/pariskwsto/cube/internal/linux"
	"github.com/pariskwsto/cube/internal/utils"
	"github.com/spf13/cobra"
)

// createNewUserCmd represents the createNewUser command
var createNewUserCmd = &cobra.Command{
	Use:   "create-new-user",
	Short: "Create a new system user and grant admin privileges",
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintCommand("> Starting new system user creation...")
		
		username, err := linux.StdinUsernamePrompt("Enter new username for system: ")
		if err != nil || username == "" {
			utils.PrintError("Error reading username or no username provided")
			return
		}

		if err := linux.CreateAndConfigureUser(username); err != nil {
			utils.PrintError(fmt.Sprintf("Error creating new system user: %v", err))
			return
		}

		utils.PrintSuccess(fmt.Sprintf("User %s created successfully with admin privileges", username))
	},
}

func init() {
	rootCmd.AddCommand(createNewUserCmd)
}
