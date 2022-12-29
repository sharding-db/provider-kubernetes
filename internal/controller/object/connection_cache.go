package object

import (
	"time"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/patrickmn/go-cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var globalK8sClientCache = cache.New(cache.NoExpiration, time.Hour)

func getClientFromCache(providerConfig string) *resource.ClientApplicator {
	if k, ok := globalK8sClientCache.Get(providerConfig); ok {
		return k.(*resource.ClientApplicator)
	}
	return nil
}

func setClientToCache(providerConfig string, cli client.Client) {
	globalK8sClientCache.Set(providerConfig, &resource.ClientApplicator{
		Client:     cli,
		Applicator: resource.NewAPIPatchingApplicator(cli),
	}, 0)
}
