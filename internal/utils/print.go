package utils

import (
	"fmt"

	"github.com/pariskwsto/cube/internal/styles"
)

// PrintCommand prints a styled command message.
func PrintCommand(text string) {
	fmt.Println(styles.Command(text))
}

// PrintSubCommand prints a styled sub-command message.
func PrintSubCommand(text string) {
	fmt.Println(styles.SubCommand(text))
}

// PrintSuccess prints a styled success message.
func PrintSuccess(text string) {
	fmt.Println(styles.Success(text))
}

// PrintError prints a styled error message.
func PrintError(text string) {
	fmt.Println(styles.Error(text))
}

// PrintWarning prints a styled warning message.
func PrintWarning(text string) {
	fmt.Println(styles.Warning(text))
}

// PrintSuggestion prints a styled suggestion message.
func PrintSuggestion(text string) {
	fmt.Println(styles.Suggestion(text))
}

// PrintInstruction prints a styled instruction message.
func PrintInstruction(text string) {
	fmt.Println(styles.Instruction(text))
}
