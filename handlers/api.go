package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetWeatherRapidAPI(city string) (string, error) {
	url := fmt.Sprintf("https://weather-by-api-ninjas.p.rapidapi.com/v1/weather?city=%s", city)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("X-RapidAPI-Key", os.Getenv("RAPIDAPI"))
	req.Header.Add("X-RapidAPI-Host", "weather-by-api-ninjas.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
