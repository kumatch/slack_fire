package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	contentType = "application/json"
)

func post(url, json string) error {
	resp, err := http.Post(url, contentType, strings.NewReader(json))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", body)
	}

	return nil
}
