package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	webhooksURL = kingpin.Flag("webhooks", "Incoming webhooks URL").Short('w').Required().URL()
	text        = kingpin.Arg("text", "A message text or JSON string.").Required().String()

	channel   = kingpin.Flag("channel", "Channel").Short('c').String()
	username  = kingpin.Flag("username", "Username").Short('u').String()
	iconEmoji = kingpin.Flag("icon-emoji", "Slack ICON Emoji").Short('e').String()
	iconURL   = kingpin.Flag("icon-url", "URL of an ICON image").Short('i').URL()
	isJSON    = kingpin.Flag("json", "The argument is JSON string").Short('j').Bool()
)

func main() {
	kingpin.Parse()

	var iconURLString string
	if (*iconURL) != nil {
		iconURLString = (*iconURL).String()
	}

	p := &params{
		channel:   *channel,
		username:  *username,
		iconEmoji: *iconEmoji,
		iconURL:   iconURLString,
	}

	var jsonMap string
	var err error

	if *isJSON {
		jsonMap, err = overwriteJSONParameter(*text, p)
	} else {
		jsonMap, err = createJSONParameter(*text, p)
	}

	if err != nil {
		outError(err)
	}

	err = post((*webhooksURL).String(), jsonMap)
	if err != nil {
		outError(err)
	}
}

func outError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
