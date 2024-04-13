package main

import (
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
	println("init() project initializing...")
}

func main() {
	println("main function")
}
