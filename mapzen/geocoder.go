// Package mapzen is a geo-golang based Mapzen geocode/reverse client
package mapzen

import (
	geo "github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Geocoding struct {
			Query struct {
				Text string
			}
		}

		Features []struct {
			Geometry struct {
				Coordinates []float64
			}
			Properties struct {
				Name        string
				HouseNumber string
				Street      string
				PostalCode  string
				Country     string
				CountryCode string `json:"country_a"`
				Region      string
				RegionCode  string `json:"region_a"`
				County      string
				Label       string
			}
		}
	}
)

// Geocoder constructs Mapzen geocoder
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
