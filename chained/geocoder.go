package chained

import (
	"github.com/codingsince1985/geo-golang"
)

type chainedGeocoder struct{ Geocoders []geo.Geocoder }

// Geocoder creates a chain of Geocoders to lookup address and fallback on
func Geocoder(geocoders ...geo.Geocoder) geo.Geocoder {
	_ = "STUB: not implemented"
	return *new(geo.Geocoder)
}

// Geocode returns location for address
func (c chainedGeocoder) Geocode(address string) (*geo.Location, error) {
	_ = "STUB: not implemented"
	// Geocode address by each geocoder until we get a real location response
	return nil, nil
}

// skip error and try the next geocoder

// No geocoders found a result

// ReverseGeocode returns address for location
func (c chainedGeocoder) ReverseGeocode(lat, lng float64) (*geo.Address, error) {
	_ = "STUB: not implemented"
	// Geocode address by each geocoder until we get a real location response
	return nil, nil
}

// skip error and try the next geocoder

// No geocoders found a result
