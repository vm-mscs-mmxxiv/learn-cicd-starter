package auth

import (
    "testing"
	"net/http"
)

func TestGetAPIKey(t *testing.T) {
    var authTests = []struct {
        header http.Header
        want string
    }{
        {
            header: http.Header{"Authorization": []string{"ApiKey key-xyz"}},
            want: "key-xyz",
        },
        {
            header: http.Header{"Authorization": []string{"ApiKey aasdfsaf-adsaf-asdfsa"}},
            want: "aasdfsaf-adsaf-asdfsa",
        },
        {
            header: http.Header{"Authorization": []string{"ApiKey 1234123345345"}},
            want: "1234123345345",
        },
        {
            header: http.Header{"Authorization": []string{"ApiKey key-2134234-sasdf"}},
            want: "key-2134234-sasdf",
        },
        {
            header: http.Header{"Authorization": []string{"ApiKey key-xyz-sadfadsfsafs"}},
            want: "key-xyz-sadfadsfsafs",
        },
    }

	for _, tt := range authTests {
        got, err := GetAPIKey(tt.header)
        if err != nil {
            t.Fatalf("something went wrong: %v", err)
        }
        if got == "" {
            t.Fatalf("expected Authorization header to be present, got empty")
        }
        if got != tt.want {
            t.Fatalf("unexpected Authorization value: got %q want %q", got, tt.want)
        }
	}
}
