package main

import (
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

func write(s map[string]Student) {

	buf, err := toml.Marshal(s)

	if err != nil {
		log.Fatal(err)
	}

	err1 := os.WriteFile("studentDB.toml", buf, 0644)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func read() map[string]Student {
	s := make(map[string]Student)
	buf, err := os.ReadFile("studentDB.toml")
	if err != nil {
		log.Fatal(err)
	}
	err1 := toml.Unmarshal(buf, &s)
	if err1 != nil {
		log.Fatal(err1)
	}

	return s
}
