package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"time"
)

type myMetrics struct {
	stateActive prometheus.GaugeVec
}

type cluster struct {
	name             string
	stateActive      prometheus.GaugeVec
	statePending     prometheus.GaugeVec
	stateRemoving    prometheus.GaugeVec
	stateUnavailable prometheus.GaugeVec
}

func new() myMetrics {
	mm := myMetrics{
		stateActive: *prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "stateActive",
				Help: "Shows if this cluster is active",
			},
			[]string{"name", "state"},
		),
	}

	//Register metrics with prometheus
	prometheus.MustRegister(mm.stateActive)

	// Return instance of struct
	return mm
}

func Collect() {
	//Create new metrics struct and seed with zero values
	mm := new()

	// Update values every 30s (simply increment them by 2)
	ticker := time.NewTicker(3 * time.Second)

	for range ticker.C {
		log.Info("Updating metric values")
		mm.stateActive.WithLabelValues("cluster1", "active").Set(1)
		mm.stateActive.WithLabelValues("cluster2", "active").Set(1)

	}

}
