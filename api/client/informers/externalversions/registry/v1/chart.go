/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	versioned "tkestack.io/tke/api/client/clientset/versioned"
	internalinterfaces "tkestack.io/tke/api/client/informers/externalversions/internalinterfaces"
	v1 "tkestack.io/tke/api/client/listers/registry/v1"
	registryv1 "tkestack.io/tke/api/registry/v1"
)

// ChartInformer provides access to a shared informer and lister for
// Charts.
type ChartInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ChartLister
}

type chartInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewChartInformer constructs a new informer for Chart type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewChartInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredChartInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredChartInformer constructs a new informer for Chart type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredChartInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RegistryV1().Charts(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RegistryV1().Charts(namespace).Watch(options)
			},
		},
		&registryv1.Chart{},
		resyncPeriod,
		indexers,
	)
}

func (f *chartInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredChartInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *chartInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&registryv1.Chart{}, f.defaultInformer)
}

func (f *chartInformer) Lister() v1.ChartLister {
	return v1.NewChartLister(f.Informer().GetIndexer())
}