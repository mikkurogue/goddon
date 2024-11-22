package installer

import (
	"log"
	"os"
)

func Install(addon, v, dir string) {
	// TODO: make this install an addon provided
}

func Update(addon, v, dir string) {
	// TODO: make this update an addon provided
}

func InstalledAddons(v, dir string) []string {
	entries, err := os.ReadDir("Z:/World of Warcraft/_retail_/interface/addons")
	if err != nil {
		log.Fatal("Fatal error no wow addon folder found for retail")
	}

	var names []string
	for _, e := range entries {
		names = append(names, e.Name())
	}

	return names
}
