package msg

import "github.com/fatih/color"

func Error(err error) {
	color.Red(err.Error())
}

func MessageWithFormat(message string, value ...any) {
		c := color.New(color.FgCyan)
		c.Printf(message, value...)
}

func Message(message any) {
	c := color.New(color.FgCyan)
	c.Println(message)
}

func Prompt(prompt any, newLine bool) {
	c := color.New(color.FgBlue)
	if newLine {
		c.Println(prompt)
		return
	}
	c.Print(prompt)
}

func Success(message any) {
	success := color.New(color.FgGreen)
	success.Println(message)
}

func Data(text string, value any) {
	textColor := color.New(color.FgCyan)
	valueColor := color.New(color.FgBlue).Add(color.Bold)

	textColor.Print(text)
	valueColor.Println(value)
}
