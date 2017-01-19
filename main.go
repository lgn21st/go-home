package main

import (
	"os"

	"github.com/lgn21st/go-home/call"
	_ "github.com/lgn21st/go-home/help"
)

func main() {
	args := os.Args

	if len(args) >= 2 {
		call.Start(args[1])
	}
}
