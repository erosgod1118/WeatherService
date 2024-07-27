package main

import (
	"net/http"
	"time"

	"github.com/erosgod1118/WeatherService/handlers"
	"github.com/gorilla/mux"
)

func WeatherServer(serverAddr string) *http.Server {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/current_weather", handlers.CurrentWeatherHandler)

	svr := &http.Server{
		Addr:           serverAddr,
		Handler:        rtr,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return svr
}
