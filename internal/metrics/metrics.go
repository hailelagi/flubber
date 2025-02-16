package metrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "fuse_operations_total",
			Help: "Total number of FUSE operations",
		},
		[]string{"operation", "status"},
	)

	opDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "fuse_operation_duration_seconds",
			Help:    "Latency of FUSE operations",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"operation"},
	)
)

func init() {
	prometheus.MustRegister(opCount, opDuration)
}

func startMetricsServer() {
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Starting metrics server on :2112")
	log.Fatal(http.ListenAndServe(":2112", nil))
}
