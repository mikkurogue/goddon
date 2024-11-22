package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	entries, err := os.ReadDir("Z:/World of Warcraft/_retail_/interface/addons")
	if err != nil {
		log.Fatal("Fatal error no wow addon folder found for retail")
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}

}
