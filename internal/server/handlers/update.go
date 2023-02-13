package handlers

import (
	"fmt"
	"github.com/lastbyte32/go-metric/internal/server/storage"
	"net/http"
	"strconv"
	"strings"
)

func UpdateHandle(gaugeStorage storage.GaugeStorage, counterStorage storage.CounterStorage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("UpdateHandle")

		urlPath := strings.Split(strings.TrimLeft(r.URL.Path, "update/"), "/")
		if len(urlPath) != 3 {
			fmt.Printf("parse err\n Count: %d\nParam: %v", len(urlPath), urlPath)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		switch urlPath[0] {
		case "gauge":
			fmt.Println("case gauge")

			valueGauge, err := strconv.ParseFloat(urlPath[2], 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			gaugeStorage.Update(urlPath[1], valueGauge)

		case "counter":
			fmt.Println("case counter")

			valueCounter, err := strconv.ParseInt(urlPath[2], 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			counterStorage.Update(urlPath[1], valueCounter)

		default:
			fmt.Println("NotImplemented")
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
		_, err := w.Write([]byte(fmt.Sprintf("Type: %s\nName: %s\nValue: %s", urlPath[0], urlPath[1], urlPath[2])))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		//w.Header().Set("Content-Type", "text/plain")

	}
}
