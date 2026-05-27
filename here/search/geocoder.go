// Package search is a geo-golang based HERE geocode/reverse geocode client for the Geocoding and Search API
package search

import (
	"github.com/codingsince1985/geo-golang"
)

type (
	baseURL         struct{ forGeocode, forReverseGeocode string }
	geocodeResponse struct {
		Items []struct {
			Address struct {
				Label       string
				CountryCode string
				CountryName string
				StateCode   string
				State       string
				County      string
				District    string
				City        string
				Street      string
				PostalCode  string
				HouseNumber string
			}
			Position struct {
				Lat float64
				Lng float64
			}
		}
	}
)

// Geocoder constructs HERE geocoder
func Geocoder(apiKey string, baseURLs ...string) geo.Geocoder {
	_ = "STUB: not implemented"
	return *new(geo.Geocoder)
}

func getGeocodeURL(p string, baseURLs ...string) string { _ = "STUB: not implemented"; return "" }

func getReverseGeocodeURL(p string, baseURLs ...string) string {
	_ = "STUB: not implemented"
	return ""
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
