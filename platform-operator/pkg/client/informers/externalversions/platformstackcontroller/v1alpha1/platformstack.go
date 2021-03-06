/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	platformstackcontrollerv1alpha1 "github.com/cloud-ark/kubeplus/platform-operator/pkg/apis/platformstackcontroller/v1alpha1"
	versioned "github.com/cloud-ark/kubeplus/platform-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/cloud-ark/kubeplus/platform-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/cloud-ark/kubeplus/platform-operator/pkg/client/listers/platformstackcontroller/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PlatformStackInformer provides access to a shared informer and lister for
// PlatformStacks.
type PlatformStackInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PlatformStackLister
}

type platformStackInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPlatformStackInformer constructs a new informer for PlatformStack type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPlatformStackInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPlatformStackInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPlatformStackInformer constructs a new informer for PlatformStack type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPlatformStackInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PlatformstackV1alpha1().PlatformStacks(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PlatformstackV1alpha1().PlatformStacks(namespace).Watch(options)
			},
		},
		&platformstackcontrollerv1alpha1.PlatformStack{},
		resyncPeriod,
		indexers,
	)
}

func (f *platformStackInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPlatformStackInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *platformStackInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&platformstackcontrollerv1alpha1.PlatformStack{}, f.defaultInformer)
}

func (f *platformStackInformer) Lister() v1alpha1.PlatformStackLister {
	return v1alpha1.NewPlatformStackLister(f.Informer().GetIndexer())
}
