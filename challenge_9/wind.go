package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	url := "https://jsonplaceholder.typicode.com/posts"

	for {
		// generate random values for water and wind
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		// determine status based on water and wind values
		waterStatus := ""
		windStatus := ""

		if water < 5 {
			waterStatus = "aman"
		} else if water >= 6 && water <= 8 {
			waterStatus = "siaga"
		} else {
			waterStatus = "bahaya"
		}

		if wind < 6 {
			windStatus = "aman"
		} else if wind >= 7 && wind <= 15 {
			windStatus = "siaga"
		} else {
			windStatus = "bahaya"
		}

		// create JSON payload
		data := map[string]interface{}{
			"water":       water,
			"waterStatus": waterStatus,
			"wind":        wind,
			"windStatus":  windStatus,
		}

		payload, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			continue
		}

		// make POST request
		resp, err := http.Post(url, "application/json", strings.NewReader(string(payload)))
		if err != nil {
			fmt.Println("Error making POST request:", err)
			continue
		}

		// print response
		fmt.Printf("Sent data: %s\n", string(payload))
		fmt.Printf("Status water: %s\n", waterStatus)
		fmt.Printf("Status wind: %s\n", windStatus)
		fmt.Printf("Status Code: %d\n\n", resp.StatusCode)

		time.Sleep(15 * time.Second)
	}
}
