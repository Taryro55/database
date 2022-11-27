package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

// func open() *os.File {
// 	f, err := os.OpenFile(fPath, os.O_RDWR|os.O_APPEND, 0660);
// 	if err != nil {
// 		// failed to open the file
// 		log.Fatal(err)
// 	}
// 	return f
// }

func write(s []Student) {
	for _, v := range s {
		fmt.Println(v)
		buf, err := toml.Marshal(v)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(buf))
		err1 := os.WriteFile("studentDB.toml", buf, 0644)
		if err1 != nil {
			log.Fatal(err1)
		} 
	}
}

func load() Student {
	s := Student{}
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
