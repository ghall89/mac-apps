package main

import (
	"fmt"
	"mac-apps/internal/apps"
)

func main() {
	appList := apps.GetApps()

	for _, e := range appList {
		fmt.Println(e.Name())
	}
}
