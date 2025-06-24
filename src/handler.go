package main

import(
	"fmt"
	"net/http"
)

func digitOccurrenceHandler(w http.ResponseWriter, r *http.Request) {

	valStr := r.URL.Query().Get("seriesStart")
	seriesStart, err := strconv.ParseInt(valStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid seriesStart parameter",http.StatusBadRequest)
		return
	}
	valStr = r.URL.Query().Get("seriesEnd")
	seriesEnd, err := strconv.ParseInt(valStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid seriesEnd parameter",http.StatusBadRequest)
		return
	}
	valEnd = r.URL.Query().Get("seriesIncrement")
	seriesIncrement, err := strconv.ParseInt(valStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid seriesIncrement parameter",http.StatusBadRequest)
		return
	}

	specifiedDigit := r.URL.Query().Get("specifiedDigit")
	seriesType := r.URL.Query().Get("seriesType")

	fmt.Fprintf(w, "Received params:\nseriesStart=%s\nseriesEnd=%s\nseriesIncrement=%s\nspecifiedDigit=%s\nseriesType=%s\n", seriesStart, seriesEnd, seriesIncrement, specifiedDigit, seriesType)

}
