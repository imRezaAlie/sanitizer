package sanitize

import (
	"strings"
	"unicode"
)

func MaskMobile(num string) string {
	d := digitsOnly(num)
	if len(d) < 7 {
		return num
	}

	first := d[:5]
	last := d[len(d)-3:]
	masked := first + "***" + last

	// اگر ورودی با + شروع شده و پیش‌شماره 98 بود، + رو نگه دار
	if strings.HasPrefix(strings.TrimSpace(num), "+") && strings.HasPrefix(d, "98") {
		return "+" + masked
	}
	return masked
}

func MaskEmail(email string) string {
	at := strings.Index(email, "@")
	if at <= 1 {
		return email
	}
	local := email[:at]
	domain := email[at:]

	if len(local) <= 3 {
		return local[:1] + "***" + domain
	}
	return local[:3] + "***" + domain
}

func MaskCardPan(pan string) string {
	digits := make([]rune, 0, len(pan))
	for _, r := range pan {
		if unicode.IsDigit(r) {
			digits = append(digits, r)
		}
	}
	if len(digits) < 10 {
		return "***"
	}
	first6 := string(digits[:6])
	last4 := string(digits[len(digits)-4:])
	return first6 + "******" + last4
}

// digitsOnly extracts only numeric characters from a string
func digitsOnly(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		if unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}
