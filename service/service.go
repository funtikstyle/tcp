package service

import "fmt"

var Base map[string]string

func GET(k string) string {
	return Base[k]
}

func SET(k string, v string) {
	Base[k] = v
}

func KEYS(s string) string {
	for key, _ := range Base {
		return key
	}
}

func DEL(k string) bool {
	delete(Base, k)
}
