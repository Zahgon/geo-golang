// Package opencage is a geo-golang based OpenCage geocode/reverse geocode client
package opencage

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/osm"
)

type (
	baseURL string

	geocodeResponse struct {
		Results []struct {
			Formatted  string
			Geometry   geo.Location
			Components osm.Address
		}
		Status struct {
			Code    int
			Message string
		}
	}
)

// Geocoder constructs OpenCage geocoder
func Geocoder(key string, baseURLs ...string) geo.Geocoder {
	_ = "STUB: not implemented"
	return *new(geo.Geocoder)
}

func getURL(key string, baseURLs ...string) string { _ = "STUB: not implemented"; return "" }

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
