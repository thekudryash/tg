// GENERATED BY i2s. DO NOT EDIT.
package transport

import (
	"time"

	kitPrometheus "github.com/go-kit/kit/metrics/prometheus"
	stdPrometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var srvMetrics *fasthttp.Server

var RequestCount = kitPrometheus.NewCounterFrom(stdPrometheus.CounterOpts{
	Help:      "Number of requests received",
	Name:      "count",
	Namespace: "service",
	Subsystem: "requests",
}, []string{"method", "service", "success"})

var RequestCountAll = kitPrometheus.NewCounterFrom(stdPrometheus.CounterOpts{
	Help:      "Number of all requests received",
	Name:      "all_count",
	Namespace: "service",
	Subsystem: "requests",
}, []string{"method", "service"})

var RequestLatency = kitPrometheus.NewSummaryFrom(stdPrometheus.SummaryOpts{
	Help:      "Total duration of requests in microseconds",
	Name:      "latency_microseconds",
	Namespace: "service",
	Subsystem: "requests",
}, []string{"method", "service", "success"})

func ServeMetrics(log logrus.FieldLogger, address string) {

	srvMetrics = &fasthttp.Server{
		Handler:     fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler()),
		ReadTimeout: time.Second * 10,
	}

	go func() {
		err := srvMetrics.ListenAndServe(address)
		ExitOnError(log, err, "serve metrics on "+address)
	}()
}
