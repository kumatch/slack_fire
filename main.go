package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	webhooksURL = kingpin.Flag("webhooks", "Incoming webhooks URL").Short('w').Required().URL()

	text      = kingpin.Arg("text", "A message text or JSON string.").String()
	channel   = kingpin.Flag("channel", "Channel").Short('c').String()
	username  = kingpin.Flag("username", "Username").Short('u').String()
	iconEmoji = kingpin.Flag("icon-emoji", "Slack ICON Emoji").Short('e').String()
	iconURL   = kingpin.Flag("icon-url", "URL of an ICON image").Short('i').URL()

	streamMode = kingpin.Flag("stream", "Stream mode (use stdin for JSON input)").Short('S').Bool()
)

func main() {
	kingpin.Parse()

	var iconURLString string
	if (*iconURL) != nil {
		iconURLString = (*iconURL).String()
	}

	url := (*webhooksURL).String()

	if *streamMode {
		runStreamMode(url)
	} else {
		if *text == "" {
			outError(fmt.Errorf("A message is blank, stopped."))
			return
		}

		p := &params{
			channel:   *channel,
			username:  *username,
			iconEmoji: *iconEmoji,
			iconURL:   iconURLString,
		}

		runArgsMode(url, *text, p)
	}
}

func runStreamMode(url string) {
	err := postStream(url, os.Stdin)
	if err != nil {
		outError(err)
	}
}

func runArgsMode(url, text string, p *params) {
	jsonString, err := createJSONParameter(text, p)
	if err != nil {
		outError(err)
	}

	err = post(url, jsonString)
	if err != nil {
		outError(err)
	}
}

func outError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
