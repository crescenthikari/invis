package handler

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/gocarina/gocsv"
)

func ResponseOk(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"status": 0,
		"result": data,
	}
	js, err := json.Marshal(response)
	if err != nil {
		resp := map[string]interface{}{
			"status": 1,
			"error":  fmt.Sprintf("%v", err),
		}
		js, _ = json.Marshal(resp)
	}
	w.WriteHeader(200)
	w.Write(js)
}

func WriteCsv(w http.ResponseWriter, filename string, data interface{}) {
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename="+filename+".csv")

	gocsv.Marshal(data, w)
}

func ResponseErr(w http.ResponseWriter, err error) {
	resp := map[string]interface{}{
		"status": 1,
		"error":  fmt.Sprintf("%v", err),
	}
	js, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	w.Write(js)
}
