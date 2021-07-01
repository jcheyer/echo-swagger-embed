package echoswaggerembed

import "io"

// URL presents the url pointing to API definition (normally swagger.json or swagger.yaml).
func WithURL(url string) func(s *Swagger) {
	return func(s *Swagger) {
		s.URL = url
	}
}

func WithVersion(version string) func(s *Swagger) {
	return func(s *Swagger) {
		s.Version = version
	}
}

func WithSpecs(i io.Reader) func(s *Swagger) {
	return func(s *Swagger) {
		b, _ := io.ReadAll(i)
		s.Specs = string(b)
	}
}
