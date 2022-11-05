package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

// func read() []Student {
// 	// loads array from toml
// 	// toml.DecodeFile()
// 	return nil
// }

func write(slice []Student) {
	if err := toml.NewEncoder(file).Encode(slice); err != nil {
		// failed to encode
		log.Fatal(err)
	}
}