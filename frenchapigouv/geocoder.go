// Package frenchapigouv is a geo-golang based French API Gouv geocode/reverse geocode client
package frenchapigouv

import (
	"github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Type     string
		Version  string
		Features []struct {
			Type     string
			Geometry struct {
				Type        string
				Coordinates []float64
			}
			Properties struct {
				Label       string
				Score       float64
				Housenumber string
				Citycode    string
				Context     string
				Postcode    string
				Name        string
				ID          string
				Y           float64
				Importance  float64
				Type        string
				City        string
				X           float64
				Street      string
			}
		}
		Attribution string
		Licence     string
		Query       string
		Limit       int
	}
	context struct {
		state      string
		county     string
		countyCode string
	}
)

// Geocoder constructs FrenchApiGouv geocoder
func Geocoder() geo.Geocoder { _ = "STUB: not implemented"; return *new(geo.Geocoder) }

// GeocoderWithURL constructs French API Gouv geocoder using a custom installation of Nominatim
func GeocoderWithURL(url string) geo.Geocoder { _ = "STUB: not implemented"; return *new(geo.Geocoder) }

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

func (r *geocodeResponse) parseContext() *context { _ = "STUB: not implemented"; return nil }
