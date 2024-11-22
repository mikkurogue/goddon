package installer

import (
	"log"
	"os"
	"slices"
)

func Install(addon, v, dir string) {
	// TODO: make this install an addon provided
	if isInstalled(addon, dir) {
		return
	}
}

func Update(addon, v, dir string) {
	// TODO: make this update an addon provided
	if isInstalled(addon, dir) {
		return
	}
}

func Remove(addon, dir string) {
	// TODO: make this remove an addon provided
	if isInstalled(addon, dir) {
		return
	}
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

// util to check if an addon is already installed.
func isInstalled(addon, dir string) bool {
	installed := InstalledAddons("retail", dir)
	return slices.Contains(installed, addon)
}
