package location

import (
	"fmt"

	"googlemaps.github.io/maps"
)

// Location Struct
type Location struct {
	Address string
}

// Get Location func to get data from cache or maps api
func (l *Location) GetLocation() maps.GeocodingResult {
	var location maps.GeocodingResult
	// cache, err := getCachedLocation(l.Address)

	// if err != nil {
	// 	fmt.Errorf("Error fetching location")
	// }

	// if cache != nil {
	// 	return cache
	// }

	location = getFreshLocation(l.Address)

	fmt.Printf("%+v\n", location)

	return location
}

// Get location first check stored results on cache then on google maps API
func getFreshLocation(address string) maps.GeocodingResult {
	client := createClient("AIzaSyCiN6qvrrxpqwL3xJNOSV4FxDPb08WtXy4")
	results, err := client.queryLoc(address)

	if err != nil {
		fmt.Errorf("Error fetching location from Google Maps API")
		return maps.GeocodingResult{}
	}

	if len(results) == 0 {
		fmt.Errorf("No results found")
		return maps.GeocodingResult{}
	}

	return results[0]
}

// func getCachedLocation(address string) (*maps.GeocodingResult, error) {

// }
