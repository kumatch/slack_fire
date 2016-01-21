package main

import "encoding/json"

type params struct {
	channel   string
	username  string
	iconEmoji string
	iconURL   string
}

func (p *params) Map() *map[string]string {
	m := map[string]string{
		"channel":    p.channel,
		"username":   p.username,
		"icon_emoji": p.iconEmoji,
		"icon_url":   p.iconURL,
	}

	for k, v := range m {
		if v == "" {
			delete(m, k)
		}
	}

	return &m
}

func createJSONParameter(text string, p *params) (string, error) {
	mapData := p.Map()
	(*mapData)["text"] = text

	return json.Marshal(mapData)
}
