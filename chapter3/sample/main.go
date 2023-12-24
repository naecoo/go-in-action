package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Animal struct {
	Name string `json:"name"`
}

const filename = "data.json"

func main() {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	var animals []*Animal
	err = json.NewDecoder(file).Decode(&animals)
	if err != nil {
		panic(err)
	}

	for _, animal := range animals {
		fmt.Printf("animal name is %s\n", animal.Name)
	}
}
