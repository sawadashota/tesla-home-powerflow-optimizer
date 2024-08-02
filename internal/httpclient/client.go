package httpclient

import "net/http"

type (
	Provider interface {
		HTTPClient() *http.Client
	}
)
