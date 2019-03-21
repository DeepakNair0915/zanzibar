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

package barendpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/thriftrw/ptr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	workflow "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/workflow"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"

	defaultExample "github.com/uber/zanzibar/examples/example-gateway/middlewares/default/default_example"
	defaultExample2 "github.com/uber/zanzibar/examples/example-gateway/middlewares/default/default_example2"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/module"
)

// BarArgWithNestedQueryParamsHandler is the handler for "/bar/argWithNestedQueryParams"
type BarArgWithNestedQueryParamsHandler struct {
	Dependencies *module.Dependencies
	endpoint     *zanzibar.RouterEndpoint
}

// NewBarArgWithNestedQueryParamsHandler creates a handler
func NewBarArgWithNestedQueryParamsHandler(deps *module.Dependencies) *BarArgWithNestedQueryParamsHandler {
	handler := &BarArgWithNestedQueryParamsHandler{
		Dependencies: deps,
	}
	handler.endpoint = zanzibar.NewRouterEndpoint(
		deps.Default.ContextExtractor, deps.Default,
		"bar", "argWithNestedQueryParams",
		zanzibar.NewStack([]zanzibar.MiddlewareHandle{
			deps.Middleware.DefaultExample2.NewMiddlewareHandle(
				defaultExample2.Options{},
			),
			deps.Middleware.DefaultExample.NewMiddlewareHandle(
				defaultExample.Options{},
			),
		}, handler.HandleRequest).Handle,
	)

	return handler
}

// Register adds the http handler to the gateway's http router
func (h *BarArgWithNestedQueryParamsHandler) Register(g *zanzibar.Gateway) error {
	return g.HTTPRouter.Handle(
		"GET", "/bar/argWithNestedQueryParams",
		http.HandlerFunc(h.endpoint.HandleRequest),
	)
}

// HandleRequest handles "/bar/argWithNestedQueryParams".
func (h *BarArgWithNestedQueryParamsHandler) HandleRequest(
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
				"Endpoint failure: endpoint panic",
				zap.Error(e),
				zap.String("stacktrace", stacktrace),
				zap.String("endpoint", h.endpoint.EndpointName))

			h.Dependencies.Default.ContextMetrics.IncCounter(ctx, zanzibar.MetricEndpointPanics, 1)
			res.SendError(502, "Unexpected workflow panic, recovered at endpoint.", nil)
		}
	}()

	var requestBody endpointsBarBar.Bar_ArgWithNestedQueryParams_Args

	if requestBody.Request == nil {
		requestBody.Request = &endpointsBarBar.QueryParamsStruct{}
	}
	xUUIDValue, xUUIDValueExists := req.Header.Get("x-uuid")
	if xUUIDValueExists {
		requestBody.Request.AuthUUID = ptr.String(xUUIDValue)
	}
	xUUID2Value, xUUID2ValueExists := req.Header.Get("x-uuid2")
	if xUUID2ValueExists {
		requestBody.Request.AuthUUID2 = ptr.String(xUUID2Value)
	}

	if requestBody.Request == nil {
		requestBody.Request = &endpointsBarBar.QueryParamsStruct{}
	}
	requestNameOk := req.CheckQueryValue("request.name")
	if !requestNameOk {
		return
	}
	requestNameQuery, ok := req.GetQueryValue("request.name")
	if !ok {
		return
	}
	requestBody.Request.Name = requestNameQuery

	requestUserUUIDOk := req.HasQueryValue("request.userUUID")
	if requestUserUUIDOk {
		requestUserUUIDQuery, ok := req.GetQueryValue("request.userUUID")
		if !ok {
			return
		}
		requestBody.Request.UserUUID = ptr.String(requestUserUUIDQuery)
	}

	requestFooOk := req.CheckQueryValue("request.foo")
	if !requestFooOk {
		return
	}
	requestFooQuery, ok := req.GetQueryValues("request.foo")
	if !ok {
		return
	}
	requestBody.Request.Foo = requestFooQuery

	if req.HasQueryPrefix("opt") || requestBody.Opt != nil {
		if requestBody.Opt == nil {
			requestBody.Opt = &endpointsBarBar.QueryParamsOptsStruct{}
		}
		optNameOk := req.CheckQueryValue("opt.name")
		if !optNameOk {
			return
		}
		optNameQuery, ok := req.GetQueryValue("opt.name")
		if !ok {
			return
		}
		requestBody.Opt.Name = optNameQuery

		optUserUUIDOk := req.HasQueryValue("opt.userUUID")
		if optUserUUIDOk {
			optUserUUIDQuery, ok := req.GetQueryValue("opt.userUUID")
			if !ok {
				return
			}
			requestBody.Opt.UserUUID = ptr.String(optUserUUIDQuery)
		}

		optAuthUUIDOk := req.HasQueryValue("opt.authUUID")
		if optAuthUUIDOk {
			optAuthUUIDQuery, ok := req.GetQueryValue("opt.authUUID")
			if !ok {
				return
			}
			requestBody.Opt.AuthUUID = ptr.String(optAuthUUIDQuery)
		}

		optAuthUUID2Ok := req.HasQueryValue("opt.authUUID2")
		if optAuthUUID2Ok {
			optAuthUUID2Query, ok := req.GetQueryValue("opt.authUUID2")
			if !ok {
				return
			}
			requestBody.Opt.AuthUUID2 = ptr.String(optAuthUUID2Query)
		}

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

	w := workflow.NewBarArgWithNestedQueryParamsWorkflow(h.Dependencies)
	if span := req.GetSpan(); span != nil {
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	response, cliRespHeaders, err := w.Handle(ctx, req.Header, &requestBody)

	// log downstream response to endpoint
	if ce := h.Dependencies.Default.ContextLogger.Check(zapcore.DebugLevel, "stub"); ce != nil {
		zfields := []zapcore.Field{
			zap.String("endpoint", h.endpoint.EndpointName),
		}
		if body, err := json.Marshal(response); err == nil {
			zfields = append(zfields, zap.String("body", fmt.Sprintf("%s", body)))
		}
		for _, k := range cliRespHeaders.Keys() {
			if val, ok := cliRespHeaders.Get(k); ok {
				zfields = append(zfields, zap.String(k, val))
			}
		}
		if traceKey, ok := req.Header.Get("x-trace-id"); ok {
			zfields = append(zfields, zap.String("x-trace-id", traceKey))
		}
		h.Dependencies.Default.ContextLogger.Debug(ctx, "downstream service response", zfields...)
	}

	if err != nil {
		res.SendError(500, "Unexpected server error", err)
		return

	}

	res.WriteJSON(200, cliRespHeaders, response)
}
