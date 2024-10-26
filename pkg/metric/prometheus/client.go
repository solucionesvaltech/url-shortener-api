package prometheus

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
	"url-shortener/pkg/metric"

	"url-shortener/pkg/log"
)

var histogramMetrics = []Histogram{
	newHistogram(metric.DURATION, "Time elapsed in processing", []string{"use_case"}, []float64{0.5, 1, 3, 5, 10}),
}

var countersMetrics = []Counter{
	newCounter(metric.REQUEST, "Number of total requests received", []string{"use_case"}),
	newCounter(metric.OK, "Number of total events processed successfully by handler", []string{"use_case"}),
	newCounter(metric.ERROR, "Number of total events with custom_errors", []string{"use_case", "error_type"}),
}

type Histogram struct {
	Metric *prometheus.HistogramVec
	Name   string
}

type Counter struct {
	Metric *prometheus.CounterVec
	Name   string
}

type Client struct {
	counters   []Counter
	histograms []Histogram
}

func NewClient() *Client {
	return &Client{
		counters:   countersMetrics,
		histograms: histogramMetrics,
	}
}

func (mc *Client) Start() error {
	for _, c := range mc.counters {
		err := mc.addCounter(c.Metric)
		if err != nil {
			log.Log.Errorf("error initializing counter metrics %s", c.Name)
			return err
		}
		log.Log.Debugf("metric initialized %s", c.Name)
	}

	for _, h := range mc.histograms {
		err := mc.addHistogram(h.Metric)
		if err != nil {
			log.Log.Errorf("error initializing histogram metrics %s", h.Name)
			return err
		}
		log.Log.Debugf("metric initialized %s", h.Name)
	}

	return nil
}

func newHistogram(name string, description string, labels []string, buckets []float64) Histogram {
	histogram := Histogram{
		Metric: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    fmt.Sprintf("%s_%s", metric.APP, name),
				Help:    description,
				Buckets: buckets,
			},
			labels,
		),
		Name: name,
	}

	return histogram
}

func newCounter(name string, description string, labels []string) Counter {
	counter := Counter{
		Metric: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: fmt.Sprintf("%s_%s", metric.APP, name),
				Help: description,
			},
			labels,
		),
		Name: name,
	}

	return counter
}

func (mc *Client) addHistogram(metric *prometheus.HistogramVec) error {
	return prometheus.Register(metric)
}

func (mc *Client) addCounter(metric *prometheus.CounterVec) error {
	return prometheus.Register(metric)
}

func (mc *Client) getCounterByName(name string) *prometheus.CounterVec {
	log.Log.Debugf("getting counter by name %s", name)
	for _, c := range mc.counters {
		if c.Name == name {
			return c.Metric
		}
	}
	log.Log.Errorf("error getting counter by name %s", name)
	return nil
}

func (mc *Client) getHistogramByName(name string) *prometheus.HistogramVec {
	log.Log.Debugf("getting histogram by name %s", name)
	for _, h := range mc.histograms {
		if h.Name == name {
			return h.Metric
		}
	}
	log.Log.Errorf("error getting histogram by name %s", name)
	return nil
}

func (mc *Client) IncrementCounter(name string, labels ...string) {
	log.Log.WithFields(
		logrus.Fields{
			"name": name,
		},
	).Debug("incrementing counter metric")

	if counter := mc.getCounterByName(name); counter != nil {
		counter.WithLabelValues(labels...).Inc()
	}
}

func (mc *Client) ObserveHistogram(name string, start time.Time, labels ...string) {
	duration := time.Since(start).Seconds()

	log.Log.WithFields(
		logrus.Fields{
			"name":     name,
			"duration": strconv.FormatFloat(duration, 'f', -1, 64),
		},
	).Debug("observing histogram metric")

	if counter := mc.getHistogramByName(name); counter != nil {
		counter.WithLabelValues(labels...).Observe(duration)
	}
}
