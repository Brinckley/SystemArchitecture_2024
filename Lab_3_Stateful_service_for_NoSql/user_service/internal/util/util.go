package util

import (
	"net/http"
)

func CopyHeadersToRequest(req *http.Request, r *http.Request) {
	for name, values := range r.Header {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}
}

func CopyHeadersToWriter(resp *http.Response, writer http.ResponseWriter) {
	for name, values := range resp.Header {
		for _, value := range values {
			writer.Header().Add(name, value)
		}
	}
}

func CreateProxyRequest(req *http.Request, url string) (*http.Request, error) {
	proxyReq, err := http.NewRequest(req.Method, url, req.Body)
	if err != nil {
		return nil, err
	}
	CopyHeadersToRequest(proxyReq, req)
	return proxyReq, nil
}
