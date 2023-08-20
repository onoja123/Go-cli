package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fatih/color"
)

type Condition struct {
	Text string `json:"text"`
}

type Hour struct {
	TimeEpoch    int64     `json:"time_epoch"`
	TempC        float64   `json:"temp_c"`
	Condition    Condition `json:"condition"`
	ChanceOfRain float64   `json:"chance_of_rain"`
}

type Forecastday struct {
	Hour []Hour `json:"hour"`
}

type Location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Current struct {
	TempC     float64   `json:"temp_c"`
	Condition Condition `json:"condition"`
}

type Weather struct {
	Location Location    `json:"location"`
	Current  Current     `json:"current"`
	Forecast Forecastday `json:"forecast"`
}

func fetchWeatherData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Weather API not available")
	}

	return io.ReadAll(res.Body)
}

func isFutureDate(date time.Time) bool {
	return date.After(time.Now())
}

func printWeatherInfo(location Location, current Current) {
	fmt.Printf(
		"%s, %s: %.0fC, %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)
}

func main() {
	weatherData, err := fetchWeatherData("http://localhost")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var weather Weather
	err = json.Unmarshal(weatherData, &weather)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	printWeatherInfo(location, current)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		if !isFutureDate(date) {
			continue
		}

		message := fmt.Sprintf(
			"%s - %.0fC, %.0f%%, %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)

		if hour.ChanceOfRain < 40 {
			fmt.Print(message)
		} else {
			color.Red(message) // You may need to include the color package.
		}
	}
}
