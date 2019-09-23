package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	counter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:      "analytics",
			Help:      "Data Endpoint was called",
			Subsystem: "gometheus",
		},
		[]string{"data"},
	)

	_ = prometheus.Register(counter)

	router.Handle("/data/{name}", dataHandler(counter)).Methods("GET")
	router.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":9090", router))
}

func dataHandler(counter *prometheus.CounterVec) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		counter.WithLabelValues(vars["name"]).Inc()
		w.Write([]byte(fmt.Sprintf("Hello %v",vars["name"])))
	})
}
