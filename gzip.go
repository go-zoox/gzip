package gzip

import (
	"compress/gzip"

	"github.com/go-zoox/zoox"
)

const (
	// BestCompression is the level of best compression.
	BestCompression = gzip.BestCompression
	// BestSpeed is the level of best speed.
	BestSpeed = gzip.BestSpeed
	// DefaultCompression is the default compression level.
	DefaultCompression = gzip.DefaultCompression
	// NoCompression is the level of no compression.
	NoCompression = gzip.NoCompression
)

// Gzip is a gzip middleware for zoox
func Gzip(level int, options ...Option) zoox.HandlerFunc {
	return newGzipHandler(level, options...).Handler
}

// gzipWriter is a custom writer for gzip response.
type gzipWriter struct {
	zoox.ResponseWriter
	writer *gzip.Writer
}

func (g *gzipWriter) WriteString(s string) (int, error) {
	g.Header().Del("Content-Length")
	return g.writer.Write([]byte(s))
}

func (g *gzipWriter) Write(data []byte) (int, error) {
	g.Header().Del("Content-Length")
	return g.writer.Write(data)
}

// Fix: https://github.com/mholt/caddy/issues/38
func (g *gzipWriter) WriteHeader(code int) {
	g.Header().Del("Content-Length")
	g.ResponseWriter.WriteHeader(code)
}
