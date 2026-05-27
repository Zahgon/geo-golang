// Package tomtom is a geo-golang based TomTom geocode/reverse geocode client
package tomtom

import (
	geo "github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Summary struct {
			Query string
		}

		Results []struct {
			Position struct {
				Lat float64
				Lon float64
			}
		}

		// Reverse Geocoding response
		Addresses []struct {
			Address struct {
				BuildingNumber              string
				StreetNumber                string
				Street                      string
				StreetName                  string
				StreetNameAndNumber         string
				CountryCode                 string
				CountrySubdivision          string // state code
				CountrySecondarySubdivision string
				CountryTertiarySubdivision  string
				Municipality                string // city
				PostalCode                  string
				Country                     string
				CountryCodeISO3             string
				FreeformAddress             string
				CountrySubdivisionName      string
			}
		}
	}
)

// Geocoder constructs TomTom geocoder
func Geocoder(key string, baseURLs ...string) geo.Geocoder {
	_ = "STUB: not implemented"
	return *new(geo.Geocoder)
}

func getUrl(key string, baseURLs ...string) string { _ = "STUB: not implemented"; return "" }

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
