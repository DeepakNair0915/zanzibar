// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package googlenowendpoint

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/googlenow/module"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	workflow "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/googlenow/workflow"
	endpointsGooglenowGooglenow "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/googlenow/googlenow"
)

// GoogleNowAddCredentialsHandler is the handler for "/googlenow/add-credentials"
type GoogleNowAddCredentialsHandler struct {
	Dependencies *module.Dependencies
	endpoint     *zanzibar.RouterEndpoint
}

// NewGoogleNowAddCredentialsHandler creates a handler
func NewGoogleNowAddCredentialsHandler(deps *module.Dependencies) *GoogleNowAddCredentialsHandler {
	handler := &GoogleNowAddCredentialsHandler{
		Dependencies: deps,
	}
	handler.endpoint = zanzibar.NewRouterEndpointContext(
		deps.Default.ContextExtractor, deps.Default.Logger, deps.Default.Scope, deps.Default.Tracer,
		"googlenow", "addCredentials",
		handler.HandleRequest,
	)

	return handler
}

// Register adds the http handler to the gateway's http router
func (h *GoogleNowAddCredentialsHandler) Register(g *zanzibar.Gateway) error {
	return g.HTTPRouter.Register(
		"POST", "/googlenow/add-credentials",
		h.endpoint,
	)
}

// HandleRequest handles "/googlenow/add-credentials".
func (h *GoogleNowAddCredentialsHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	defer func() {
		if r := recover(); r != nil {
			stacktrace := string(debug.Stack())
			e := errors.Errorf("enpoint panic: %v, stacktrace: %v", r, stacktrace)
			h.Dependencies.Default.ContextLogger.Error(
				ctx,
				"endpoint panic",
				zap.Error(e),
				zap.String("stacktrace", stacktrace),
				zap.String("endpoint", h.endpoint.EndpointName))

			h.endpoint.ContextMetrics.EndpointMetrics.Panic.Inc(1)
			res.SendError(502, "Unexpected workflow panic, recovered at endpoint.", nil)
		}
	}()

	if !req.CheckHeaders([]string{"X-Token", "X-Uuid"}) {
		return
	}
	var requestBody endpointsGooglenowGooglenow.GoogleNow_AddCredentials_Args
	if ok := req.ReadAndUnmarshalBody(&requestBody); !ok {
		return
	}

	// log endpoint request to downstream services
	if ce := h.Dependencies.Default.ContextLogger.Check(zapcore.DebugLevel, "stub"); ce != nil {
		zfields := []zapcore.Field{
			zap.String("endpoint", h.endpoint.EndpointName),
		}
		zfields = append(zfields, zap.String("body", fmt.Sprintf("%s", req.GetRawBody())))
		for _, k := range req.Header.Keys() {
			if val, ok := req.Header.Get(k); ok {
				zfields = append(zfields, zap.String(k, val))
			}
		}
		h.Dependencies.Default.ContextLogger.Debug(ctx, "endpoint request to downstream", zfields...)
	}

	w := workflow.NewGoogleNowAddCredentialsWorkflow(h.Dependencies)
	if span := req.GetSpan(); span != nil {
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	cliRespHeaders, err := w.Handle(ctx, req.Header, &requestBody)
	if err != nil {
		res.SendError(500, "Unexpected server error", err)
		return

	}

	res.WriteJSONBytes(202, cliRespHeaders, nil)
}
