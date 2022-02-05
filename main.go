package main

import (
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

// Config sets application variables
type Config struct {
	Port        string `envconfig:"PORT" default:"8080"`
	StaticsPath string `envconfig:"STATICS_PATH" default:"./static"`
}

func main() {
	config := new(Config)
	err := envconfig.Process("", config)
	if err != nil {
		log.Fatalf("Can't process config: %v", err)
	}

	fs := http.FileServer(http.Dir(config.StaticsPath))
	http.Handle("/", fs)

	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		log.Fatalf("Error while serving: %v", err)
	}
}
