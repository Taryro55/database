package main

import (
	// "fmt"
	"log"
	"os"
	"path"
	toml "github.com/BurntSushi/toml"
)

type Student struct {
	fName   string
	lName   string
	age     int
	citizen bool
	grade   int
}

func init() {
	var PATH, _ = os.Getwd()
	path := path.Join(PATH, "//studentDB.toml")

	// Creates a file if it doesnt exist
	if _, err := os.Stat(path); err != nil {
		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

}

func load() []Student {
	// loads array from toml
	// toml.DecodeFile()
	return nil
}

func printDB([]Student) {
	
}

func main() {
	db := load()
	printDB(db)
	// print the current values (starting values would be nothing)
	// have an add function 
		// asks for everything and adds that to a Student struct, then it adds it to an array, and to a toml
	// have a remove func
		// asks for id to remove and edits the toml
	// have a refresh func
		// re-reads the toml file.

	// IFF using raylib, have a click on an area to sort by that



}
