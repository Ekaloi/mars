package api

import "net/http"

const BaseUrl string = "https://api.nasa.gov/insight_weather/"

type Client struct {
	httpClient http.Client
}

func NewClient() *Client {
	return &Client{httpClient: http.Client{}}
}
