package controllers

import (
	"context"

	"github.com/kong/go-kong/kong"
	"sigs.k8s.io/controller-runtime/pkg/client"

	k8sobj "github.com/kong/kubernetes-ingress-controller/v2/internal/util/kubernetes/object"
)

// DataPlane is a common interface that is used by reconcilers to interact
// with the Kong dataplane.
type DataPlane interface {
	Listeners(ctx context.Context) ([]kong.ProxyListener, []kong.StreamListener, error)
	AreKubernetesObjectReportsEnabled() bool
	KubernetesObjectConfigurationStatus(obj client.Object) k8sobj.ConfigurationStatus
	KubernetesObjectIsConfigured(obj client.Object) bool
}
