package sanitize

import (
	"regexp"
	"strings"
	"sync"
)

type Action int

const (
	ActionMask Action = iota
	ActionRemove
)

type MaskFunc func(key string, value any) (any, bool)

type Rule struct {
	Name       string
	KeyRegex   *regexp.Regexp
	ValueRegex *regexp.Regexp
	Action     Action
	Mask       MaskFunc
}

type Registry struct {
	mu    sync.RWMutex
	rules []Rule
}

func NewRegistry() *Registry { return &Registry{} }

func (r *Registry) Register(rule Rule) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.rules = append(r.rules, rule)
}

func (r *Registry) Apply(key string, val any) (any, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	k := strings.ToLower(strings.TrimSpace(key))

	for _, rule := range r.rules {
		if rule.KeyRegex != nil && rule.KeyRegex.MatchString(k) {
			if rule.Mask != nil {
				return rule.Mask(key, val)
			}
			return defaultAction(rule.Action), true
		}
		if rule.ValueRegex != nil {
			s, ok := val.(string)
			if ok && rule.ValueRegex.MatchString(s) {
				if rule.Mask != nil {
					return rule.Mask(key, val)
				}
				return defaultAction(rule.Action), true
			}
		}
	}
	return val, false
}

func defaultAction(a Action) any {
	if a == ActionRemove {
		return ""
	}
	return "***"
}
