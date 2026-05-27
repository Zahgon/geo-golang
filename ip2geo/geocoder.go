// Package ip2geo is a geo-golang based ip2geo.dev IP geolocation client
package ip2geo

import (
	"github.com/codingsince1985/geo-golang"
)

type geocoder struct {
	apiKey  string
	baseURL string
}

type apiResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		IP        string `json:"ip"`
		Type      string `json:"type"`
		Continent struct {
			Name    string `json:"name"`
			Code    string `json:"code"`
			Country struct {
				Name        string `json:"name"`
				Code        string `json:"code"`
				PhoneCode   string `json:"phone_code"`
				Capital     string `json:"capital"`
				Subdivision struct {
					Name string `json:"name"`
					Code string `json:"code"`
				} `json:"subdivision"`
				City struct {
					Name           string  `json:"name"`
					Latitude       float64 `json:"latitude"`
					Longitude      float64 `json:"longitude"`
					PostalCode     string  `json:"postal_code"`
					AccuracyRadius int     `json:"accuracy_radius"`
					Timezone       struct {
						Name string `json:"name"`
					} `json:"timezone"`
				} `json:"city"`
			} `json:"country"`
		} `json:"continent"`
		ASN struct {
			Number int    `json:"number"`
			Name   string `json:"name"`
		} `json:"asn"`
		RegisteredCountry struct {
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"registered_country"`
	} `json:"data"`
}

// Geocoder constructs an ip2geo geocoder
func Geocoder(apiKey string, baseURLs ...string) geo.Geocoder {
	_ = "STUB: not implemented"
	return *new(geo.Geocoder)
}

// Geocode returns location for the given IP address
func (g *geocoder) Geocode(address string) (*geo.Location, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReverseGeocode returns address for location.
// ip2geo is an IP geolocation service and does not support reverse geocoding.
func (g *geocoder) ReverseGeocode(lat, lng float64) (*geo.Address, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *geocoder) fetch(ip string) (*apiResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
