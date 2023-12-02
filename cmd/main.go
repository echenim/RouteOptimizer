package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const apiKey = "Your_API_KEY"



func getRoute(from, to string) {
	url := fmt.Sprintf("https://api.tomtom.com/routing/1/calculateRoute/%s:%s/json?key=%s", from, to, tomtomAPIKey)

	// Send a GET request to the TomTom API
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error occurred while sending request to TomTom API: %s", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error occurred while reading the response body: %s", err)
	}

	// Unmarshal the JSON response into a RouteResponse struct
	var routeResponse RouteResponse
	if err := json.Unmarshal(body, &routeResponse); err != nil {
		log.Fatalf("Error occurred while unmarshalling the JSON response: %s", err)
	}

	// Process and print route information
	if len(routeResponse.Routes) > 0 {
		route := routeResponse.Routes[0]
		fmt.Printf("Length: %d meters\n", route.Summary.LengthInMeters)
		fmt.Printf("Travel Time: %d seconds\n", route.Summary.TravelTimeInSeconds)
		fmt.Printf("Traffic Delay: %d seconds\n", route.Summary.TrafficDelayInSeconds)
	} else {
		fmt.Println("No route found")
	}
}

func main() {
	getRoute("52.50931,13.42936", "52.50274,13.43872")
}
