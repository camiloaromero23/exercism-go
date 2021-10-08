package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var stars string
	for i := 0; i < numStarsPerLine; i++ {
		stars += "*"
	}
	return fmt.Sprintf("%s\n%s\n%s", stars, welcomeMsg, stars)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	msg := strings.Trim(oldMsg, "*\n")
	msg = strings.TrimLeft(msg, " ")
	msg = strings.TrimRight(msg, " ")
	return msg
}
