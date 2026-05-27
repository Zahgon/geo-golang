// Package pickpoint is a geo-golang based PickPoint geocode/reverse geocode client
package pickpoint

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/osm"
)

type (
	baseURL         string
	geocodeResponse struct {
		DisplayName string `json:"display_name"`
		Lat         string
		Lon         string
		Error       string
		Addr        osm.Address `json:"address"`
	}
)

var key string

// Geocoder constructs PickPoint geocoder
func Geocoder(apiKey string, baseURLs ...string) geo.Geocoder {
	_ = "STUB: not implemented"
	return *new(geo.Geocoder)
}

func getURL(baseURLs ...string) string { _ = "STUB: not implemented"; return "" }

func (b baseURL) GeocodeURL(address string) string { _ = "STUB: not implemented"; return "" }

func (b baseURL) ReverseGeocodeURL(l geo.Location) string { _ = "STUB: not implemented"; return "" }

func (r *geocodeResponse) Location() (*geo.Location, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
