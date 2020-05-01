package urlutil

import (
	"net/url"
	"os"
	"strings"
)

// Rewriter is a URL re-writer based on a template. See Rewrite() for
// details on templates.
type Rewriter string

// Rewrite takes an input URL and re-writes it to conform to a template. A template
// is a string with environment variables that get replaced from the in url.
//
// The set of variables are scheme, host, path, query, and request_uri.
func (r Rewriter) Rewrite(in *url.URL) (result *url.URL, err error) {
	target := os.Expand(string(r), func(s string) string {
		switch strings.ToLower(s) {
		case "scheme":
			return in.Scheme
		case "host":
			return in.Host
		case "path":
			return in.Path
		case "query":
			return in.RawQuery
		case "request_uri":
			return in.RequestURI()
		default:
			return ""
		}
	})

	return url.Parse(target)
}
