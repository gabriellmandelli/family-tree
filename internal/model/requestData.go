package model

import (
	"net/http"
	"net/url"
)

type RequestData struct {
	Headers     http.Header
	QueryParams url.Values
}
