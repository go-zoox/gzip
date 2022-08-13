# GZIP zoox's middleware

[![Run Tests](https://github.com/go-zoox/gzip/actions/workflows/go.yml/badge.svg)](https://github.com/go-zoox/gzip/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/go-zoox/gzip/branch/master/graph/badge.svg)](https://codecov.io/gh/go-zoox/gzip)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/gzip)](https://goreportcard.com/report/github.com/go-zoox/gzip)
[![GoDoc](https://godoc.org/github.com/go-zoox/gzip?status.svg)](https://godoc.org/github.com/go-zoox/gzip)
[![Join the chat at https://gitter.im/go-zoox/zoox](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/go-zoox/zoox)

Zoox middleware to enable `GZIP` support.

## Usage

Download and install it:

```sh
go get github.com/go-zoox/gzip
```

Import it in your code:

```go
import "github.com/go-zoox/gzip"
```

Canonical example:

```go
package main

import (
  "fmt"
  "net/http"
  "time"

  "github.com/go-zoox/gzip"
  "github.com/go-zoox/zoox"
)

func main() {
  r := zoox.New()
  r.Use(gzip.Gzip(gzip.DefaultCompression))
  r.Get("/ping", func(c *zoox.Context) {
    c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // Listen and Server in 0.0.0.0:8080
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

Customized Excluded Extensions

```go
package main

import (
  "fmt"
  "net/http"
  "time"

  "github.com/go-zoox/gzip"
  "github.com/go-zoox/zoox"
)

func main() {
  r := zoox.Default()
  r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{".pdf", ".mp4"})))
  r.Get("/ping", func(c *zoox.Context) {
    c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // Listen and Server in 0.0.0.0:8080
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

Customized Excluded Paths

```go
package main

import (
  "fmt"
  "net/http"
  "time"

  "github.com/go-zoox/gzip"
  "github.com/go-zoox/zoox"
)

func main() {
  r := zoox.Default()
  r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{"/api/"})))
  r.Get("/ping", func(c *zoox.Context) {
    c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // Listen and Server in 0.0.0.0:8080
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

Customized Excluded Paths

```go
package main

import (
  "fmt"
  "net/http"
  "time"

  "github.com/go-zoox/gzip"
  "github.com/go-zoox/zoox"
)

func main() {
  r := zoox.Default()
  r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPathsRegexs([]string{".*"})))
  r.Get("/ping", func(c *zoox.Context) {
    c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
  })

  // Listen and Server in 0.0.0.0:8080
  if err := r.Run(":8080"); err != nil {
    log.Fatal(err)
  }
}
```

## Thanks

- This project is modified from [gin-gzip](https://github.com/gin-gonic/gin).

## License

GoZoox is released under the [MIT License](./LICENSE).
