package cached

import (
	"github.com/codingsince1985/geo-golang"
	"github.com/patrickmn/go-cache"
)

type cachedGeocoder struct {
	Geocoder geo.Geocoder
	Cache    *cache.Cache
}

// Geocoder creates a chain of Geocoders to lookup address and fallback on
func Geocoder(geocoder geo.Geocoder, cache *cache.Cache) geo.Geocoder {
	_ = "STUB: not implemented"
	return *new(geo.Geocoder)
}

// Geocode returns location for address
func (c cachedGeocoder) Geocode(address string) (*geo.Location, error) {
	_ = "STUB: not implemented"
	// Check if we've cached this response
	return nil, nil
}

// ReverseGeocode returns address for location
func (c cachedGeocoder) ReverseGeocode(lat, lng float64) (*geo.Address, error) {
	_ = "STUB: not implemented"
	// Check if we've cached this response
	return nil, nil
}
