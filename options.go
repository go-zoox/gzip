package gzip

import (
	"compress/gzip"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-zoox/zoox"
)

var (
	// DefaultExcludedExtentions is the default excluded extensions.
	DefaultExcludedExtentions = NewExcludedExtensions([]string{
		".png", ".gif", ".jpeg", ".jpg",
	})
	// DefaultOptions is the default options for gzip middleware.
	DefaultOptions = &Options{
		ExcludedExtensions: DefaultExcludedExtentions,
	}
)

// Options is the options for gzip middleware.
type Options struct {
	ExcludedExtensions   ExcludedExtensions
	ExcludedPaths        ExcludedPaths
	ExcludedPathesRegexs ExcludedPathesRegexs
	DecompressFn         func(c *zoox.Context)
}

// Option is the type of function that can be passed to WithExcludedExtensions, WithExcludedPaths, WithExcludedPathsRegexs, WithDecompressFn.
type Option func(*Options)

// WithExcludedExtensions is an option for excluding compression for certain file extensions.
func WithExcludedExtensions(args []string) Option {
	return func(o *Options) {
		o.ExcludedExtensions = NewExcludedExtensions(args)
	}
}

// WithExcludedPaths is an option for excluding compression for certain paths.
func WithExcludedPaths(args []string) Option {
	return func(o *Options) {
		o.ExcludedPaths = NewExcludedPaths(args)
	}
}

// WithExcludedPathsRegexs is an option for excluding compression for certain paths.
func WithExcludedPathsRegexs(args []string) Option {
	return func(o *Options) {
		o.ExcludedPathesRegexs = NewExcludedPathesRegexs(args)
	}
}

// WithDecompressFn is an option for setting a custom decompress function.
func WithDecompressFn(decompressFn func(c *zoox.Context)) Option {
	return func(o *Options) {
		o.DecompressFn = decompressFn
	}
}

// ExcludedExtensions is extensions map, using map for better lookup performance
type ExcludedExtensions map[string]bool

// NewExcludedExtensions creates a new ExcludedExtensions map.
func NewExcludedExtensions(extensions []string) ExcludedExtensions {
	res := make(ExcludedExtensions)
	for _, e := range extensions {
		res[e] = true
	}
	return res
}

// Contains returns true if the extension is excluded.
func (e ExcludedExtensions) Contains(target string) bool {
	_, ok := e[target]
	return ok
}

// ExcludedPaths ...
type ExcludedPaths []string

// NewExcludedPaths creates a new ExcludedPaths.
func NewExcludedPaths(paths []string) ExcludedPaths {
	return ExcludedPaths(paths)
}

// Contains returns true if the path is excluded.
func (e ExcludedPaths) Contains(requestURI string) bool {
	for _, path := range e {
		if strings.HasPrefix(requestURI, path) {
			return true
		}
	}
	return false
}

// ExcludedPathesRegexs ...
type ExcludedPathesRegexs []*regexp.Regexp

// NewExcludedPathesRegexs creates a new ExcludedPathesRegexs.
func NewExcludedPathesRegexs(regexs []string) ExcludedPathesRegexs {
	result := make([]*regexp.Regexp, len(regexs))
	for i, reg := range regexs {
		result[i] = regexp.MustCompile(reg)
	}
	return result
}

// Contains returns true if the path is excluded.
func (e ExcludedPathesRegexs) Contains(requestURI string) bool {
	for _, reg := range e {
		if reg.MatchString(requestURI) {
			return true
		}
	}
	return false
}

// DefaultDecompressHandle is the default decompress handle.
func DefaultDecompressHandle(c *zoox.Context) {
	if c.Request.Body == nil {
		return
	}

	r, err := gzip.NewReader(c.Request.Body)
	if err != nil {
		// _ = c.AbortWithError(http.StatusBadRequest, err)
		c.Fail(err, http.StatusBadRequest, "Bad Request", http.StatusBadRequest)
		return
	}

	c.Request.Header.Del("Content-Encoding")
	c.Request.Header.Del("Content-Length")
	c.Request.Body = r
}
