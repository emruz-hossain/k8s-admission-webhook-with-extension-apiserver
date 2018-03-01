/*
Copyright 2018 The Voyager Authors.

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

package v1alpha1

import (
	kubecar_v1alpha1 "k8s-admission-webhook-with-extension-apiserver/apis/kubecar/v1alpha1"
	versioned "k8s-admission-webhook-with-extension-apiserver/client/clientset/versioned"
	internalinterfaces "k8s-admission-webhook-with-extension-apiserver/client/informers/externalversions/internalinterfaces"
	v1alpha1 "k8s-admission-webhook-with-extension-apiserver/client/listers/kubecar/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// KubecarInformer provides access to a shared informer and lister for
// Kubecars.
type KubecarInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KubecarLister
}

type kubecarInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKubecarInformer constructs a new informer for Kubecar type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKubecarInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKubecarInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKubecarInformer constructs a new informer for Kubecar type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKubecarInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubecarV1alpha1().Kubecars(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubecarV1alpha1().Kubecars(namespace).Watch(options)
			},
		},
		&kubecar_v1alpha1.Kubecar{},
		resyncPeriod,
		indexers,
	)
}

func (f *kubecarInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKubecarInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kubecarInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubecar_v1alpha1.Kubecar{}, f.defaultInformer)
}

func (f *kubecarInformer) Lister() v1alpha1.KubecarLister {
	return v1alpha1.NewKubecarLister(f.Informer().GetIndexer())
}