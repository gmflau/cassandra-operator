/*
Copyright 2018 The etcd-operator Authors

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

package v1beta2

import (
	cassandra_v1beta2 "github.com/gmflau/cassandra-operator/pkg/apis/cassandra/v1beta2"
	versioned "github.com/gmflau/cassandra-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/gmflau/cassandra-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1beta2 "github.com/gmflau/cassandra-operator/pkg/generated/listers/cassandra/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// CassandraClusterInformer provides access to a shared informer and lister for
// CassandraClusters.
type CassandraClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta2.CassandraClusterLister
}

type cassandraClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewCassandraClusterInformer constructs a new informer for CassandraCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCassandraClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.CassandraV1beta2().CassandraClusters(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.CassandraV1beta2().CassandraClusters(namespace).Watch(options)
			},
		},
		&cassandra_v1beta2.CassandraCluster{},
		resyncPeriod,
		indexers,
	)
}

func defaultCassandraClusterInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewCassandraClusterInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *cassandraClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cassandra_v1beta2.CassandraCluster{}, defaultCassandraClusterInformer)
}

func (f *cassandraClusterInformer) Lister() v1beta2.CassandraClusterLister {
	return v1beta2.NewCassandraClusterLister(f.Informer().GetIndexer())
}
