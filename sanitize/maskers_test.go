package sanitize

import "testing"

func TestMaskEmail(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"a@gmail.com", "a@gmail.com"},            // at<=1 => بدون تغییر
		{"ab@gmail.com", "a***@gmail.com"},        // local کوتاه
		{"ali@gmail.com", "a***@gmail.com"},       // len<=3
		{"reza.ab@gmail.com", "rez***@gmail.com"}, // معمولی
	}

	for _, tt := range tests {
		got := MaskEmail(tt.in)
		if got != tt.want {
			t.Fatalf("MaskEmail(%q)=%q; want %q", tt.in, got, tt.want)
		}
	}
}

func TestMaskMobile(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"123", "123"},
		{"09123456789", "09123***789"},
		{"+989123456789", "+9891***789"}, // چون len و slicing مستقیمه، خروجی بر اساس رشته خامه
	}

	for _, tt := range tests {
		got := MaskMobile(tt.in)
		if got != tt.want {
			t.Fatalf("MaskMobile(%q)=%q; want %q", tt.in, got, tt.want)
		}
	}
}

func TestMaskCardPan(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"6037991890123456", "603799******3456"},
		{"6037-9918-9012-3456", "603799******3456"},
		{"123456789", "***"}, // کوتاه
	}

	for _, tt := range tests {
		got := MaskCardPan(tt.in)
		if got != tt.want {
			t.Fatalf("MaskCardPan(%q)=%q; want %q", tt.in, got, tt.want)
		}
	}
}
