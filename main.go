package main

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {

	rate := uint64(100) // per second
	duration := 4 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://localhost:8080/",
	})
	attacker := vegeta.NewAttacker()

	var results vegeta.Results
	for res := range attacker.Attack(targeter, rate, duration) {
		results = append(results, *res)
	}

	fmt.Printf("%d results\n", results.Len())
}
