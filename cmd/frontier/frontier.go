package main

import (
	"fmt"
	"os"

	"github.com/gofrontier-com/frontier/pkg/cmd/frontier"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	command := frontier.NewRootCmd(version, commit, date)
	if err := command.Execute(); err != nil {
		fmt.Println("Error from main")
		fmt.Println(err)
		os.Exit(1)
	}
}
