package main

import (
	"encoding/json"

	"gopkg.in/alecthomas/kingpin.v2"
)

type Arguments struct {
	WebHooksURL string
	Text        string
	IsStream    bool
	Option      *Option
}

func (args *Arguments) CreateJSON() (string, error) {
	mapData := args.Option.Map()
	mapData["text"] = args.Text

	buf, err := json.Marshal(mapData)

	return string(buf), err
}

type Option struct {
	channel   string
	username  string
	iconEmoji string
	iconURL   string
}

func (o *Option) Map() map[string]string {
	m := map[string]string{
		"channel":    o.channel,
		"username":   o.username,
		"icon_emoji": o.iconEmoji,
		"icon_url":   o.iconURL,
	}

	for k, v := range m {
		if v == "" {
			delete(m, k)
		}
	}

	return m
}

var (
	webhooksURL = kingpin.Flag("webhooks", "Incoming webhooks URL").Short('w').Required().URL()

	text      = kingpin.Arg("text", "A message text or JSON string.").String()
	channel   = kingpin.Flag("channel", "Channel").Short('c').String()
	username  = kingpin.Flag("username", "Username").Short('u').String()
	iconEmoji = kingpin.Flag("icon-emoji", "Slack ICON Emoji").Short('e').String()
	iconURL   = kingpin.Flag("icon-url", "URL of an ICON image").Short('i').URL()

	streamMode = kingpin.Flag("stream", "Stream mode (use stdin for JSON input)").Short('S').Bool()
)

func parseArguments() *Arguments {
	kingpin.Parse()

	option := &Option{
		channel:   *channel,
		username:  *username,
		iconEmoji: *iconEmoji,
	}

	if (*iconURL) != nil {
		option.iconURL = (*iconURL).String()
	}

	return &Arguments{
		WebHooksURL: (*webhooksURL).String(),
		Text:        *text,
		IsStream:    *streamMode,
		Option:      option,
	}
}
