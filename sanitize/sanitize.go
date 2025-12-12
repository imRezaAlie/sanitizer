package sanitize

func (r *Registry) SanitizeAny(key string, v any) any {
	if out, ok := r.Apply(key, v); ok {
		return out
	}

	switch x := v.(type) {
	case map[string]any:
		out := make(map[string]any, len(x))
		for k, vv := range x {
			out[k] = r.SanitizeAny(k, vv)
		}
		return out

	case []any:
		out := make([]any, len(x))
		for i := range x {
			out[i] = r.SanitizeAny("", x[i])
		}
		return out

	default:
		return v
	}
}
