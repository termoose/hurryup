package speed

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
	"time"
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

	var totalDownloadTime time.Duration
	var totalSize int64
	for _, t := range targets {
		var firstByteTime time.Time
		req, _ := http.NewRequest(http.MethodGet, t, nil)
		trace := &httptrace.ClientTrace{
			GotFirstResponseByte: func() {
				firstByteTime = time.Now()
			},
		}

		req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
		resp, err := http.DefaultTransport.RoundTrip(req)
		if err != nil {
			panic(err)
		}

		// Discard the body but count the number of bytes in it
		bytes, _ := io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()

		totalDownloadTime += time.Since(firstByteTime)
		totalSize += bytes / 1024 / 1024 * 8
	}

	fmt.Printf("transfer time: %s size: %d Mbit speed: %f Mbit/s\n", totalDownloadTime, totalSize,
		float64(totalSize)/totalDownloadTime.Seconds())
}

func getTargetURLs() ([]string, error) {
	url := fmt.Sprintf(TARGET_URL_ENDPOINT, FAST_TOKEN, URL_COUNT)
	resp, err := http.Get(url)

	if err != nil {
		return []string{}, fmt.Errorf("could not get target urls: %w", err)
	}

	if resp.StatusCode == http.StatusForbidden {
		return []string{}, errors.New("forbidden, invalid token")
	}

	var response targetResponse
	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		return []string{}, fmt.Errorf("could not parse json response: %w", err)
	}

	fmt.Printf("client: %+v\n", response.Client)

	var result []string
	for _, u := range response.Targets {
		result = append(result, u.Url)
	}

	return result, nil
}
