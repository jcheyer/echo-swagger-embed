package echoswaggerembed

import "io"

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

func WithTryItOut(b bool) func(s *Swagger) {
	return func(s *Swagger) {
		s.TryItOut = b
	}
}

func WithDeepLinking(b bool) func(s *Swagger) {
	return func(s *Swagger) {
		s.DeepLinking = b
	}
}
