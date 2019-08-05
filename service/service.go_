package service

import "fmt"

type dbservice struct {
	Base map[string]string
}

//var Base map[string]string

func (db dbservice) GET(k string) string {
	return db.Base[k]
}

func (db dbservice) SET(k string, v string) {
	db.Base[k] = v
}

func (db dbservice) KEYS(s string) string {
	for key, _ := range db.Base {
		return key
	}
}

func (db dbservice) DEL(k string) bool {
	delete(db.Base, k)
}
