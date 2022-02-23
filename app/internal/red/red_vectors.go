package red

import "github.com/prometheus/client_golang/prometheus"

// const (
// 	labelHandler = "my_handler"
// 	labelMethod  = "my_method"
// 	labelStatus  = "my_status"
// 	labelQuery   = "my_query"
// 	labelResult  = "my_result"
// 	labelService = "my_service"
// )

var (
	duration = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "backend2_duration_seconds",
			Help:       "Summary of request duration in seconds",
			Objectives: map[float64]float64{0.9: 0.01, 0.95: 0.005, 0.99: 0.001},
		},
		[]string{"URI", "METHOD"},
	)
	errorsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "backend2_errors_total",
			Help: "Total number of errors",
		},
		[]string{"URI", "METHOD", "CODE"},
	)
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "backend2_request_total",
			Help: "Total number of requests",
		},
		[]string{"URI", "METHOD"},
	)
)

var (
	DurationDBFunc = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "backend2_duration_db_method_seconds",
			Help:       "Summary of database method request duration in seconds",
			Objectives: map[float64]float64{0.9: 0.01, 0.95: 0.005, 0.99: 0.001},
		},
		[]string{"FUNC"},
	)

	DurationDBErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "backend2_duration_query_db_method_seconds",
			Help: "Summary of database Query method request duration in seconds",
		},
		[]string{"URI", "METHOD"},
	)

	DurationDBRequests = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "backend2_duration_query_row_db_method_seconds",
			Help:       "Summary of database QueryRow method request duration in seconds",
			Objectives: map[float64]float64{0.9: 0.01, 0.95: 0.005, 0.99: 0.001},
		},
		[]string{"URI", "METHOD"},
	)
)

func init() {
	prometheus.MustRegister(duration)
	prometheus.MustRegister(errorsTotal)
	prometheus.MustRegister(requestsTotal)

	prometheus.MustRegister(DurationDBFunc)
	prometheus.MustRegister(DurationDBErrors)
	prometheus.MustRegister(DurationDBRequests)
}

// todo: написать обертку для обращения к БД, втиснуть в нее метрики
