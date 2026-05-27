// Package here is a geo-golang based HERE geocode/reverse geocode client for the legacy geocoder API
package here

import (
	"github.com/codingsince1985/geo-golang"
)

type (
	baseURL         struct{ forGeocode, forReverseGeocode string }
	geocodeResponse struct {
		Response struct {
			View []struct {
				Result []struct {
					Location struct {
						DisplayPosition struct {
							Latitude, Longitude float64
						}
						Address struct {
							Label          string
							Country        string
							State          string
							County         string
							City           string
							District       string
							Street         string
							HouseNumber    string
							PostalCode     string
							AdditionalData []struct {
								Key   string
								Value string
							}
						}
					}
				}
			}
		}
	}
)

// Key*Name represents constants for geocoding more address detail
const (
	KeyCountryName = "CountryName"
	KeyStateName   = "StateName"
	KeyCountyName  = "CountyName"
)

var r = 100

// Geocoder constructs HERE geocoder
func Geocoder(id, code string, radius int, baseURLs ...string) geo.Geocoder {
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
