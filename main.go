package main

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Name 		string
	Rate 		uint64
	Duration 	uint64
}

func main() {

	var config tomlConfig
	if _, err := toml.DecodeFile("example.toml", &config); err != nil {
		fmt.Println("No config file, using default settings.")
		config.Name = "Default"
	}
	fmt.Printf("Title: %s\n", config.Name)

	rate := uint64(100) // per second
	duration := 4 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:8080/",
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration) {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("%d results\n", metrics.Requests)
	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}
