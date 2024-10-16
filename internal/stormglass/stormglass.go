package stormglass

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"lr-surf-forecast/config"
)

type StormglassWeatherPointApiResponse struct {
	Hours []Hour `json:"hours"`
	Meta  Meta   `json:"meta"`
}

type Hour struct {
	AirTemperature   Source    `json:"airTemperature"`
	CurrentSpeed     Source    `json:"currentSpeed"`
	SeaLevel         Source    `json:"seaLevel"`
	SwellDirection   Source    `json:"swellDirection"`
	SwellHeight      Source    `json:"swellHeight"`
	SwellPeriod      Source    `json:"swellPeriod"`
	Time             time.Time `json:"time"`
	WaterTemperature Source    `json:"waterTemperature"`
	WaveDirection    Source    `json:"waveDirection"`
	WaveHeight       Source    `json:"waveHeight"`
	WavePeriod       Source    `json:"wavePeriod"`
	WindDirection    Source    `json:"windDirection"`
	WindSpeed        Source    `json:"windSpeed"`
}

type Source struct {
	Sg float64 `json:"sg"`
}

type Meta struct {
	Cost         int      `json:"cost,omitempty"`
	DailyQuota   int      `json:"dailyQuota,omitempty"`
	End          string   `json:"end,omitempty"`
	Lat          float64  `json:"lat,omitempty"`
	Lng          float64  `json:"lng,omitempty"`
	Params       []string `json:"params,omitempty"`
	RequestCount int      `json:"requestCount,omitempty"`
	Start        string   `json:"start,omitempty"`
}

// call stormglass api endpoint v2/weather/point
func GetWeatherData(lat, lng float64, start time.Time, duration int) (*StormglassWeatherPointApiResponse, error) {
	stormglassApiKey := os.Getenv("STORMGLASS_API_KEY")
	if stormglassApiKey == "" {
		return nil, fmt.Errorf("STORMGLASS_API_KEY environment variable is not set")
	}

	baseURL, err := url.Parse(config.StormglassApiEndpoint)
	if err != nil {
		return nil, err
	}
	baseURL.Path += "/weather/point"

	params := url.Values{}
	params.Add("lat", fmt.Sprintf("%f", lat))
	params.Add("lng", fmt.Sprintf("%f", lng))
	params.Add("params", "airTemperature,currentSpeed,seaLevel,swellDirection,swellHeight,swellPeriod,waterTemperature,waveDirection,waveHeight,wavePeriod,windDirection,windSpeed")
	params.Add("start", fmt.Sprintf("%d", start.Unix()))
	end := start.Add(time.Duration(duration) * 24 * time.Hour).Unix()
	params.Add("end", fmt.Sprintf("%d", end))
	params.Add("source", "sg")
	baseURL.RawQuery = params.Encode()

	log.Default().Printf("Calling stormglass API %s", baseURL.Host)

	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", stormglassApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get data: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherPointApiResponse StormglassWeatherPointApiResponse
	if err := json.Unmarshal(body, &weatherPointApiResponse); err != nil {
		return nil, err
	}

	return &weatherPointApiResponse, nil
}
