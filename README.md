# echo-swagger-embed


echo handler to embed swagger-ui

Current default swagger-ui version used: 3.51.0

## Usage

### Start using it

1. Download [echo-swagger-embed](https://github.com/jhcheyer/echo-swagger-embed) by using:
```sh
$ go get -u github.com/jcheyer/echo-swagger-embed
```

2. Import following in your code:
```go
import echoswaggerembed "github.com/jcheyer/echo-swagger-embed"
```

### Example:

```go
package main

import (
    _ "embed"
    "github.com/labstack/echo/v4"
    echoswaggerembed "github.com/jcheyer/echo-swagger-embed"
)

//go:embed swagger.yaml
var specs []byte


func main() {
    e := echo.New()

    r := bytes.NewReader(specs)
    swagger := echoswaggerembed.New(echoswaggerembed.WithSpecs(r))

    e.GET("/swagger/*", swagger.Handle)

    e.Logger.Fatal(e.Start(":8080"))
}

```

3. have a swagger.yaml file in place (see //go:embed swagger.yaml)

4. Run it, and browser to http://localhost:8080/swagger/index.html

### Options

use remote specs file
```
WithURL("https://your.domain/swagger.yaml")
```

use different versions of swagger-ui
```
WithVersion("3.19.4")
```
