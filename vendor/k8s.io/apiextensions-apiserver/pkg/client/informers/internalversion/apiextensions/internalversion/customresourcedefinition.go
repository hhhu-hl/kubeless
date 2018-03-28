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
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	internalclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/internalclientset"
	internalinterfaces "k8s.io/apiextensions-apiserver/pkg/client/informers/internalversion/internalinterfaces"
	internalversion "k8s.io/apiextensions-apiserver/pkg/client/listers/apiextensions/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// CustomResourceDefinitionInformer provides access to a shared informer and lister for
// CustomResourceDefinitions.
type CustomResourceDefinitionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.CustomResourceDefinitionLister
}

type customResourceDefinitionInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewCustomResourceDefinitionInformer constructs a new informer for CustomResourceDefinition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCustomResourceDefinitionInformer(client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Apiextensions().CustomResourceDefinitions().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Apiextensions().CustomResourceDefinitions().Watch(options)
			},
		},
		&apiextensions.CustomResourceDefinition{},
		resyncPeriod,
		indexers,
	)
}

func defaultCustomResourceDefinitionInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewCustomResourceDefinitionInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *customResourceDefinitionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiextensions.CustomResourceDefinition{}, defaultCustomResourceDefinitionInformer)
}

func (f *customResourceDefinitionInformer) Lister() internalversion.CustomResourceDefinitionLister {
	return internalversion.NewCustomResourceDefinitionLister(f.Informer().GetIndexer())
}