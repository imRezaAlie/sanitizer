package sanitize

import (
	"net/url"
	"testing"
)

func TestSanitizeQuery(t *testing.T) {
	r := NewRegistry()
	RegisterDefaults(r)

	q := url.Values{}
	q.Set("email", "ali@gmail.com")
	q.Set("password", "123456")
	q.Set("token", "ey...")
	q.Add("mobile", "09123456789")
	q.Add("mobile", "09120001111")

	out := r.SanitizeQuery(q)

	if out["email"] != "a***@gmail.com" {
		t.Fatalf("email=%v; want a***@gmail.com", out["email"])
	}
	if out["password"] != "***" {
		t.Fatalf("password=%v; want ***", out["password"])
	}
	if out["token"] != "***" {
		t.Fatalf("token=%v; want ***", out["token"])
	}

	mobiles, ok := out["mobile"].([]any)
	if !ok || len(mobiles) != 2 {
		t.Fatalf("mobile=%T/%v; want []any len=2", out["mobile"], out["mobile"])
	}
	if mobiles[0] != "09123***789" {
		t.Fatalf("mobile[0]=%v; want 09123***789", mobiles[0])
	}
	if mobiles[1] != "09120***111" {
		t.Fatalf("mobile[1]=%v; want 09120***111", mobiles[1])
	}
}
