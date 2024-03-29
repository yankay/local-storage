// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	versioned "github.com/hwameistor/local-storage/pkg/apis/client/clientset/versioned"
	internalinterfaces "github.com/hwameistor/local-storage/pkg/apis/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/hwameistor/local-storage/pkg/apis/client/listers/localstorage/v1alpha1"
	localstoragev1alpha1 "github.com/hwameistor/local-storage/pkg/apis/localstorage/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// LocalStorageAlertInformer provides access to a shared informer and lister for
// LocalStorageAlerts.
type LocalStorageAlertInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.LocalStorageAlertLister
}

type localStorageAlertInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewLocalStorageAlertInformer constructs a new informer for LocalStorageAlert type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLocalStorageAlertInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLocalStorageAlertInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredLocalStorageAlertInformer constructs a new informer for LocalStorageAlert type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLocalStorageAlertInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LocalStorageV1alpha1().LocalStorageAlerts().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LocalStorageV1alpha1().LocalStorageAlerts().Watch(context.TODO(), options)
			},
		},
		&localstoragev1alpha1.LocalStorageAlert{},
		resyncPeriod,
		indexers,
	)
}

func (f *localStorageAlertInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLocalStorageAlertInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *localStorageAlertInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&localstoragev1alpha1.LocalStorageAlert{}, f.defaultInformer)
}

func (f *localStorageAlertInformer) Lister() v1alpha1.LocalStorageAlertLister {
	return v1alpha1.NewLocalStorageAlertLister(f.Informer().GetIndexer())
}
