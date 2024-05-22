package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/echenim/TravelPathOptimizer/internal/models/entities"
	"github.com/echenim/TravelPathOptimizer/internal/models/views"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/planar"
)

func OptimizerRoute(tomtomAPIKey string, from, to string, driverHourThreshold, fuelCapacityThreshold int, truck entities.Truck) (views.RouteResponse, error) {
	routeURL := buildRouteURL(tomtomAPIKey, from, to, truck)
	return fetchRoute(routeURL)
}

func buildRouteURL(tomtomAPIKey, from, to string, truck entities.Truck) string {
	baseURL := fmt.Sprintf("https://api.tomtom.com/routing/1/calculateRoute/%s:%s/json?key=%s", from, to, tomtomAPIKey)
	params := prepareRouteParameters(truck)
	fullURL := fmt.Sprintf("%s%s", baseURL, params)
	return fullURL
}

func prepareRouteParameters(truck entities.Truck) string {
	var params []string
	params = append(params, formatVehicleSpecs(truck))
	params = append(params, handleRouteRestrictions(truck.RouteRestrictions)...)
	params = append(params, handleTrafficPatterns(truck.TrafficPatterns)...)
	params = append(params, handleFuelNeeds(truck)...)
	params = append(params, handleDriverHours(truck)...)
	return strings.Join(params, "&")
}

// Separate functions for different concerns
func formatVehicleSpecs(truck entities.Truck) string {
	return fmt.Sprintf("&maxHeight=%d&maxWeight=%d", truck.Dimensions.Height, truck.Weight)
}

func handleRouteRestrictions(restrictions []string) []string {
	var params []string
	for _, restriction := range restrictions {
		params = append(params, fmt.Sprintf("&avoid=%s", restriction))
	}
	return params
}

func handleTrafficPatterns(patterns map[string]string) []string {
	var params []string
	for location, trafficLevel := range patterns {
		// Determine action based on the traffic level
		switch trafficLevel {
		case "high":
			avoidParam := fmt.Sprintf("avoidArea=%s", formatGeoFencing(location))
			params = append(params, avoidParam)
		case "medium":
			// For medium traffic, you might choose to simply monitor the area, or use real-time traffic data to decide dynamically
			monitorParam := fmt.Sprintf("monitorArea=%s", formatGeoFencing(location))
			params = append(params, monitorParam)
		// Low traffic doesn't require any action
		case "low":
			continue
		default:
			// Handle unknown traffic levels or implement additional logic as needed
			continue
		}
	}
	return params
}

// Helper function to convert a location identifier to a geo-fencing parameter suitable for API requests
// This is a stub and needs actual logic based on how locations are defined and how geo-fencing works with your chosen API
// TODO: implement conversion logic
func formatGeoFencing(location string) string {
	return location
}

// handleFuelNeeds determines the need for fuel based on the truck's fuel capacity and
// adds appropriate waypoints to refuel if necessary.
func handleFuelNeeds(truck entities.Truck, stations []entities.FuelStation, route *geojson.Feature) []string {
	var params []string
	fuelNeeded := calculateFuelNeeded(truck)
	if truck.FuelCapacity < fuelNeeded {
		stationsAlongRoute := findStationsAlongRoute(truck.CurrentLocation, truck.Destination, stations, route)
		optimalStations := selectOptimalFuelStations(stationsAlongRoute, truck)
		for _, station := range optimalStations {
			params = append(params, fmt.Sprintf("&addWaypoints=%s", station.Location))
		}
	}
	return params
}

// calculateFuelNeeded estimates the fuel required for the journey based on distance and truck's efficiency.
// TODO: implement
func calculateFuelNeeded(truck entities.Truck) int {
	distance := calculateDistance(truck.CurrentLocation, truck.Destination)
	return distance
}

// selectOptimalFuelStations selects the most strategically placed fuel stations to minimize route deviations.
func selectOptimalFuelStations(stations []entities.FuelStation, truck entities.Truck) []entities.FuelStation {
	return stations
}

// calculateDistance simulates distance calculation between two geographic locations.
// TODO: full implementation of distance calculation is needed
func calculateDistance(from, to entities.Location) int {
	return 100
}

func handleDriverHours(truck entities.Truck) []string {
	var params []string
	// Add logic for driver hours
	return params
}

func fetchRoute(url string) (views.RouteResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return views.RouteResponse{}, fmt.Errorf("failed to send request to TomTom API: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return views.RouteResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}
	var routeResponse views.RouteResponse
	if err := json.Unmarshal(body, &routeResponse); err != nil {
		return views.RouteResponse{}, fmt.Errorf("failed to unmarshal JSON response: %w", err)
	}
	return routeResponse, nil
}

// Improved function to find stations along the route using geographic calculations
func findStationsAlongRoute(from, to entities.Location, stations []entities.FuelStation, route *geojson.Feature) []entities.FuelStation {
	var stationsAlongRoute []entities.FuelStation

	routeLine, _ := route.Geometry.(orb.LineString)
	for _, station := range stations {
		if isStationAlongRoute(routeLine, station.Location, 1000) { // 1000 meters tolerance
			stationsAlongRoute = append(stationsAlongRoute, station)
		}
	}
	return stationsAlongRoute
}

// Uses geospatial logic to check if a station is close enough to the route
func isStationAlongRoute(route orb.LineString, stationLocation orb.Point, tolerance float64) bool {
	distance := planar.LineStringDistanceFromPoint(route, stationLocation)
	return distance <= tolerance
}

// TODO: implement
func calculateDetour(location string) string {
	return "someDetourArea"
}
