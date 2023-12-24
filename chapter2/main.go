package main

import (
	"log"
	"os"

	_ "github.com/naecoo/go-in-action/chapter2/matchers"
	"github.com/naecoo/go-in-action/chapter2/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
