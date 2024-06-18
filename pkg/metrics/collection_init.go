package metrics

import (
	"github.com/martin-helmich/prometheus-nginxlog-exporter/pkg/config"
	"github.com/martin-helmich/prometheus-nginxlog-exporter/pkg/relabeling"
	"github.com/prometheus/client_golang/prometheus"
)

// Init initializes a metrics struct
func (m *Collection) Init(cfg *config.NamespaceConfig) {
	cfg.MustCompile()

	labels := cfg.OrderedLabelNames
	counterLabels := labels

	relabelings := relabeling.NewRelabelings(cfg.RelabelConfigs)
	relabelings = append(relabelings, relabeling.DefaultRelabelings...)
	relabelings = relabeling.UniqueRelabelings(relabelings)

	for _, r := range relabelings {
		if !r.OnlyCounter {
			labels = append(labels, r.TargetLabel)
		}
		counterLabels = append(counterLabels, r.TargetLabel)
	}

	m.RequestBytesTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   cfg.NamespacePrefix,
		ConstLabels: cfg.NamespaceLabels,
		Name:        "http_request_size_bytes",
		Help:        "Total amount of received bytes",
	}, labels)

	m.ResponseSeconds = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace:   cfg.NamespacePrefix,
		ConstLabels: cfg.NamespaceLabels,
		Name:        "http_response_time_seconds",
		Help:        "Time needed by NGINX to handle requests",
		Objectives:  map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, labels)

	m.ResponseSecondsHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   cfg.NamespacePrefix,
		ConstLabels: cfg.NamespaceLabels,
		Name:        "http_response_time_seconds_hist",
		Help:        "Time needed by NGINX to handle requests",
		Buckets:     cfg.HistogramBuckets,
	}, labels)

	m.ParseErrorsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace:   cfg.NamespacePrefix,
		ConstLabels: cfg.NamespaceLabels,
		Name:        "parse_errors_total",
		Help:        "Total number of log file lines that could not be parsed",
	})
	m.TotalRequestNumber = prometheus.NewCounter(prometheus.CounterOpts{ // Initialize new field
		Namespace:   cfg.NamespacePrefix,
		ConstLabels: cfg.NamespaceLabels,
		Name:        "total_request_number",
		Help:        "Total number of requests",
	})
	m.RequestMethodTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   cfg.NamespacePrefix,
		ConstLabels: cfg.NamespaceLabels,
		Name:        "request_method_total",
		Help:        "Total number of requests by method",
	}, labels)
	m.ResponseStatusCodeTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   cfg.NamespacePrefix,
		ConstLabels: cfg.NamespaceLabels,
		Name:        "response_status_code_total",
		Help:        "Total number of responses by status code",
	}, labels)
	m.RemoteAddressTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   cfg.NamespacePrefix,
		ConstLabels: cfg.NamespaceLabels,
		Name:        "remote_address_total",
		Help:        "Total number of requests by remote address",
	}, labels)
}
