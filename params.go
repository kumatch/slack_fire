package main

import (
	"encoding/json"
	"strings"

	"github.com/golang/golang/go/src/fmt"
)

type params struct {
	channel   string
	username  string
	iconEmoji string
	iconURL   string
}

func (p *params) Map() *map[string]string {
	m := map[string]string{
		"channel":   p.channel,
		"username":  p.username,
		"iconEmoji": p.iconEmoji,
		"iconURL":   p.iconEmoji,
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

	return convertJSONString(mapData)
}

func overwriteJSONParameter(jsonStr string, p *params) (string, error) {
	var data interface{}
	dec := json.NewDecoder(strings.NewReader(jsonStr))
	err := dec.Decode(&data)
	if err != nil {
		return "", err
	}

	var mapData map[string]interface{}

	switch data.(type) {
	case map[string]interface{}:
		mapData = data.(map[string]interface{})
	default:
		return "", fmt.Errorf("Invalid JSON format.")
	}

	for k, v := range *(p.Map()) {
		if v != "" {
			mapData[k] = v
		}
	}

	return convertJSONString(mapData)
}

func convertJSONString(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
