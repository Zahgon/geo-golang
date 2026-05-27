package geo

import (
	"context"
	"errors"
	"time"
)

// DefaultTimeout for the request execution
const DefaultTimeout = time.Second * 8

// ErrTimeout occurs when no response returned within timeoutInSeconds
var ErrTimeout = errors.New("TIMEOUT")

// EndpointBuilder defines functions that build urls for geocode/reverse geocode
type EndpointBuilder interface {
	GeocodeURL(string) string
	ReverseGeocodeURL(Location) string
}

// ResponseParserFactory creates a new ResponseParser
type ResponseParserFactory func() ResponseParser

// ResponseParser defines functions that parse response of geocode/reverse geocode
type ResponseParser interface {
	Location() (*Location, error)
	Address() (*Address, error)
}

// HTTPGeocoder has EndpointBuilder and ResponseParser
type HTTPGeocoder struct {
	EndpointBuilder
	ResponseParserFactory
	ResponseUnmarshaler
}

func (g HTTPGeocoder) geocodeWithContext(ctx context.Context, address string) (*Location, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Geocode returns location for address
func (g HTTPGeocoder) Geocode(address string) (*Location, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReverseGeocode returns address for location
func (g HTTPGeocoder) ReverseGeocode(lat, lng float64) (*Address, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ResponseUnmarshaler interface {
	Unmarshal(data []byte, v any) error
}

type JSONUnmarshaler struct{}

func (*JSONUnmarshaler) Unmarshal(data []byte, v any) error { _ = "STUB: not implemented"; return nil }

type XMLUnmarshaler struct{}

func (*XMLUnmarshaler) Unmarshal(data []byte, v any) error { _ = "STUB: not implemented"; return nil }

// Response gets response from url
func response(ctx context.Context, url string, unmarshaler ResponseUnmarshaler, obj ResponseParser) error {
	_ = "STUB: not implemented"
	return nil
}

// ParseFloat is a helper to parse a string to a float
func ParseFloat(value string) float64 { _ = "STUB: not implemented"; return 0 }
