package analytics

import "github.com/prometheus/client_golang/prometheus"

func init() {
	prometheus.MustRegister(
		segmentBatchUploadCount,
		segmentBatchUploadSize,
		segmentBatchUploadTriggerCount,
		segmentBatchUploadMessagesGauge,
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

var segmentBatchUploadSize = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "segment_batch_upload_size",
		Help: "segment_batch_upload_size",
	},
)
