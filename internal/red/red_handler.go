package red

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	labelHandler = ""
	labelMethod  = ""
	labelStatus  = ""
)

var duration = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "duration_seconds",
		Help:       "Summary of request duration in seconds",
		Objectives: map[float64]float64{0.9: 0.01, 0.95: 0.005, 0.99: 0.001},
	},
	[]string{labelHandler, labelMethod, labelStatus},
)
var errorsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "errors_total",
		Help: "Total number of errors",
	},
	[]string{labelHandler, labelMethod, labelStatus},
)
var requestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "request_total",
		Help: "Total number of requests",
	},
	[]string{labelHandler, labelMethod},
)

// MeasurableHandler ...
var MeasurableHandler = func(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		m := r.Method
		p := r.URL.Path

		requestsTotal.WithLabelValues(p, m).Inc()
		mw := newMeasurableWriter(w)
		h(mw, r)
		if mw.Status()/100 > 3 {
			errorsTotal.WithLabelValues(p, m, strconv.Itoa(mw.Status())).Inc()
		}
		duration.WithLabelValues(p, m, strconv.Itoa(mw.Status())).Observe(time.Since(t).Seconds())
	}
}
