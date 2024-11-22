package main

import (
	"fmt"
	"goddons/installer"
)

func main() {

	fmt.Println("Installed addons")
	for _, item := range installer.InstalledAddons("retail", "Z:/World of Warcraft/_retail_/interface/addons") {
		fmt.Println(item)
	}

}
