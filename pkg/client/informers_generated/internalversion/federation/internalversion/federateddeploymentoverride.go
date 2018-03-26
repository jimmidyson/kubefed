/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package internalversion

import (
	federation "github.com/marun/fnord/pkg/apis/federation"
	internalclientset "github.com/marun/fnord/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "github.com/marun/fnord/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "github.com/marun/fnord/pkg/client/listers_generated/federation/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// FederatedDeploymentOverrideInformer provides access to a shared informer and lister for
// FederatedDeploymentOverrides.
type FederatedDeploymentOverrideInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.FederatedDeploymentOverrideLister
}

type federatedDeploymentOverrideInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFederatedDeploymentOverrideInformer constructs a new informer for FederatedDeploymentOverride type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFederatedDeploymentOverrideInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFederatedDeploymentOverrideInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFederatedDeploymentOverrideInformer constructs a new informer for FederatedDeploymentOverride type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFederatedDeploymentOverrideInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Federation().FederatedDeploymentOverrides(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Federation().FederatedDeploymentOverrides(namespace).Watch(options)
			},
		},
		&federation.FederatedDeploymentOverride{},
		resyncPeriod,
		indexers,
	)
}

func (f *federatedDeploymentOverrideInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFederatedDeploymentOverrideInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *federatedDeploymentOverrideInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&federation.FederatedDeploymentOverride{}, f.defaultInformer)
}

func (f *federatedDeploymentOverrideInformer) Lister() internalversion.FederatedDeploymentOverrideLister {
	return internalversion.NewFederatedDeploymentOverrideLister(f.Informer().GetIndexer())
}
