// Package open is a geo-golang based MapRequest Open geocode/reverse geocode client
package open

import (
	"github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Results []struct {
			Locations []struct {
				LatLng struct {
					Lat float64
					Lng float64
				}
				PostalCode string
				Street     string
				AdminArea6 string // neighbourhood
				AdminArea5 string // city
				AdminArea4 string // county
				AdminArea3 string // state
				AdminArea1 string // country (ISO 3166-1 alpha-2 code)
			}
		}
	}
)

// Geocoder constructs MapRequest Open geocoder
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
