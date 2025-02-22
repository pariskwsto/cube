package styles

var colorReset = "\033[0m"
var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorOrange = "\033[38;5;208m"
// var colorBlue = "\033[34m"
var colorPurple = "\033[35m"
var colorCyan = "\033[36m"
var colorWhite = "\033[37m"

// Command applies a cyan style to command text.
func Command(text string) string {
	return colorCyan + text + colorReset
}

// SubCommand applies a yellow style to sub-command text.
func SubCommand(text string) string {
	return colorYellow + text + colorReset
}

// Success applies a green style to command text.
func Success(text string) string {
	return colorGreen + text + colorReset
}

// Error applies a red style to error messages.
func Error(text string) string {
	return colorRed + text + colorReset
}

// Warning applies an orange style to warnings.
func Warning(text string) string {
	return colorOrange + text + colorReset
}

// Suggestion applies a purple style to suggestions.
func Suggestion(text string) string {
	return colorPurple + text + colorReset
}

// Instruction applies a white style to instructions.
func Instruction(text string) string {
	return colorWhite + text + colorReset
}
