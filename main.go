package main

import (
	"os"

	"github.com/parasraut21/GoChain/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}
