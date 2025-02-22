package linux

import (
	"fmt"

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
