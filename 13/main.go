package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var outp string
	outp += strings.Repeat("*", numStarsPerLine)
	outp += "\n" + welcomeMsg + "\n"
	outp += strings.Repeat("*", numStarsPerLine)
	return outp
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	var outp string
	for _, c := range oldMsg {
		if c == '*' {
			continue
		}
		outp += string(c)
	}
	return strings.TrimSpace(outp)
}

// message := `
// **************************
// *    BUY NOW, SAVE 10%   *
// **************************
// `

// CleanUpMessage(message)
// // => BUY NOW, SAVE 10%
