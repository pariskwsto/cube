package linux

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/pariskwsto/cube/internal/styles"
	"github.com/pariskwsto/cube/internal/utils"
)

// CommandsSuggestions prints useful linux commands suggestions
func CommandsSuggestions() {
	utils.PrintSubCommand("System Related Commands:")
	utils.PrintSuggestion("Check OS version: " + styles.Command("$ cat /etc/*-release"))
	utils.PrintSuggestion("Check users list: " + styles.Command("$ less /etc/passwd"))
	utils.PrintSubCommand("SSH keys Related Commands:")
	utils.PrintSuggestion("List SSH keys: " + styles.Command("$ ls -al ~/.ssh"))
	utils.PrintSuggestion("Start the ssh-agent: " + styles.Command(`$ eval "$(ssh-agent -s)"`))
	utils.PrintSuggestion("Add your SSH private key: " + styles.Command("$ ssh-add ~/.ssh/id_rsa"))
	utils.PrintSuggestion("Start ssh-agent and add SSH key: " + styles.Command(`$ eval "$(ssh-agent)" && ssh-add ~/.ssh/id_rsa`))
}

// StdinUsernamePrompt prompts the user to provide a valid username
func StdinUsernamePrompt(promptMsg string) (string, error) {
	label := "Please provide a valid username"
	if promptMsg != "" {
		label = promptMsg
	}

	utils.PrintSuggestion(label)
	
	const usernamePattern = `^[a-zA-Z0-9_-]+$`
	re := regexp.MustCompile(usernamePattern)

	validate := func(input string) error {
		trimmed := strings.TrimSpace(input)
		if len(trimmed) == 0 {
			return errors.New("username cannot be empty")
		}
		if !re.MatchString(trimmed) {
			return fmt.Errorf("invalid username: must match pattern %s", usernamePattern)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "username",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return "", err
	}

	return strings.TrimSpace(result), nil
}