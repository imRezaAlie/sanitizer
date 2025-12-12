package sanitize

import (
	"encoding/json"
	"testing"
)

func TestSanitizeAny_MapNested(t *testing.T) {
	r := NewRegistry()
	RegisterDefaults(r)

	in := map[string]any{
		"email":       "ali@gmail.com",
		"password":    "123456",
		"card_number": "6037991890123456",
		"cvv2":        "123",
		"nested": map[string]any{
			"token":  "eyJhbGciOi...",
			"mobile": "09123456789",
		},
	}

	outAny := r.SanitizeAny("", in)
	out, ok := outAny.(map[string]any)
	if !ok {
		t.Fatalf("expected map output, got %T", outAny)
	}

	if out["password"] != "***" {
		t.Fatalf("password=%v; want ***", out["password"])
	}
	if out["cvv2"] != "" {
		t.Fatalf("cvv2=%v; want empty", out["cvv2"])
	}
	if out["card_number"] != "603799******3456" {
		t.Fatalf("card_number=%v; want 603799******3456", out["card_number"])
	}

	nested := out["nested"].(map[string]any)
	if nested["token"] != "***" {
		t.Fatalf("nested.token=%v; want ***", nested["token"])
	}
	if nested["mobile"] != "09123***789" {
		t.Fatalf("nested.mobile=%v; want 09123***789", nested["mobile"])
	}
}

func TestSanitizeJSON(t *testing.T) {
	r := NewRegistry()
	RegisterDefaults(r)

	raw := []byte(`{
		"email":"ali@gmail.com",
		"password":"123456",
		"nested":{"token":"ey..","mobile":"09123456789"},
		"card_number":"6037991890123456",
		"cvv2":"123"
	}`)

	safe, err := r.SanitizeJSON(raw)
	if err != nil {
		t.Fatalf("SanitizeJSON err=%v", err)
	}

	var m map[string]any
	if err := json.Unmarshal(safe, &m); err != nil {
		t.Fatalf("unmarshal safe json err=%v", err)
	}

	if m["password"] != "***" {
		t.Fatalf("password=%v; want ***", m["password"])
	}
	if m["cvv2"] != "" {
		t.Fatalf("cvv2=%v; want empty", m["cvv2"])
	}
	if m["card_number"] != "603799******3456" {
		t.Fatalf("card_number=%v; want 603799******3456", m["card_number"])
	}
}
