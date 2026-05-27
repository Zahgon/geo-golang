// Package openstreetmap is a geo-golang based OpenStreetMap geocode/reverse geocode client
package openstreetmap

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

// Geocoder constructs OpenStreetMap geocoder
func Geocoder() geo.Geocoder { _ = "STUB: not implemented"; return *new(geo.Geocoder) }

// GeocoderWithURL constructs OpenStreetMap geocoder using a custom installation of Nominatim
func GeocoderWithURL(nominatimURL string) geo.Geocoder {
	_ = "STUB: not implemented"
	return *new(geo.Geocoder)
}

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
