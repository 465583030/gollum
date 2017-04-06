// Copyright 2015-2017 trivago GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package router

import (
	"github.com/trivago/gollum/core"
)

// Route stream plugin
// Messages will be routed to all streams configured. Each target stream can
// hold another stream configuration, too, so this is not directly sending to
// the producers attached to the target streams.
// Configuration example
//
// 	myrouter":
//    Type: "router.Route"
//    Stream: "mystream"
// 	  Routes:
//      - "foo"
//      - "bar"
//
// Routes defines a 1:n stream remapping.
// Messages are reassigned to all of stream(s) in this list.
// If no route is set messages are forwarded on the incoming router.
// When routing to multiple streams, the incoming stream has to be listed explicitly to be used.
type Route struct {
	Broadcast
	streams []core.Router
}

func init() {
	core.TypeRegistry.Register(Route{})
}

// Configure initializes this distributor with values from a plugin config.
func (router *Route) Configure(conf core.PluginConfigReader) error {
	router.Broadcast.Configure(conf)

	boundStreamIDs := conf.GetStreamArray("Streams", []core.MessageStreamID{})
	for _, streamID := range boundStreamIDs {
		route := core.StreamRegistry.GetStreamOrFallback(streamID)
		router.streams = append(router.streams, route)
	}

	return conf.Errors.OrNil()
}

func (router *Route) route(msg *core.Message, route core.Router) {
	if router.StreamID() == router.StreamID() {
		router.Broadcast.Enqueue(msg)
	} else {
		msg.SetStreamID(route.StreamID())
		core.Route(msg, route)
	}
}

// Enqueue enques a message to the router
func (router *Route) Enqueue(msg *core.Message) error {
	numStreams := len(router.streams)

	switch numStreams {
	case 0:
		return core.NewModulateResultError("No producers configured for stream %s", router.GetID())

	case 1:
		router.route(msg, router.streams[0])

	default:
		lastStreamIdx := numStreams - 1
		for streamIdx := 0; streamIdx < lastStreamIdx; streamIdx++ {
			router.route(msg.Clone(), router.streams[streamIdx])
		}
		router.route(msg, router.streams[lastStreamIdx])
	}

	return nil
}
