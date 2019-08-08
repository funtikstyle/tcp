package service

type DBService struct {
	Base map[string]string
}

func NewDevice() *DBService {
	return &DBService{Base: make(map[string]string)}

}

func (db DBService) GET(k string) string {
	return db.Base[k]
}

func (db *DBService) SET(k string, v string) {
	db.Base[k] = v
}

func (db DBService) KEYS() string {
	result := ""
	for key, _ := range db.Base {
		result += key + " "
	}
	return result
}

func (db *DBService) DEL(k string) {
	delete(db.Base, k)
}
