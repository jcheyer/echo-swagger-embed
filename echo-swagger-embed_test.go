package echoswaggerembed

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/labstack/echo/v4"
)

func TestOperations(t *testing.T) {
	handler := echo.New()

	s := New()
	handler.GET("/swagger/*", s.Handle)

	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	e.GET("/swagger/egal").
		Expect().
		Status(http.StatusNotFound)

	e.GET("/swagger/" + s.URL).Expect().Status(http.StatusNotFound)

	testBody := "Hallo"
	s.Specs = testBody

	e.GET("/swagger/" + s.URL).Expect().Status(http.StatusOK).Body().Equal(testBody)

	s.DeepLinking = true
	e.GET("/swagger/").Expect().Status(http.StatusOK).Body().Contains("deepLinking: true ,")

	s.DeepLinking = false
	e.GET("/swagger/").Expect().Status(http.StatusOK).Body().Contains("deepLinking: false ,")

	s.TryItOut = true
	e.GET("/swagger/").Expect().Status(http.StatusOK).Body().Contains("tryItOutEnabled: true ,")

	s.TryItOut = false
	e.GET("/swagger/").Expect().Status(http.StatusOK).Body().Contains("tryItOutEnabled: false ,")

	s.Version = "XX.XX.XX"
	e.GET("/swagger/").Expect().Status(http.StatusOK).Body().Contains("/XX.XX.XX/swagger-ui-bundle.js")

	// TODO: fix this test
	//s.URL = "https://my.domain.com/swagger.yaml"
	//e.GET("/swagger/index.html").Expect().Status(http.StatusOK).Body().Contains("url: \"https://my.domain.com/swagger.yaml\",")
}
