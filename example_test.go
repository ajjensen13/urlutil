package urlutil_test

import (
	"fmt"
	"github.com/ajjensen13/urlutil"
	"net/url"
)

func ExampleRewrite() {
	rw := urlutil.Rewriter("https://${host}/root${request_uri}")
	in, err := url.Parse("http://example.com")
	if err != nil {
		panic(err)
	}
	rewritten, err := rw.Rewrite(in)
	if err != nil {
		panic(err)
	}
	fmt.Print(rewritten.String())

	// Output:
	// https://example.com/root/
}
