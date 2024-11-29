package main

import (
	"encoding/json"
	"fmt"
	"io"
	_ "io/ioutil"
	"net/http"
)

/*
	Goal is to hit some website
	Fetch the value asynchronously
*/

type Weather struct {
	Main MainDetails `json:"main"`
	Name string      `json:"name"`
}

type MainDetails struct {
	Temperature        float64 `json:"temp"`
	FeelsLike          float64 `json:"feels_like"`
	MinimumTemperature float64 `json:"temp_min"`
	MaximumTemperature float64 `json:"temp_max"`
	Pressure           int64   `json:"pressure"`
	Humidity           int64   `json:"humidity"`
}

func FetchWeather() {
	resp, err := http.Get(`https://api.openweathermap.org/data/2.5/weather?q=New%20York&appid=de9a37c14517b0a371c5d284b35f0bef&units=metric`)
	if err != nil {
		fmt.Println("Error fetching weather data:", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing body:", err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
	}
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Println("Error parsing body:", err)
	}
	fmt.Println(weather.Main.Humidity)
}

func main() {
	FetchWeather()
}
