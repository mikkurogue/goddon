package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func (c *Configuration) UpdateConfiguration() {

	if c.GameVersion == "" {
		log.Fatal("No game version configured")
	}

	if c.AddonsDirectory == "" {
		log.Fatal("No addons directory found")
	}

}

func (c *Configuration) UpdateGameVersion(v, ad string) {

	if c.GameVersion != ad {
		c.AddonsDirectory = ad
	}

	if c.GameVersion != v {
		c.GameVersion = v
	}

	fmt.Println("Updated Configuration")
}

func Validate() bool {
	_, err := os.Stat("config.json")
	return !os.IsNotExist(err)
}

func ReadConfiguration() (*Configuration, error) {

	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	var c Configuration
	if json.Unmarshal(data, &c) != nil {
		return nil, err
	}

	return &c, nil
}

func DefaultConfig() {
	c := Configuration{
		GameVersion:     "retail",
		AddonsDirectory: "Z:/World of Warcraft/_retail_/interface/addons",
	}

	j, err := json.MarshalIndent(c, "", "")
	if err != nil {
		log.Fatal(err)
	}

	filename := "config.json"
	err = os.WriteFile(filename, j, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
