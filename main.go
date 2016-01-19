package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	webhooksUrl = kingpin.Flag("webhooks", "Incoming webhooks URL").Short('w').Required().URL()
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

	json, err := createJSONParameter(*text, p)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = post((*webhooksUrl).String(), json)
	if err != nil {
		fmt.Println(err.Error())
	}
}
