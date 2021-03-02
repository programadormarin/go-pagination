package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"worten/paginate"
)

type RequestHandler struct {
}

func (rh *RequestHandler) validateRequest(r *http.Request) (map[string]int, error) {
	current_page, ok := r.URL.Query()["current_page"]
	total_pages, ok := r.URL.Query()["total_pages"]
	boundaries, ok := r.URL.Query()["boundaries"]
	around, ok := r.URL.Query()["around"]

	res := make(map[string]int)

	if !ok || len(current_page[0]) < 1 {
		return res, errors.New("Url Param 'current_page' is missing")
	}

	if !ok || len(total_pages[0]) < 1 {
		return res, errors.New("Url Param 'total_pages' is missing")
	}

	if !ok || len(boundaries[0]) < 1 {
		return res, errors.New("Url Param 'boundaries' is missing")
	}

	if !ok || len(around[0]) < 1 {
		return res, errors.New("Url Param 'around' is missing")
	}

	var e error

	res["cp"], e = strconv.Atoi(current_page[0])

	if e != nil {
		return res, errors.New("Url Param 'current_page' is invalid")
	}

	res["tp"], e = strconv.Atoi(total_pages[0])
	if e != nil {
		return res, errors.New("Url Param 'total_pages' is invalid")
	}

	res["b"], e = strconv.Atoi(boundaries[0])
	if e != nil {
		return res, errors.New("Url Param 'boundaries' is invalid")
	}

	res["a"], e = strconv.Atoi(around[0])
	if e != nil {
		return res, errors.New("Url Param 'around' is invalid")
	}

	return res, nil
}

func (rh *RequestHandler) getPagination(w http.ResponseWriter, r *http.Request) {

	params, err := rh.validateRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	p := paginate.Pagination{
		CurrentPage: params["cp"],
		TotalPages:  params["tp"],
		Boundaries:  params["b"],
		Around:      params["a"],
	}

	fmt.Fprintf(w, p.GetPages())
	return
}

func (rh *RequestHandler) Handle() {
	http.HandleFunc("/", rh.getPagination)

	fmt.Println(http.ListenAndServe(":80", nil))
}
