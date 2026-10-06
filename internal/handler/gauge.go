package handler

import (
	"net/http"
	"strconv"

	"github.com/Vladeeg/gometrics/internal/service/gauge"
)

func HandleGauge(res http.ResponseWriter, req *http.Request) {
	name := req.PathValue("name")
	value := req.PathValue("value")

	parsedValue, err := strconv.ParseFloat(value, 64)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)

		return
	}

	gauge.SetGauge(name, parsedValue)

	res.WriteHeader(http.StatusOK)
}
