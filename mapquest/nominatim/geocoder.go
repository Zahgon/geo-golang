// Package nominatim is a geo-golang based MapRequest Nominatim geocode/reverse geocode client
package nominatim

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/osm"
)

type (
	baseURL string

	geocodeResponse struct {
		DisplayName     string `json:"display_name"`
		Lat, Lon, Error string
		Addr            osm.Address `json:"address"`
	}
)

var key string

// Geocoder constructs MapRequest Nominatim geocoder
func Geocoder(k string, baseURLs ...string) geo.Geocoder {
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
