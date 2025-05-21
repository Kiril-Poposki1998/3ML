package main

import (
	"github.com/charmbracelet/huh"
)

func main() {
	var selectedColors []string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().Title("Select your favorite color").Options(
				huh.NewOption("Red", "red"),
				huh.NewOption("Green", "green"),
				huh.NewOption("Blue", "blue"),
			).Value(&selectedColors),
		),
	)
	form.Run()
	for _, color := range selectedColors {
		println("Selected color:", color)
	}
	// Handle form submission or cancellation
	// You can access the submitted values using form.Values
	// form.Values["name"], form.Values["email"], form.Values["age"]
}
