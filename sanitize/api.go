package sanitize

import (
	"net/http"
	"net/url"
)

func SanitizeAny(v any) any {
	return DefaultRegistry.SanitizeAny("", v)
}

func SanitizeJSON(raw []byte) ([]byte, error) {
	return DefaultRegistry.SanitizeJSON(raw)
}

func SanitizeQuery(q url.Values) map[string]any {
	return DefaultRegistry.SanitizeQuery(q)
}

func SanitizeHeaders(h http.Header) map[string]any {
	return DefaultRegistry.SanitizeHeaders(h)
}
