package analytics

import "github.com/prometheus/client_golang/prometheus"

func init() {
	prometheus.MustRegister(
		segmentBatchUploadCount,
		segmentBatchUploadSize,
	)
}

var segmentBatchUploadCount = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "segment_batch_upload_count",
		Help: "segment_batch_upload_count",
	},
)
var segmentBatchUploadSize = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "segment_batch_upload_size",
		Help: "segment_batch_upload_size",
	},
)
