package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func GetWeatherRapidAPI(city string) (string, error) {
	url := fmt.Sprintf("https://weather-by-api-ninjas.p.rapidapi.com/v1/weather?city=%s", city)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("X-RapidAPI-Key", "1930c8ef0dmsh589af469b5e014ap177d48jsna6e21baafd45")
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
