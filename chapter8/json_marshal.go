package main

import (
	"encoding/json"
	"fmt"
)

type Other struct {
	Field string `json:"field"`
}
type Person struct {
	Name  string `json:"name"`
	Age   uint8  `json:"age"`
	Other Other  `json:"other"`
}

func main() {

	var p = Person{
		Name:  "ocean",
		Age:   25,
		Other: Other{"test"},
	}

	// var str, err = json.Marshal(p)
	var str, err = json.MarshalIndent(p, "", "    ")
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println(string(str))
	}

}
