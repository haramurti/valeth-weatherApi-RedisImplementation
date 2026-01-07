package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type VisualCrossingResponse struct {
	Days []struct {
		Temp       float64 `json:"temp"`
		Humidity   float64 `json:"humidity"`
		WindSpeed  float64 `json:"windspeed"`
		UVIndex    float64 `json:"uvindex"`
		Conditions string  `json:"conditions"`
		Datetime   string  `json:"datetime"`
		Sunrise    string  `json:"sunrise"`
		Sunset     string  `json:"sunset"`
	} `json:"days"`
}

type Weather struct {
	ID        int     `json:"id"`
	City      string  `json:"city"`
	Weather   string  `json:"weather"`
	Time      string  `json:"time"`
	AvgTemp   float64 `json:"temp_celcius"`
	Humidity  int     `json:"humidity"`
	WindSpeed float64 `json:"wind_speed"`
	UVindex   int     `json:"uv_index"`
	Sunrise   string  `json:"sunrise"`
	Sunset    string  `json:"sunset"`
}

func GetWelcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to weather api"))
}

func GetCityWeather(w http.ResponseWriter, r *http.Request) {

	cityParam := r.PathValue("city")

	apiKey := os.Getenv("WEATHER_API_KEY")
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=%s&contentType=json", cityParam, apiKey)
	fmt.Println("shooting url :" + url)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "failed to shoot url : "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		http.Error(w, "error fetching url api", resp.StatusCode)
		return
	}

	var vcResponse VisualCrossingResponse
	if err := json.NewDecoder(resp.Body).Decode(&vcResponse); err != nil {
		http.Error(w, "failed to read weather", http.StatusInternalServerError)
		return
	}
	today := vcResponse.Days[0]

	finalData := Weather{
		ID:        1,
		City:      cityParam,
		Weather:   today.Conditions,
		Time:      today.Datetime,
		AvgTemp:   today.Temp,
		Humidity:  int(today.Humidity),
		WindSpeed: today.WindSpeed,
		UVindex:   int(today.UVIndex),
		Sunrise:   today.Sunrise,
		Sunset:    today.Sunset,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(finalData)

}
