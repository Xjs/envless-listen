package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func help() {
	fmt.Println(`# envless-listen

This is to demonstrate that on Windows, the SYSTEMROOT variable needs to be set in order for a ` + "`net.Listener`" + ` to work. This is especially not the case when executed from an os.Cmd with a custom environment.

## Setup

` + "`cd listener && go build`" + `
` + "`go build`" + `

## Working case

` + "`./envless-listen -working`" + `

## Failing case

` + "`./envless-listen -failing`" + `

`)
}

func main() {
	var working, failing bool
	flag.BoolVar(&working, "working", working, "Run the working example")
	flag.BoolVar(&failing, "failing", failing, "Run the failing example")
	flag.Parse()

	if working == failing {
		help()
		return
	}

	env := []string{}
	switch {
	case working:
		env = append(env, "SYSTEMROOT="+os.Getenv("SYSTEMROOT"))
	}

	cmd := exec.Command("listener/listener")
	cmd.Env = env

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
