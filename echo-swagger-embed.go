package echoswaggerembed

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Option func(s *Swagger)

type Swagger struct {
	index       *template.Template
	Specs       string
	URL         string
	Version     string
	DeepLinking bool
	TryItOut    bool
}

func New(options ...Option) *Swagger {
	swagger := &Swagger{
		URL:         "doc.json",
		Version:     "3.51.0",
		Specs:       "",
		DeepLinking: true,
		TryItOut:    false,
	}

	for _, o := range options {
		o(swagger)
	}

	t := template.New("swagger_index.html")
	swagger.index, _ = t.Parse(indexTemplate)

	return swagger
}

func (s *Swagger) Handle(c echo.Context) error {

	param := c.ParamValues()
	if len(param) != 1 {
		return c.String(http.StatusBadRequest, "no param")
	}

	switch param[0] {
	case "", "index.html":
		return s.index.Execute(c.Response().Writer, s)
	case s.URL:
		if s.Specs == "" {
			return c.String(http.StatusNotFound, "not found")
		}
		return c.String(http.StatusOK, s.Specs)
	}

	return c.String(http.StatusNotFound, "not found")

}

const indexTemplate = `<!-- HTML for static distribution bundle build -->
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <title>Swagger UI</title>
    <link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/{{.Version}}/swagger-ui.css" />
    <!--<link rel="icon" type="image/png" href="./favicon-32x32.png" sizes="32x32" />-->
    <!--<link rel="icon" type="image/png" href="./favicon-16x16.png" sizes="16x16" />-->
    <style>
      html
      {
        box-sizing: border-box;
        overflow: -moz-scrollbars-vertical;
        overflow-y: scroll;
      }

      *,
      *:before,
      *:after
      {
        box-sizing: inherit;
      }

      body
      {
        margin:0;
        background: #fafafa;
      }
    </style>
  </head>

  <body>
    <div id="swagger-ui"></div>

    <script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/{{.Version}}/swagger-ui-bundle.js" charset="UTF-8"> </script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/{{.Version}}/swagger-ui-standalone-preset.js" charset="UTF-8"> </script>
    <script>
    window.onload = function() {
      // Begin Swagger UI call region
      const ui = SwaggerUIBundle({
        url: "{{.URL}}",
        dom_id: '#swagger-ui',
        deepLinking:{{.DeepLinking}},
        tryItOutEnabled:{{.TryItOut}},
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIStandalonePreset
        ],
        plugins: [
          SwaggerUIBundle.plugins.DownloadUrl
        ],
        layout: "StandaloneLayout"
      });
      // End Swagger UI call region

      window.ui = ui;
    };
  </script>
  </body>
</html>
`
