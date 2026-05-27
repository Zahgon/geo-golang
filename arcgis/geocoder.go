// Package arcgis is a geo-golang based ArcGIS geocode/reverse client
package arcgis

import (
	geo "github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Candidates []struct {
			Address  string
			Location struct {
				X float64
				Y float64
			}
		}

		ReverseAddress struct {
			MatchAddr    string `json:"Match_addr"`
			LongLabel    string
			ShortLabel   string
			AddNum       string
			Address      string
			Neighborhood string
			City         string
			Subregion    string
			Region       string
			Postal       string
			CountryCode  string
		} `json:"address"`
	}
)

// Geocoder constructs ArcGIS geocoder
func Geocoder(token string, baseURLs ...string) geo.Geocoder {
	_ = "STUB: not implemented"
	return *new(geo.Geocoder)
}

func getUrl(token string, baseURLs ...string) string { _ = "STUB: not implemented"; return "" }

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
