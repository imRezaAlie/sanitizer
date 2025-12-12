package sanitize

import "regexp"

var DefaultRegistry = NewRegistry()

func init() {
	RegisterDefaults(DefaultRegistry)
}

func RegisterDefaults(r *Registry) {
	// Auth
	r.Register(Rule{
		Name:     "password",
		KeyRegex: regexp.MustCompile(`(?i)\b(pass|password|pwd)\b`),
		Action:   ActionMask,
		Mask:     func(_ string, _ any) (any, bool) { return "***", true },
	})

	r.Register(Rule{
		Name:     "token",
		KeyRegex: regexp.MustCompile(`(?i)\b(token|jwt|access_token|refresh_token|authorization)\b`),
		Action:   ActionMask,
		Mask:     func(_ string, _ any) (any, bool) { return "***", true },
	})

	r.Register(Rule{
		Name:     "otp",
		KeyRegex: regexp.MustCompile(`(?i)\b(otp|one[_-]?time|code)\b`),
		Action:   ActionMask,
		Mask:     func(_ string, _ any) (any, bool) { return "***", true },
	})

	r.Register(Rule{
		Name:     "keys",
		KeyRegex: regexp.MustCompile(`(?i)\b(api[_-]?key|secret|secret[_-]?key|private[_-]?key|session[_-]?id)\b`),
		Action:   ActionMask,
		Mask:     func(_ string, _ any) (any, bool) { return "***", true },
	})

	// Financial
	r.Register(Rule{
		Name:     "cvv",
		KeyRegex: regexp.MustCompile(`(?i)\b(cvv2?|cvc)\b`),
		Action:   ActionRemove,
		Mask:     func(_ string, _ any) (any, bool) { return "", true },
	})

	r.Register(Rule{
		Name:     "card_pan_by_key",
		KeyRegex: regexp.MustCompile(`(?i)\b(card|pan|card_number|cardnumber)\b`),
		Action:   ActionMask,
		Mask: func(_ string, v any) (any, bool) {
			s, ok := v.(string)
			if !ok {
				return "***", true
			}
			return MaskCardPan(s), true
		},
	})

	// PII
	r.Register(Rule{
		Name:     "mobile",
		KeyRegex: regexp.MustCompile(`(?i)\b(mobile|phone|cell|msisdn)\b`),
		Action:   ActionMask,
		Mask: func(_ string, v any) (any, bool) {
			s, ok := v.(string)
			if !ok {
				return "***", true
			}
			return MaskMobile(s), true
		},
	})

	r.Register(Rule{
		Name:     "email",
		KeyRegex: regexp.MustCompile(`(?i)\b(email|mail)\b`),
		Action:   ActionMask,
		Mask: func(_ string, v any) (any, bool) {
			s, ok := v.(string)
			if !ok {
				return "***", true
			}
			return MaskEmail(s), true
		},
	})
	// IBAN (مثلاً IR.... یا DE....)
	r.Register(Rule{
		Name:       "iban_by_value",
		ValueRegex: regexp.MustCompile(`(?i)\b[A-Z]{2}\d{2}[A-Z0-9]{10,30}\b`),
		Action:     ActionMask,
		Mask: func(_ string, v any) (any, bool) {
			s, ok := v.(string)
			if !ok || len(s) < 8 {
				return "***", true
			}
			// 4 اول + ... + 4 آخر
			if len(s) <= 10 {
				return "***", true
			}
			return s[:4] + "****" + s[len(s)-4:], true
		},
	})

	// کارت 16 رقمی در هر جایی
	r.Register(Rule{
		Name:       "card_pan_by_value_16",
		ValueRegex: regexp.MustCompile(`\b\d{16}\b`),
		Action:     ActionMask,
		Mask: func(_ string, v any) (any, bool) {
			s, ok := v.(string)
			if !ok {
				return "***", true
			}
			return MaskCardPan(s), true
		},
	})

}
