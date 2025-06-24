package main

import(
	//"fmt"
	"net/http"
	"strconv"
	"encoding/json"
)

type Response struct{
	Occurrences int `json:"occurrences"`
	Error string `json:"error,omitempty"`
}
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
	valStr = r.URL.Query().Get("seriesIncrement")
	seriesIncrement, err := strconv.ParseInt(valStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid seriesIncrement parameter",http.StatusBadRequest)
		return
	}
	valStr = r.URL.Query().Get("specifiedDigit")

	/*specifiedDigit*/
	val_temp1_64, err := strconv.ParseInt(valStr, 10, 64)//int not int64
	if err != nil {
		http.Error(w, "Invalid specifiedDigit parameter",http.StatusBadRequest)
		return
	}
	specifiedDigit := int(val_temp1_64)

	valStr = r.URL.Query().Get("seriesType")
	/*seriesType*/
	val_temp2_64, err := strconv.ParseInt(valStr, 10, 32)//int not 64
	if err != nil {
		http.Error(w, "Invalid seriesType parameter",http.StatusBadRequest)
		return
	}
	seriesType := int(val_temp2_64)

//	fmt.Fprintf(w, "Received params:\nseriesStart=%s\nseriesEnd=%s\nseriesIncrement=%s\nspecifiedDigit=%s\nseriesType=%s\n", seriesStart, seriesEnd, seriesIncrement, specifiedDigit, seriesType)


	occurrences, err := DigitOccurrence(seriesStart, seriesEnd, seriesIncrement, specifiedDigit, seriesType)
	resp := Response{}
	if err != nil {
		resp.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		resp.Occurrences = occurrences
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)



}
