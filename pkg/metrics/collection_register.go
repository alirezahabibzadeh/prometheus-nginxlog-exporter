package metrics

import "github.com/prometheus/client_golang/prometheus"

func (c *Collection) MustRegister(r *prometheus.Registry) {
	r.MustRegister(c.RequestBytesTotal)
	r.MustRegister(c.ResponseSeconds)
	r.MustRegister(c.ResponseSecondsHist)
	r.MustRegister(c.ParseErrorsTotal)
	r.MustRegister(c.TotalRequestNumber)
	r.MustRegister(c.RequestMethodTotal)
	r.MustRegister(c.ResponseStatusCodeTotal)
	r.MustRegister(c.RemoteAddressTotal)
}
