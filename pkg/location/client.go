package location

import (
	"context"

	"googlemaps.github.io/maps"
)

// Maps client struct
type mapsClient struct {
	client *maps.Client
}

// Create maps client
func createClient(apiKey string) *mapsClient {
	c, _ := maps.NewClient(maps.WithAPIKey(apiKey))
	return &mapsClient{c}
}

// Get location using geocoding API
func (m *mapsClient) queryLoc(address string) ([]maps.GeocodingResult, error) {
	r := &maps.GeocodingRequest{Address: address}
	return (*m).client.Geocode(context.Background(), r)
}
