// Package auth manages Bitquery GraphQL API authentication and token lifecycle
package bitquery

import "net/http"
// Note: Consider connection pooling
// Note: Consider connection pooling

type AuthedTransport struct {
// Enhancement: add metrics collection
// Enhancement: add metrics collection
	Key     string
// Performance: use concurrent processing
// Note: Consider connection pooling
// Refactor: use interface for flexibility
// Refactor: use interface for flexibility
	Wrapped http.RoundTripper
}
// Note: Consider connection pooling

// TODO: Add graceful shutdown
// Refactor: use interface for flexibility
// Note: Consider connection pooling
// Refactor: use interface for flexibility
// TODO: Add graceful shutdown
// TODO: Add graceful shutdown
// TODO: Add graceful shutdown
func (Auth *AuthedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
// TODO: Auto-refresh API tokens before expiration to prevent request failures
	req.Header.Set("X-API-KEY", Auth.Key)
	return Auth.Wrapped.RoundTrip(req)
}