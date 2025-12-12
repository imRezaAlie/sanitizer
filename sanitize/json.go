package sanitize

import "encoding/json"

func (r *Registry) SanitizeJSON(raw []byte) ([]byte, error) {
	if len(raw) == 0 {
		return raw, nil
	}

	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return raw, nil
	}

	safe := r.SanitizeAny("", v)
	return json.Marshal(safe)
}
