package main

import (
	"encoding/json"
)

type params struct {
	text      string
	channel   string
	username  string
	iconEmoji string
	iconURL   string
}

func (p *params) CreateMap() *map[string]string {
	m := make(map[string]string)

	m["text"] = p.text

	if p.channel != "" {
		m["channel"] = p.channel
	}
	if p.username != "" {
		m["username"] = p.username
	}
	if p.iconEmoji != "" {
		m["icon_emoji"] = p.iconEmoji
	}
	if p.iconURL != "" {
		m["icon_url"] = p.iconURL
	}

	return &m
}

func createJSONParameter(text string, p *params) (string, error) {
	p.text = text
	b, err := json.Marshal(p.CreateMap())
	if err != nil {
		return "", err
	}

	return string(b), nil
}

/*
func overwriteJSONParameter(json string, p *params) (string, error) {
	p.text = text
	m := p.CreateMap()

	b, err := json.Marshal(m)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
*/
