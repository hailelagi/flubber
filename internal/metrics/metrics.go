package metrics

import (
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/hailelagi/flubber/internal/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
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

func StartMetricsServer() {
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Starting metrics server on :2112")
	log.Fatal(http.ListenAndServe(":2112", nil))
}

func StartMetricsPprof(config *config.Mount) {
	var profFile, memProfFile io.Writer
	var err error

	if config.Profile != "" {
		profFile, err = os.Create(config.Profile)
		if err != nil {
			zap.L().Warn("cannot create cpu profile", zap.Error(err))
		}
	}
	if config.MemProfile != "" {
		memProfFile, err = os.Create(config.MemProfile)
		if err != nil {
			zap.L().Warn("cannot create mem profile", zap.Error(err))
		}
	}

	runtime.GC()

	if profFile != nil {
		err := pprof.StartCPUProfile(profFile)

		if err != nil {
			zap.L().Info(err.Error())
		}
		defer pprof.StopCPUProfile()
	}

	if memProfFile != nil {
		err := pprof.WriteHeapProfile(memProfFile)
		if err != nil {
			zap.L().Info(err.Error())
		}
	}
}
