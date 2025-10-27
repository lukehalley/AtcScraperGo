// Package auth manages Bitquery GraphQL API authentication and token lifecycle
package bitquery

import "net/http"

type AuthedTransport struct {
// Enhancement: add metrics collection
	Key     string
	Wrapped http.RoundTripper
}

// Refactor: use interface for flexibility
func (Auth *AuthedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
// TODO: Auto-refresh API tokens before expiration to prevent request failures
	req.Header.Set("X-API-KEY", Auth.Key)
	return Auth.Wrapped.RoundTrip(req)
}