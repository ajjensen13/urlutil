package urlutil

import (
	"net/url"
	"testing"
)

func TestRewriter_Rewrite(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name       string
		r          Rewriter
		args       args
		wantResult string
		wantErr    bool
	}{
		{
			"no variables",
			Rewriter("http://example.com"),
			args{in: "http://test.net/path?query=true#fragment"},
			"http://example.com",
			false,
		},
		{
			"to https",
			Rewriter("https://${host}${request_uri}"),
			args{in: "http://test.net/path?query=true"},
			"https://test.net/path?query=true",
			false,
		},
		{
			"to host",
			Rewriter("https://newhost.com${request_uri}"),
			args{in: "http://test.net/path?query=true"},
			"https://newhost.com/path?query=true",
			false,
		},
		{
			"to subpath",
			Rewriter("https://${host}/parent${request_uri}"),
			args{in: "http://test.net/path?query=true"},
			"https://test.net/parent/path?query=true",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in, err := url.Parse(tt.args.in)
			if err != nil {
				t.Error(err)
			}

			gotResult, err := tt.r.Rewrite(in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rewrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotResultString := gotResult.String(); gotResultString != tt.wantResult {
				t.Errorf("Rewrite() gotResult = %v, want %v", gotResultString, tt.wantResult)
			}
		})
	}
}
