package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const BaseUrl string = "https://api.nasa.gov/insight_weather/?api_key="

type Client struct {
	httpClient http.Client
}

func NewClient() *Client {
	return &Client{httpClient: http.Client{}}
}


func (c *Client) GetWeather() (NasaResponse, error) {
	apiKey := os.Getenv("API_Key")
	url := fmt.Sprintf("%s/%s&%s",BaseUrl,apiKey,"feedtype=json&ver=1.0")
	
	req, err := http.NewRequest("GET",url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return NasaResponse{}, err
	}
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return NasaResponse{}, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var nasaData NasaResponse
	err = decoder.Decode(&nasaData)
	if err != nil {
		fmt.Printf("Error decoding the response: %v\n", err)
		return NasaResponse{}, err
	}

	return nasaData, nil

}