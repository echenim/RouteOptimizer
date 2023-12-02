package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/plak3com/route-optimizer-engine/internal/models/entities"
	"github.com/plak3com/route-optimizer-engine/internal/models/views"
)

func OptimizerRoute(tomtom_api_key, from, to string, driverHourThreshold, fuelCapacityThreshold int, truck entities.Truck) (views.RouteResponse, error) {
	return findBestRoutes(tomtom_api_key, from, to, driverHourThreshold, fuelCapacityThreshold, truck)
}

func findBestRoutes(tomtom_api_key, from, to string, driverHourThreshold, fuelCapacityThreshold int, truck entities.Truck) (views.RouteResponse, error) {
	// Construct the base URL with truck dimensions and weight
	url := fmt.Sprintf("https://api.tomtom.com/routing/1/calculateRoute/%s:%s/json?maxHeight=%d&maxWeight=%d&key=%s",
		from, to, truck.Dimensions.Height, truck.Weight, tomtom_api_key)

	// Add additional parameters based on truck properties
	if truck.Hazmat {
		url += "&avoid=tunnels"
	}

	// Consider route restrictions based on truck properties
	for _, restriction := range truck.RouteRestrictions {
		switch restriction {
		case "tollRoads":
			url += "&avoid=tollRoads"
		case "motorways":
			url += "&avoid=motorways"
		case "ferries":
			url += "&avoid=ferries"
			// Add more cases as needed
		}
	}

	// Logic to avoid high-traffic areas
	for location, pattern := range truck.TrafficPatterns {
		if pattern == "high" {
			// Logic to adjust the route to avoid this location
			// This could involve adding an intermediate waypoint that routes around the high traffic area
			// or using TomTom API features to avoid certain areas
			avoidArea := calculateDetour(location)
			url += fmt.Sprintf("&avoid=%s", avoidArea)
		}
	}

	// Consider Fuel Stations if fuel capacity is a constraint
	// Logic for Fuel Capacity
	if truck.FuelCapacity < fuelCapacityThreshold {
		// Find fuel stations along the route
		stationsAlongRoute := findStationsAlongRoute(from, to, truck.FuelStations)

		// Modify the URL to include waypoints for fuel stations
		// Adding multiple fuel stations as waypoints could make the routing less efficient.
		// We will need a strategy to select the most optimal fuel stations, considering factors
		// like distance from the route, fuel prices, and station amenities.
		for _, station := range stationsAlongRoute {
			url += fmt.Sprintf("&waypoints=%s", station.Location)
		}
	}

	// Handle refrigeration needs
	if truck.Refrigeration {
		// Logic to prioritize routes with shorter travel times or specific conditions
		// This could be a placeholder as TomTom API might not support this directly
	}

	// Add logic for axle weight (assuming TomTom API supports this)
	if len(truck.AxleWeight) > 0 {
		axleWeights := make([]string, len(truck.AxleWeight))
		for i, weight := range truck.AxleWeight {
			axleWeights[i] = strconv.Itoa(weight)
		}
		url += "&axleWeight=" + strings.Join(axleWeights, ",")
	}

	// Handle driver hours - determine if a rest stop is needed
	if truck.DriverHours > driverHourThreshold {
		// Logic to include rest areas in the route
		url += "&restAreas=" + strings.Join(truck.RestAreas, ",")
	}

	// avoid toll roads if the truck doesn't have a toll system
	if truck.TollSystem == "" {
		url += "&avoid=tollRoads"
	}

	// Send request to TomTom API
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error occurred while sending request to TomTom API: %s", err)
	}

	defer resp.Body.Close()

	// Handle the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error occurred while reading the response body: %s", err)
	}

	// Assuming RouteResponse is defined to parse the relevant data
	var routeResponse views.RouteResponse
	if err := json.Unmarshal(body, &routeResponse); err != nil {
		log.Fatalf("Error occurred while unmarshalling the JSON response: %s", err)
	}

	return routeResponse, nil
}

func findStationsAlongRoute(from, to string, stations []entities.FuelStation) []entities.FuelStation {
	var stationsAlongRoute []entities.FuelStation
	// Logic to determine if a station falls along the route
	// This is a placeholder. We would need
	// a more sophisticated approach to determine if a station is along the route.
	for _, station := range stations {
		if isStationAlongRoute(from, to, station) {
			stationsAlongRoute = append(stationsAlongRoute, station)
		}
	}
	return stationsAlongRoute
}

func isStationAlongRoute(from, to string, station entities.FuelStation) bool {
	// Placeholder logic to determine if a station is along the route
	// We might use geographic calculations or additional API queries
	return true
}

func calculateDetour(location string) string {
	// Placeholder logic to calculate a detour around a high-traffic area
	// to determine an effective detour
	return "someDetourArea"
}
