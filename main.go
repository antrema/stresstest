package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpSimultaneousRequests = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "simultaneous",
		Help: "Number of simultaneous requests",
	})
)

func simulateConnectionMetrics() {
	httpSimultaneousRequests.Set(500)
	go func() {
		for {
			httpSimultaneousRequests.Add(float64(rand.Intn(10) - 2))
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {

	prometheus.MustRegister(httpSimultaneousRequests)

	simulateConnectionMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
