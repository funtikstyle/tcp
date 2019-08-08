package handler

import (
	"fmt"
	"strings"
	"tcp/domain"
)

type DBHandler struct {
	srv domain.DBService
}

func NewDBHandler(srv domain.DBService) *DBHandler {
	return &DBHandler{srv: srv}
}

func (h DBHandler) Req(data string) string {
	res := ""
	sep := strings.Split(string(data), " ")
	fmt.Printf("read: %v \n", string(data))

	switch sep[0] {
	case "GET":
		res = h.srv.GET(sep[1])
	case "SET":
		h.srv.SET(sep[1], sep[2])
		res = "запись добавлена"
		//res = "запись изменена"
	case "KEYS":
		res = h.srv.KEYS()
	case "DEL":
		h.srv.DEL(sep[1])
		res = "запись удалена"
		//res = "запись не найдена"
	default:
		res = "команда не найдена"
	}
	return res + "\n"
}
