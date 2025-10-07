package main

import (
	"encoding/json"
	"fmt"
	"mac-apps/internal/apps"
	"mac-apps/internal/input"
)

func main() {
	appList := apps.GetApps()

	selectedApp := input.AppPicker(appList)

	data := apps.GetInfo(selectedApp)

	json, _ := json.MarshalIndent(data, "", "\t")
	fmt.Print(string(json))
}
