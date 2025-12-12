# Sanitizer

A lightweight, extensible Go module for **sanitizing and masking sensitive data**  
built to prevent accidental leakage of secrets, credentials, PII, and financial
information into logs, monitoring systems, or third-party tools.

Designed with **real-world logging safety** in mind.

---

## ✨ Features

- 🔐 Mask authentication data  
  (passwords, tokens, API keys, session IDs, OTPs)
- 💳 Mask financial data  
  (card PAN, CVV, IBAN)
- 🧑‍💼 Mask PII  
  (email, mobile numbers, identifiers)
- 🧠 Rule-based and extensible registry
- 🧬 Recursive sanitization (nested maps & slices)
- 📦 Built-in helpers for:
    - JSON payloads (`[]byte`)
    - Query parameters (`url.Values`)
    - HTTP headers (`http.Header`)
- 🧪 Fully unit-tested (table-driven tests)

---

## 📦 Installation

```bash
go get github.com/imRezaAlie/sanitizer@latest
```

## 🚀 Quick Start
### Sanitize any payload
```go
import "github.com/imRezaAlie/sanitizer/sanitize"

payload := map[string]any{
  "email":    "ali@gmail.com",
  "password": "123456",
  "token":    "eyJhbGciOi...",
}

safe := sanitize.SanitizeAny(payload)
```
### Output
```go
map[string]any{
  "email":    "a***@gmail.com",
  "password": "***",
  "token":    "***",
}

```
🧩 Supported Data Types
1. Any / map / slice (recursive)
```go
sanitize.SanitizeAny(data)
```
2. JSON payloads
```go
safeJSON, err := sanitize.SanitizeJSON(rawJSON)
```
If the JSON is invalid, the original input is returned safely.
4. HTTP Headers

```go
safeHeaders := sanitize.SanitizeHeaders(req.Header)
```
Authorization, cookies, and sensitive headers are masked automatically.

---

## ⚙️ Advanced Usage – Custom Rules
You can define your own registry and rules:
```go
r := sanitize.NewRegistry()
sanitize.RegisterDefaults(r)

r.Register(sanitize.Rule{
  Name:     "custom-secret",
  KeyRegex: regexp.MustCompile(`(?i)secret_value`),
  Action:   sanitize.ActionMask,
})

safe := r.SanitizeAny(payload)

```

---
## 🛡️ What Gets Sanitized by Default
### 🔐 Authentication & Secrets

* password / pwd
* token / jwt / access_token / refresh_token
* api_key / secret_key
* session_id
* otp

### 💳 Financial

* Card number (6 first + 4 last digits)
* CVV / CVC (removed)
* IBAN
### 🧑‍💼 PII
* Email (partial mask)
* Mobile number (digits-only masking)
*  sensitive fields
----

## 🤝 Contributing

Contributions are very welcome ❤️

1. Fork the repository
2. Create a new branch (`feat/...`, `fix/...`)
3. Add tests for new behavior
4. Run `go test ./...`
5. Open a Pull Request

---