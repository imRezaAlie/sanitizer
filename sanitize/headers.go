package sanitize

import "net/http"

func (r *Registry) SanitizeHeaders(h http.Header) map[string]any {
	out := make(map[string]any, len(h))
	for k, vals := range h {
		// Header key خودش مهمه (Authorization, Cookie, ...)
		if len(vals) == 1 {
			out[k] = r.SanitizeAny(k, vals[0])
		} else {
			arr := make([]any, 0, len(vals))
			for _, v := range vals {
				arr = append(arr, r.SanitizeAny(k, v))
			}
			out[k] = arr
		}
	}
	return out
}
