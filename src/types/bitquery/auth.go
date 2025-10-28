// Package auth manages Bitquery GraphQL API authentication and token lifecycle
package bitquery

import "net/http"

type AuthedTransport struct {
	Key     string
	Wrapped http.RoundTripper
}

func (Auth *AuthedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
// TODO: Auto-refresh API tokens before expiration to prevent request failures
	req.Header.Set("X-API-KEY", Auth.Key)
	return Auth.Wrapped.RoundTrip(req)
}