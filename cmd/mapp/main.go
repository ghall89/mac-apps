package main

import (
	"fmt"
	"mapp/internal/apps"
	"mapp/internal/input"
)

func main() {
	appList := apps.GetApps()

	selectedApp := input.AppPicker(appList)

	data := apps.GetInfo(selectedApp)

	fmt.Println(data.CFBundleExecutable + " v" + data.CFBundleShortVersionString)
	fmt.Println(data.NSHumanReadableCopyright)
}
