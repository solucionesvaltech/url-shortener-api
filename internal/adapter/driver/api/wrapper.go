package api

import (
	"github.com/labstack/echo/v4"
	"time"
	"url-shortener/internal/adapter/driver/api/common"
	"url-shortener/pkg/customerror"

	"url-shortener/pkg/log"
	"url-shortener/pkg/metric"
)

type RequestWrapper struct {
	req          *RequestHandler
	metricClient metric.Metric
	errorHandler *customerror.ErrorHandler
}

func NewRequestWrapper(
	req *RequestHandler,
	metricClient metric.Metric,
	errorHandler *customerror.ErrorHandler,
) *RequestWrapper {
	return &RequestWrapper{
		req:          req,
		metricClient: metricClient,
		errorHandler: errorHandler,
	}
}

func (r *RequestWrapper) ApplyCommon(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		startTime := time.Now()
		log.Log.Debugf("apply common operation for: %s", r.req.Path)

		r.metricClient.IncrementCounter(metric.REQUEST, r.req.Handler.Domain().String())
		err := next(ctx)
		if err != nil {
			customError := common.GetErrorFromContext(ctx)
			r.errorHandler.Handle(customError)
			r.metricClient.ObserveHistogram(metric.DURATION, startTime, r.req.Handler.Domain().String())
			return err
		}

		r.metricClient.IncrementCounter(metric.OK, r.req.Handler.Domain().String())
		r.metricClient.ObserveHistogram(metric.DURATION, startTime, r.req.Handler.Domain().String())
		return err
	}
}
