package main

import (
	"fmt"
	"os"
)

func main() {
	args := parseArguments()

	if args.IsStream {
		runStreamMode(args)
	} else {
		runArgsMode(args)
	}
}

func runStreamMode(args *Arguments) {
	err := postStream(args.WebHooksURL, os.Stdin)
	if err != nil {
		outError(err)
	}
}

func runArgsMode(args *Arguments) {
	if args.Text == "" {
		outError(fmt.Errorf("A message is blank, stopped."))
		return
	}

	jsonString, err := args.CreateJSON()
	if err != nil {
		outError(err)
	}

	err = postJSON(args.WebHooksURL, jsonString)
	if err != nil {
		outError(err)
	}
}

func outError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
