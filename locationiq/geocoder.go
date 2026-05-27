// Package locationiq is a geo-golang based LocationIQ geocode/reverse geocode client
package locationiq

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/osm"
)

type baseURL string

type geocodeResponse struct {
	DisplayName     string `json:"display_name"`
	Lat, Lon, Error string
	Addr            osm.Address `json:"address"`
}

const (
	defaultURL  = "http://locationiq.org/v1/"
	minZoom     = 0  // Min zoom level for locationiq - country level
	maxZoom     = 18 // Max zoom level for locationiq - house level
	defaultZoom = 18
)

var (
	key  string
	zoom int
)

// Geocoder constructs LocationIQ geocoder
func Geocoder(k string, z int, baseURLs ...string) geo.Geocoder {
	_ = "STUB: not implemented"
	return *new(geo.Geocoder)
}

func (b baseURL) GeocodeURL(address string) string { _ = "STUB: not implemented"; return "" }

func (b baseURL) ReverseGeocodeURL(l geo.Location) string { _ = "STUB: not implemented"; return "" }

func (r *geocodeResponse) Location() (*geo.Location, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// no result

func (r *geocodeResponse) Address() (*geo.Address, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
