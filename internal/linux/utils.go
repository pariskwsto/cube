package linux

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/pariskwsto/cube/internal/styles"
)

// CommandsSuggestions prints useful linux commands suggestions
func CommandsSuggestions() {
	fmt.Println(styles.SubCommand("System Related Commands:"))
	fmt.Println(styles.Suggestion("Check OS version: " + styles.Command("$ cat /etc/*-release")))
	fmt.Println(styles.Suggestion("Check users list: " + styles.Command("$ less /etc/passwd")))
	fmt.Println(styles.SubCommand("SSH keys Related Commands:"))
	fmt.Println(styles.Suggestion("List SSH keys: " + styles.Command("$ ls -al ~/.ssh")))
	fmt.Println(styles.Suggestion("Start the ssh-agent: " + styles.Command(`$ eval "$(ssh-agent -s)"`)))
	fmt.Println(styles.Suggestion("Add your SSH private key: " + styles.Command("$ ssh-add ~/.ssh/id_rsa")))
	fmt.Println(styles.Suggestion("Start ssh-agent and add SSH key: " + styles.Command(`$ eval "$(ssh-agent)" && ssh-add ~/.ssh/id_rsa`)))
}

// StdinUsernamePrompt prompts the user to provide a valid username
func StdinUsernamePrompt(promptMsg string) (string, error) {
	label := "Please provide a valid username"
	if promptMsg != "" {
		label = promptMsg
	}

	fmt.Println(styles.Suggestion(label))
	
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