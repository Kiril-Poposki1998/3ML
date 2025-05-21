package handlereq

import (
	"os"

	"github.com/charmbracelet/huh"
)

func createForm() {
	// Create a new form
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput[string]().Title("Project Path").Value(os.Getwd()),
		),
	)

	// Run the form
	form.Run()
}
