// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"

	"github.com/seniorGolang/tg/example/interfaces"
)

type httpJsonRPC struct {
	log          logrus.FieldLogger
	errorHandler ErrorHandler
	svc          *serverJsonRPC
}

func NewJsonRPC(log logrus.FieldLogger, svcJsonRPC interfaces.JsonRPC) (srv *httpJsonRPC) {

	srv = &httpJsonRPC{
		log: log,
		svc: newServerJsonRPC(svcJsonRPC),
	}
	return
}

func (http httpJsonRPC) Service() MiddlewareSetJsonRPC {
	return http.svc
}

func (http *httpJsonRPC) WithLog(log logrus.FieldLogger) *httpJsonRPC {
	http.svc.WithLog(log)
	return http
}

func (http *httpJsonRPC) WithTrace() *httpJsonRPC {
	http.svc.WithTrace()
	return http
}

func (http *httpJsonRPC) WithErrorHandler(handler ErrorHandler) *httpJsonRPC {
	http.errorHandler = handler
	return http
}

func (http *httpJsonRPC) SetRoutes(route *router.Router) {

	route.POST("/jsonrpc", http.serveBatch)
	route.POST("/jsonRPC/test", http.serveTest)
}
