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

package googlenowclient

import (
	"context"
	"fmt"
	"time"

	"github.com/afex/hystrix-go/hystrix"

	zanzibar "github.com/uber/zanzibar/runtime"

	module "github.com/uber/zanzibar/examples/example-gateway/build/clients/google-now/module"
	clientsGooglenowGooglenow "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/googlenow/googlenow"
)

// Client defines google-now client interface.
type Client interface {
	HTTPClient() *zanzibar.HTTPClient
	AddCredentials(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsGooglenowGooglenow.GoogleNowService_AddCredentials_Args,
	) (map[string]string, error)
	CheckCredentials(
		ctx context.Context,
		reqHeaders map[string]string,
	) (map[string]string, error)
}

// googleNowClient is the http client.
type googleNowClient struct {
	clientID               string
	httpClient             *zanzibar.HTTPClient
	circuitBreakerDisabled bool
	requestUUIDHeaderKey   string
}

// NewClient returns a new http client.
func NewClient(deps *module.Dependencies) Client {
	ip := deps.Default.Config.MustGetString("clients.google-now.ip")
	port := deps.Default.Config.MustGetInt("clients.google-now.port")
	baseURL := fmt.Sprintf("http://%s:%d", ip, port)
	timeoutVal := int(deps.Default.Config.MustGetInt("clients.google-now.timeout"))
	timeout := time.Millisecond * time.Duration(
		timeoutVal,
	)
	defaultHeaders := make(map[string]string)
	if deps.Default.Config.ContainsKey("http.defaultHeaders") {
		deps.Default.Config.MustGetStruct("http.defaultHeaders", &defaultHeaders)
	}
	if deps.Default.Config.ContainsKey("clients.google-now.defaultHeaders") {
		deps.Default.Config.MustGetStruct("clients.google-now.defaultHeaders", &defaultHeaders)
	}
	var requestUUIDHeaderKey string
	if deps.Default.Config.ContainsKey("http.clients.requestUUIDHeaderKey") {
		requestUUIDHeaderKey = deps.Default.Config.MustGetString("http.clients.requestUUIDHeaderKey")
	}

	circuitBreakerDisabled := configureCicruitBreaker(deps, timeoutVal)

	return &googleNowClient{
		clientID: "google-now",
		httpClient: zanzibar.NewHTTPClientContext(
			deps.Default.Logger, deps.Default.ContextMetrics,
			"google-now",
			map[string]string{
				"AddCredentials":   "GoogleNowService::addCredentials",
				"CheckCredentials": "GoogleNowService::checkCredentials",
			},
			baseURL,
			defaultHeaders,
			timeout,
		),
		circuitBreakerDisabled: circuitBreakerDisabled,
		requestUUIDHeaderKey:   requestUUIDHeaderKey,
	}
}

func configureCicruitBreaker(deps *module.Dependencies, timeoutVal int) bool {
	// circuitBreakerDisabled sets whether circuit-breaker should be disabled
	circuitBreakerDisabled := false
	if deps.Default.Config.ContainsKey("clients.google-now.circuitBreakerDisabled") {
		circuitBreakerDisabled = deps.Default.Config.MustGetBoolean("clients.google-now.circuitBreakerDisabled")
	}
	// sleepWindowInMilliseconds sets the amount of time, after tripping the circuit,
	// to reject requests before allowing attempts again to determine if the circuit should again be closed
	sleepWindowInMilliseconds := 5000
	if deps.Default.Config.ContainsKey("clients.google-now.sleepWindowInMilliseconds") {
		sleepWindowInMilliseconds = int(deps.Default.Config.MustGetInt("clients.google-now.sleepWindowInMilliseconds"))
	}
	// maxConcurrentRequests sets how many requests can be run at the same time, beyond which requests are rejected
	maxConcurrentRequests := 20
	if deps.Default.Config.ContainsKey("clients.google-now.maxConcurrentRequests") {
		maxConcurrentRequests = int(deps.Default.Config.MustGetInt("clients.google-now.maxConcurrentRequests"))
	}
	// errorPercentThreshold sets the error percentage at or above which the circuit should trip open
	errorPercentThreshold := 20
	if deps.Default.Config.ContainsKey("clients.google-now.errorPercentThreshold") {
		errorPercentThreshold = int(deps.Default.Config.MustGetInt("clients.google-now.errorPercentThreshold"))
	}
	// requestVolumeThreshold sets a minimum number of requests that will trip the circuit in a rolling window of 10s
	// For example, if the value is 20, then if only 19 requests are received in the rolling window of 10 seconds
	// the circuit will not trip open even if all 19 failed.
	requestVolumeThreshold := 20
	if deps.Default.Config.ContainsKey("clients.google-now.requestVolumeThreshold") {
		requestVolumeThreshold = int(deps.Default.Config.MustGetInt("clients.google-now.requestVolumeThreshold"))
	}
	if !circuitBreakerDisabled {
		hystrix.ConfigureCommand("google-now", hystrix.CommandConfig{
			MaxConcurrentRequests:  maxConcurrentRequests,
			ErrorPercentThreshold:  errorPercentThreshold,
			SleepWindow:            sleepWindowInMilliseconds,
			RequestVolumeThreshold: requestVolumeThreshold,
			Timeout:                timeoutVal,
		})
	}
	return circuitBreakerDisabled
}

// HTTPClient returns the underlying HTTP client, should only be
// used for internal testing.
func (c *googleNowClient) HTTPClient() *zanzibar.HTTPClient {
	return c.httpClient
}

// AddCredentials calls "/add-credentials" endpoint.
func (c *googleNowClient) AddCredentials(
	ctx context.Context,
	headers map[string]string,
	r *clientsGooglenowGooglenow.GoogleNowService_AddCredentials_Args,
) (map[string]string, error) {
	reqUUID := zanzibar.RequestUUIDFromCtx(ctx)
	if reqUUID != "" {
		if headers == nil {
			headers = make(map[string]string)
		}
		headers[c.requestUUIDHeaderKey] = reqUUID
	}

	req := zanzibar.NewClientHTTPRequest(ctx, c.clientID, "AddCredentials", "GoogleNowService::addCredentials", c.httpClient)

	// Generate full URL.
	fullURL := c.httpClient.BaseURL + "/add-credentials"

	err := req.WriteJSON("POST", fullURL, headers, r)
	if err != nil {
		return nil, err
	}

	headerErr := req.CheckHeaders([]string{"x-uuid"})
	if headerErr != nil {
		return nil, headerErr
	}

	var res *zanzibar.ClientHTTPResponse
	if c.circuitBreakerDisabled {
		res, err = req.Do()
	} else {
		err = hystrix.DoC(ctx, "google-now", func(ctx context.Context) error {
			res, err = req.Do()
			if res.StatusCode < 500 {
				return nil
			}
			return err
		}, nil)
	}
	if err != nil {
		return nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}
	// TODO(jakev): verify mandatory response headers

	res.CheckOKResponse([]int{202})

	switch res.StatusCode {
	case 202:
		_, err = res.ReadAll()
		if err != nil {
			return respHeaders, err
		}
		return respHeaders, nil
	default:
		_, err = res.ReadAll()
		if err != nil {
			return respHeaders, err
		}
	}

	return respHeaders, &zanzibar.UnexpectedHTTPError{
		StatusCode: res.StatusCode,
		RawBody:    res.GetRawBody(),
	}
}

// CheckCredentials calls "/check-credentials" endpoint.
func (c *googleNowClient) CheckCredentials(
	ctx context.Context,
	headers map[string]string,
) (map[string]string, error) {
	reqUUID := zanzibar.RequestUUIDFromCtx(ctx)
	if reqUUID != "" {
		if headers == nil {
			headers = make(map[string]string)
		}
		headers[c.requestUUIDHeaderKey] = reqUUID
	}

	req := zanzibar.NewClientHTTPRequest(ctx, c.clientID, "CheckCredentials", "GoogleNowService::checkCredentials", c.httpClient)

	// Generate full URL.
	fullURL := c.httpClient.BaseURL + "/check-credentials"

	err := req.WriteJSON("POST", fullURL, headers, nil)
	if err != nil {
		return nil, err
	}

	headerErr := req.CheckHeaders([]string{"x-uuid"})
	if headerErr != nil {
		return nil, headerErr
	}

	var res *zanzibar.ClientHTTPResponse
	if c.circuitBreakerDisabled {
		res, err = req.Do()
	} else {
		err = hystrix.DoC(ctx, "google-now", func(ctx context.Context) error {
			res, err = req.Do()
			if res.StatusCode < 500 {
				return nil
			}
			return err
		}, nil)
	}
	if err != nil {
		return nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}
	// TODO(jakev): verify mandatory response headers

	res.CheckOKResponse([]int{202})

	switch res.StatusCode {
	case 202:
		_, err = res.ReadAll()
		if err != nil {
			return respHeaders, err
		}
		return respHeaders, nil
	default:
		_, err = res.ReadAll()
		if err != nil {
			return respHeaders, err
		}
	}

	return respHeaders, &zanzibar.UnexpectedHTTPError{
		StatusCode: res.StatusCode,
		RawBody:    res.GetRawBody(),
	}
}
