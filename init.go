package analytics

import "github.com/prometheus/client_golang/prometheus"

func init() {
	prometheus.MustRegister(
		segmentBatchUploadCount,
		segmentBatchUploadSize,
		segmentBatchUploadTriggerCount,
		segmentBatchUploadMessagesGauge,
		segmentBatchConcurrentRequestsGauge,
		segmentUploadTime,
	)
}

var segmentBatchUploadCount = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "segment_batch_upload_count",
		Help: "segment_batch_upload_count",
	},
)

var segmentBatchUploadTriggerCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "segment_batch_upload_trigger_count",
		Help: "segment_batch_upload_trigger_count",
	},
	[]string{"trigger"},
)

var segmentBatchUploadMessagesGauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "segment_batch_upload_msgs_gauge",
		Help: "segment_batch_upload_msgs_gauge",
	},
)

var segmentBatchConcurrentRequestsGauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "segment_batch_concurrent_requests_gauge",
		Help: "segment_batch_concurrent_requests_gauge",
	},
)

var segmentBatchUploadSize = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "segment_batch_upload_size",
		Help: "segment_batch_upload_size",
	},
)

var segmentUploadTime = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:    "segment_batch_upload_time",
	Help:    "segment_batch_upload_time",
	Buckets: []float64{0.005, 0.01, 0.02, 0.04, 0.1, 0.5, 1.},
})
