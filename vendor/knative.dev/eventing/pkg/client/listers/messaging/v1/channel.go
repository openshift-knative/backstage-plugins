/*
Copyright 2021 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
	v1 "knative.dev/eventing/pkg/apis/messaging/v1"
)

// ChannelLister helps list Channels.
// All objects returned here must be treated as read-only.
type ChannelLister interface {
	// List lists all Channels in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Channel, err error)
	// Channels returns an object that can list and get Channels.
	Channels(namespace string) ChannelNamespaceLister
	ChannelListerExpansion
}

// channelLister implements the ChannelLister interface.
type channelLister struct {
	listers.ResourceIndexer[*v1.Channel]
}

// NewChannelLister returns a new ChannelLister.
func NewChannelLister(indexer cache.Indexer) ChannelLister {
	return &channelLister{listers.New[*v1.Channel](indexer, v1.Resource("channel"))}
}

// Channels returns an object that can list and get Channels.
func (s *channelLister) Channels(namespace string) ChannelNamespaceLister {
	return channelNamespaceLister{listers.NewNamespaced[*v1.Channel](s.ResourceIndexer, namespace)}
}

// ChannelNamespaceLister helps list and get Channels.
// All objects returned here must be treated as read-only.
type ChannelNamespaceLister interface {
	// List lists all Channels in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Channel, err error)
	// Get retrieves the Channel from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Channel, error)
	ChannelNamespaceListerExpansion
}

// channelNamespaceLister implements the ChannelNamespaceLister
// interface.
type channelNamespaceLister struct {
	listers.ResourceIndexer[*v1.Channel]
}
