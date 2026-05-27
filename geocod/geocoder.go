package geocod

import (
	geo "github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Results []struct {
			Components struct {
				Number  string
				Street  string
				City    string
				County  string
				State   string
				Zip     string
				Country string
			} `json:"address_components"`
			Address  string `json:"formatted_address"`
			Location struct {
				Lat float64
				Lng float64
			}
		}
	}
)

// Geocoder constructs Geocodio geocoder
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
