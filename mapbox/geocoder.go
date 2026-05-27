// Package mapbox is a geo-golang based Mapbox geocode/reverse geocode client
package mapbox

import (
	"encoding/json"

	"github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Features []struct {
			PlaceName string `json:"place_name"`
			Center    [2]float64
			Text      string          `json:"text"`    // usually street name
			Address   json.RawMessage `json:"address"` // potentially house number
			Context   []struct {
				Text      string `json:"text"`
				Id        string `json:"id"`
				ShortCode string `json:"short_code"`
				Wikidata  string `json:"wikidata"`
			}
		}
		Message string `json:"message"`
	}
)

const (
	mapboxPrefixLocality = "place"
	mapboxPrefixPostcode = "postcode"
	mapboxPrefixState    = "region"
	mapboxPrefixCountry  = "country"
)

// Geocoder constructs Mapbox geocoder
func Geocoder(token string, baseURLs ...string) geo.Geocoder {
	_ = "STUB: not implemented"
	return *new(geo.Geocoder)
}

func getURL(token string, baseURLs ...string) string { _ = "STUB: not implemented"; return "" }

func (b baseURL) GeocodeURL(address string) string { _ = "STUB: not implemented"; return "" }

func (b baseURL) ReverseGeocodeURL(l geo.Location) string { _ = "STUB: not implemented"; return "" }

func (r *geocodeResponse) Location() (*geo.Location, error) {
	_ = "STUB: not implemented"
	return nil,

		// error in response
		nil
}

// no results

func (r *geocodeResponse) Address() (*geo.Address, error) {
	_ = "STUB: not implemented"
	return nil,

		// error in response
		nil
}

// no results

func parseMapboxResponse(r *geocodeResponse) *geo.Address { _ = "STUB: not implemented"; return nil }
