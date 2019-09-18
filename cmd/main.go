package main

import (
	"fmt"

	"github.com/neomrc/go-google-maps/pkg/location"
)

func main() {
	l := &location.Location{"680666"}

	fmt.Println(l.GetLocation())
	// c, err := maps.NewClient(maps.WithAPIKey("AIzaSyCiN6qvrrxpqwL3xJNOSV4FxDPb08WtXy4"))
	// if err != nil {
	// 	log.Fatalf("fatal error: %s", err)
	// }
	// r := &maps.DirectionsRequest{
	// 	Origin:      "Sydney",
	// 	Destination: "Perth",
	// }
	// route, _, err := c.Directions(context.Background(), r)
	// if err != nil {
	// 	log.Fatalf("fatal error: %s", err)
	// }

	// pretty.Println(route)
}
