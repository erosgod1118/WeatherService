package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type WeatherResponse struct {
	Weather []struct {
		Description string `json:"description,omitempty"`
	}
	Wind struct {
		Speed float32 `json:"speed,omitempty"`
		Deg   float32 `json:"deg,omitempty"`
		Gust  float32 `json:"gust,omitempty"`
	} `json:"wind,omitempty"`
	Dt int `json:"dt,omitempty"`
}

func CurrentWeatherHandler(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")
	if lat == "" || lon == "" {
		http.Error(w, "Missing lat or lon parameter", http.StatusBadRequest)
		return
	}

	apiUrl := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s", lat, lon, os.Getenv("OPEN_WEATHER_API_KEY"))
	log.Println("Open Weather API URL: ", apiUrl)

	resp, err := http.Get(apiUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}
