package customerror

import (
	"errors"
	"github.com/sirupsen/logrus"
	"url-shortener/pkg/config"
	"url-shortener/pkg/log"
	"url-shortener/pkg/metric"
)

type ErrorHandler struct {
	conf         *config.Config
	metricClient metric.Metric
}

func NewErrorHandler(conf *config.Config, metricClient metric.Metric) *ErrorHandler {
	return &ErrorHandler{
		conf:         conf,
		metricClient: metricClient,
	}
}

func (h *ErrorHandler) Handle(e error) {
	err := new(CustomError)

	if !errors.As(e, &err) {
		log.Log.Error(e.Error())
		return
	}

	log.Log.WithFields(
		logrus.Fields{
			"extras": err.Extras,
		},
	).Error(err.Error())

	h.metricClient.IncrementCounter(metric.ERROR, err.Domain, err.Type.String())
}
