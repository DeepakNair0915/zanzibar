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

package workflow

import (
	"context"
	"net/textproto"

	"github.com/uber/zanzibar/config"

	zanzibar "github.com/uber/zanzibar/runtime"

	endpointsServerlessEndpointServerless "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/serverless-endpoint/serverless"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/serverless/module"
	"go.uber.org/zap"
)

// ServerLessServerless_responseDummyWorkflow defines the interface for ServerLessServerless_response workflow
type ServerLessServerless_responseDummyWorkflow interface {
	HandleDummy(
		ctx context.Context,
		reqHeaders zanzibar.Header,
		r *endpointsServerlessEndpointServerless.serverLess_Serverless_response_Args,
	) (*endpointsServerlessEndpointServerless.Response, zanzibar.Header, error)
}

// NewServerLessServerless_responseDummyWorkflow creates a workflow
func NewServerLessServerless_responseDummyWorkflow(deps *module.Dependencies) ServerLessServerless_responseDummyWorkflow {
	var whitelistedDynamicHeaders []string
	if deps.Default.Config.ContainsKey("clients..alternates") {
		var alternateServiceDetail config.AlternateServiceDetail
		deps.Default.Config.MustGetStruct("clients..alternates", &alternateServiceDetail)
		for _, routingConfig := range alternateServiceDetail.RoutingConfigs {
			whitelistedDynamicHeaders = append(whitelistedDynamicHeaders, textproto.CanonicalMIMEHeaderKey(routingConfig.HeaderName))
		}
	}

	return &serverLessServerlessResponseDummyWorkflow{
		Logger:                    deps.Default.Logger,
		whitelistedDynamicHeaders: whitelistedDynamicHeaders,
	}
}

// serverLessServerlessResponseDummyWorkflow calls thrift client .
type serverLessServerlessResponseDummyWorkflow struct {
	Logger                    *zap.Logger
	whitelistedDynamicHeaders []string
}

// HandleDummy calls thrift client.
func (w serverLessServerlessResponseDummyWorkflow) HandleDummy(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsServerlessEndpointServerless.ServerLess_ServerlessResponse_Args,
	r *endpointsServerlessEndpointServerless.serverLess_Serverless_response_Args,
) (*endpointsServerlessEndpointServerless.Response, zanzibar.Header, error) {

	response := convertServerless_responseDummyResponse(r)

	// Filter and map response headers from client to server response.
	resHeaders := zanzibar.ServerHTTPHeader{}

	return response, resHeaders, nil
}

func convertServerlessResponseDummyResponse(in *endpointsServerlessEndpointServerless.serverLess_Serverless_response_Args) *endpointsServerlessEndpointServerless.Response {
	out := &endpointsServerlessEndpointServerless.Response{}

	if in.Request != nil {
		out.FirstName = (*string)(in.Request.FirstName)
	}

	return out
}
