package handlers

import (
	"log"
	"net/http"
	"strconv"
)

// PerPage defines the default number of items per page
const PerPage = 15

func GetPageNum(r *http.Request) int {
	p := r.URL.Query().Get("page")

	if p == "" {
		return 1
	}
	pageNum, err := strconv.Atoi(p)
	if err != nil {
		log.Printf("%+v\n", err)

		pageNum = 1
	}
	return pageNum
}
