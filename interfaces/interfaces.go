package interfaces

import (
	"fmt"
)

type devTool struct {
	name      string
	automated bool
}

type monitorTool struct {
	name      string
	automated bool
}

func (t devTool) isAutomated() {
	if t.automated == true {
		fmt.Println(t, " is automated")
	}
}

func (t monitorTool) isAutomated() {
	if t.automated == true {
		fmt.Println(t, " is automated")
	}
}

type automatedTool interface {
	// Any tool which supports isAutomated() method, is of type automatedTool
	isAutomated()
}

func tools(t automatedTool) {
	fmt.Println(t, " is a part of automatedTool interface")
}

// Print explores interfaces
func Print() {
	fmt.Println("Exploring Interfaces:")

	// Create a devtool
	github := devTool{
		name:      "Github",
		automated: true,
	}
	// Create a monitortool
	datadog := monitorTool{
		name:      "datadog",
		automated: true,
	}

	// Calling individual tools
	github.isAutomated()
	datadog.isAutomated()

	// Calling tools function of automatedTool interface
	tools(github)
	tools(datadog)
}

/* Output:
Exploring Interfaces:
{Github true}  is automated
{datadog true}  is automated
{Github true}  is a part of automatedTool interface
{datadog true}  is a part of automatedTool interface
*/
