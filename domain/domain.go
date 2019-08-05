package domain

type DBService interface {
	GET(k string) string
	SET(k string, v string)
	KEYS() []string
	DEL(k string)
}
