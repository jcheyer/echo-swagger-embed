package echoswaggerembed

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionTryItOut(t *testing.T) {
	s := &Swagger{}

	assert.False(t, s.TryItOut)
	WithTryItOut(true)(s)
	assert.True(t, s.TryItOut)
}

func TestOptionDeeplinking(t *testing.T) {
	s := &Swagger{}

	assert.False(t, s.DeepLinking)
	WithDeepLinking(true)(s)
	assert.True(t, s.DeepLinking)
}

func TestOptionSpecs(t *testing.T) {
	s := &Swagger{}

	assert.Equal(t, "", s.Specs)
	testString := "SpecsDummy"

	WithSpecs(strings.NewReader(testString))(s)
	assert.Equal(t, testString, s.Specs)
}

func TestOptionURL(t *testing.T) {
	s := &Swagger{}

	assert.Equal(t, "", s.URL)
	testUrl := "https://my.domain/swagger.yaml"

	WithURL(testUrl)(s)
	assert.Equal(t, testUrl, s.URL)
}

func TestOptionVersion(t *testing.T) {
	s := &Swagger{}

	assert.Equal(t, "", s.Version)
	testVersion := "1.2.3.4"

	WithVersion(testVersion)(s)
	assert.Equal(t, testVersion, s.Version)
}
