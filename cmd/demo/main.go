package main

import (
	"fmt"
	"github.com/imRezaAlie/sanitizer/sanitize"
)

func main() {
	payload := map[string]any{
		"email":       "ali@gmail.com",
		"password":    "123456",
		"card_number": "6037991890123456",
		"cvv2":        "123",
		"nested": map[string]any{
			"token":  "eyJhbGciOi...",
			"mobile": "09123456789",
		},
	}

	safe := sanitize.DefaultRegistry.SanitizeAny("", payload)
	fmt.Printf("%#v\n", safe)
}
