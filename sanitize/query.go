package sanitize

import "net/url"

func (r *Registry) SanitizeQuery(q url.Values) map[string]any {
	out := make(map[string]any, len(q))
	for k, vals := range q {
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
