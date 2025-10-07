package input

import (
	"mapp/internal/apps"

	"github.com/charmbracelet/huh"
)

func AppPicker(list []apps.AppResult) string {
	var selection string

	options := getOptions(list)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Installed Apps").
				Options(options...).
				Value(&selection),
		),
	)

	form.Run()

	return selection
}

func getOptions(list []apps.AppResult) []huh.Option[string] {
	var options []huh.Option[string]

	for _, app := range list {
		option := huh.NewOption(app.Name, app.Path)
		options = append(options, option)
	}

	return options
}
