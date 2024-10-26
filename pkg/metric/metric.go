package metric

import (
	"time"
)

const (
	APP      = "shortener"
	REQUEST  = "request"
	RESPONSE = "response"
	OK       = "ok"
	ERROR    = "error"
	DURATION = "duration"
)

//go:generate mockgen -source=metric.go -destination=../../mocks/metricclient.go -package=mock

type Metric interface {
	IncrementCounter(metric string, labels ...string)
	ObserveHistogram(metric string, start time.Time, labels ...string)
	Start() error
}
