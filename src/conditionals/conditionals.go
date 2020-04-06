package conditionals

import (
	"fmt"
)

// Print explore conditional statements in Go
func Print() {
	technology := "golang"

	// Switch Condition
	switch technology {
	case "golang", "python":
		fmt.Printf("Switch Condition: You selected: %s (right choice)\n", technology)
	case "java", "node", "javascript":
		fmt.Printf("Switch Condition: You selected: %s (may not be the right choice)\n", technology)
	default:
		fmt.Println("Switch Condition: Keep Exploring")
	}

	// Same thing in if-else
	if technology == "golang" {
		fmt.Printf("If Condition: You selected: %s again (right choice)\n", technology)
	} else {
		fmt.Println("If Condition: Keep Looking")
	}
}
