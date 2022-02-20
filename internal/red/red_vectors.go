package red

import "github.com/prometheus/client_golang/prometheus"

const (
	labelHandler = "my_handler"
	labelMethod  = "my_method"
	labelStatus  = "my_status"
	labelQuery   = "my_query"
	labelResult  = "my_result"
	labelService = "my_service"
)

var (
	duration = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "duration_seconds",
			Help:       "Summary of request duration in seconds",
			Objectives: map[float64]float64{0.9: 0.01, 0.95: 0.005, 0.99: 0.001},
		},
		[]string{labelHandler, labelMethod, labelStatus},
	)
	errorsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "errors_total",
			Help: "Total number of errors",
		},
		[]string{labelHandler, labelMethod, labelStatus},
	)
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "request_total",
			Help: "Total number of requests",
		},
		[]string{labelHandler, labelMethod},
	)
)

var (
	// Duration ...
	Duration = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "out_duration_seconds",
			Help:       "Summary of ongoing request duration in seconds",
			Objectives: map[float64]float64{0.9: 0.01, 0.95: 0.005, 0.99: 0.001},
		},
		[]string{labelService, labelQuery, labelResult},
	)

	// ErrorsTotal ...
	ErrorsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "out_errors_total",
			Help: "Total number of ongoing errors",
		},
		[]string{labelService, labelQuery, labelResult},
	)

	// RequestsTotal ...
	RequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "out_request_total",
			Help: "Total number of ongoing requests",
		},
		[]string{labelService, labelQuery},
	)
)

func init() {
	prometheus.MustRegister(duration)
	prometheus.MustRegister(errorsTotal)
	prometheus.MustRegister(requestsTotal)

	prometheus.MustRegister(Duration)
	prometheus.MustRegister(ErrorsTotal)
	prometheus.MustRegister(RequestsTotal)
}
