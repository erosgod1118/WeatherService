package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
)

var serverAddr = flag.String("listen", "localhost:8080", "web server http listen address")

func main() {
	flag.Parse()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	svr := WeatherServer(*serverAddr)
	log.Printf("Launching weather server at %s", *serverAddr)
	log.Fatal(svr.ListenAndServe())
}
