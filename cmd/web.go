package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/kmonty-catamaran/deps-weather-webapp/pkg/app"
)

const (
	defaultPort = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	ipw, err := app.NewWeatherAdapter()
	if err != nil {
		exit("app.NewWeatherAdapter(): %v", err)
	}

	a := app.New(ipw)

	if err := http.ListenAndServe(":"+port, a.Handler()); err != nil {
		exit("http.ListenAndServe(): %v", err)
	}
}

func exit(f string, args ...any) {
	log.Printf("Error: "+f+"\n", args...)
	os.Exit(1)
}
