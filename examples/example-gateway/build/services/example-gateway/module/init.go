// Code generated by zanzibar
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
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

package module

import (
	barClientGenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/bar"
	bazClientGenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/baz"
	contactsClientGenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/contacts"
	googlenowClientGenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/google-now"
	barEndpointGenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar"

	barEndpointModule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/module"
	bazEndpointGenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz"

	bazEndpointModule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz/module"
	contactsEndpointGenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/contacts"

	contactsEndpointModule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/contacts/module"
	googlenowEndpointGenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/googlenow"

	googlenowEndpointModule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/googlenow/module"
	"github.com/uber/zanzibar/runtime"
)

// clientDependencies contains client dependencies
type clientDependencies struct {
	Bar       barClientGenerated.Client
	Baz       bazClientGenerated.Client
	Contacts  contactsClientGenerated.Client
	GoogleNow googlenowClientGenerated.Client
}

// endpointDependencies contains endpoint dependencies
type endpointDependencies struct {
	Bar       barEndpointGenerated.Endpoint
	Baz       bazEndpointGenerated.Endpoint
	Contacts  contactsEndpointGenerated.Endpoint
	Googlenow googlenowEndpointGenerated.Endpoint
}

func InitializeService(gateway *zanzibar.Gateway) Service {
	initializedClientDependencies := &clientDependencies{}
	initializedClientDependencies.Bar = barClientGenerated.NewClient(gateway)
	initializedClientDependencies.Baz = bazClientGenerated.NewClient(gateway)
	initializedClientDependencies.Contacts = contactsClientGenerated.NewClient(gateway)
	initializedClientDependencies.GoogleNow = googlenowClientGenerated.NewClient(gateway)

	initializedEndpointDependencies := &endpointDependencies{}
	initializedEndpointDependencies.Bar = barEndpointGenerated.NewEndpoint(gateway, &barEndpointGenerated.Dependencies{
		Client: &barEndpointModule.ClientDependencies{
			Bar: initializedClientDependencies.Bar,
		},
	})
	initializedEndpointDependencies.Baz = bazEndpointGenerated.NewEndpoint(gateway, &bazEndpointGenerated.Dependencies{
		Client: &bazEndpointModule.ClientDependencies{
			Baz: initializedClientDependencies.Baz,
		},
	})
	initializedEndpointDependencies.Contacts = contactsEndpointGenerated.NewEndpoint(gateway, &contactsEndpointGenerated.Dependencies{
		Client: &contactsEndpointModule.ClientDependencies{
			Contacts: initializedClientDependencies.Contacts,
		},
	})
	initializedEndpointDependencies.Googlenow = googlenowEndpointGenerated.NewEndpoint(gateway, &googlenowEndpointGenerated.Dependencies{
		Client: &googlenowEndpointModule.ClientDependencies{
			GoogleNow: initializedClientDependencies.GoogleNow,
		},
	})

	return NewService(gateway, &Dependencies{
		Endpoint: &EndpointDependencies{
			Bar:       initializedEndpointDependencies.Bar,
			Baz:       initializedEndpointDependencies.Baz,
			Contacts:  initializedEndpointDependencies.Contacts,
			Googlenow: initializedEndpointDependencies.Googlenow,
		},
	})
}
