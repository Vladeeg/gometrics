package handler

import (
	"net/http"
	"strconv"

	"github.com/Vladeeg/gometrics/internal/service/counter"
)

func HandleCount(res http.ResponseWriter, req *http.Request) {
	name := req.PathValue("name")
	value := req.PathValue("value")

	parsedValue, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)

		return
	}

	counter.UpdateCounter(name, parsedValue)

	res.WriteHeader(http.StatusOK)
}
