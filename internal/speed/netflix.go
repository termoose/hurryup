package speed

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var TARGET_URL_ENDPOINT = "https://api.fast.com/netflix/speedtest/v2?https=true&token=%s&urlCount=%d"
var FAST_TOKEN = "YXNkZmFzZGxmbnNkYWZoYXNkZmhrYWxm"
var URL_COUNT = 5

type Netflix struct {
}

type location struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type target struct {
	Name     string   `json:"name"`
	Url      string   `json:"url"`
	Location location `json:"location"`
}

type client struct {
	Ip       string   `json:"ip"`
	Asn      string   `json:"asn"`
	Location location `json:"location"`
}

type targetResponse struct {
	Client  client   `json:"client"`
	Targets []target `json:"targets"`
}

func NewNetflix() (*Netflix, error) {
	getDownloadSpeed()

	return &Netflix{}, nil
}

func getDownloadSpeed() {
	targets, err := getTargetURLs()

	if err != nil {
		panic(err)
	}

	for i, t := range targets {
		resp, _ := http.Get(t)

		fmt.Printf("%d: %s size: %d MB\n", i, t, resp.ContentLength/1024/1024)
	}
}

func getTargetURLs() ([]string, error) {
	url := fmt.Sprintf(TARGET_URL_ENDPOINT, FAST_TOKEN, URL_COUNT)
	resp, err := http.Get(url)

	if err != nil {
		return []string{}, fmt.Errorf("could not get target urls: %w", err)
	}

	if resp.StatusCode == http.StatusForbidden {
		return []string{}, fmt.Errorf("forbidden, invalid token")
	}

	var response targetResponse
	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		return []string{}, fmt.Errorf("could not parse json response: %w", err)
	}

	var result []string
	for _, u := range response.Targets {
		result = append(result, u.Url)
	}

	return result, nil
}
