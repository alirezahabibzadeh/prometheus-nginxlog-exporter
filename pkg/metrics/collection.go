package metrics

import "github.com/prometheus/client_golang/prometheus"

// Collection is a struct containing pointers to all metrics that should be
// exposed to Prometheus
type Collection struct {
	RequestBytesTotal          *prometheus.CounterVec
	ResponseSeconds            *prometheus.SummaryVec
	ResponseSecondsHist        *prometheus.HistogramVec
	ParseErrorsTotal            prometheus.Counter
	TotalRequestNumber          prometheus.Counter
	RequestMethodTotal		   *prometheus.CounterVec
	ResponseStatusCodeTotal	   *prometheus.CounterVec
	RemoteAddressTotal		   *prometheus.CounterVec
}