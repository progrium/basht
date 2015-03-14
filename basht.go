package main

import (
	"fmt"
	"os"

	"github.com/progrium/go-basher"
)

var Version string

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Println(Version)
		os.Exit(0)
	}
	os.Setenv("VERSION", Version)
	basher.Application(nil, []string{
		"include/basht.bash",
	}, Asset, true)
}
