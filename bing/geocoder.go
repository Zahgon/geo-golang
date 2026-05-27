// Package bing is a geo-golang based Microsoft Bing geocode/reverse geocode client
package bing

import (
	"github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		ResourceSets []struct {
			Resources []struct {
				Point struct {
					Coordinates []float64
				}
				Address struct {
					FormattedAddress string
					AddressLine      string
					AdminDistrict    string
					AdminDistrict2   string
					CountryRegion    string
					Locality         string
					PostalCode       string
				}
			}
		}
		ErrorDetails []string
	}
)

// Geocoder constructs Bing geocoder
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
