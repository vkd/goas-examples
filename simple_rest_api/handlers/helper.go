package handlers

import "net/url"

// QueryLookup ...
type QueryLookup url.Values

// Lookup ...
func (q QueryLookup) Lookup(key string) (string, bool) {
	vs, ok := q[key]
	if !ok || len(vs) < 1 {
		return "", false
	}
	return vs[0], true
}

// Lookuper ...
type Lookuper interface {
	Lookup(key string) (string, bool)
}
