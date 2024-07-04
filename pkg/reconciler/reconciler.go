package reconciler

import (
	"context"

	"github.com/zncdatadev/operator-go/pkg/client"
	"k8s.io/apimachinery/pkg/runtime"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
)

type AnySpec any

type Reconciler interface {
	GetName() string
	GetNamespace() string
	GetClient() *client.Client
	GetCtrlClient() ctrlclient.Client
	GetCtrlScheme() *runtime.Scheme
	Reconcile(ctx context.Context) *Result
	Ready(ctx context.Context) *Result
}

var _ Reconciler = &BaseReconciler[AnySpec]{}

type BaseReconciler[T AnySpec] struct {
	// Do not use ptr, to avoid other packages to modify the client
	Client *client.Client

	Name string

	Spec T
}

func (b *BaseReconciler[T]) GetClient() *client.Client {
	return b.Client
}

func (b *BaseReconciler[T]) GetName() string {
	return b.Name
}

func (b *BaseReconciler[T]) GetNamespace() string {
	return b.Client.GetOwnerNamespace()
}

func (b *BaseReconciler[T]) GetCtrlClient() ctrlclient.Client {
	return b.Client.GetCtrlClient()
}

func (b *BaseReconciler[T]) GetCtrlScheme() *runtime.Scheme {
	return b.Client.GetCtrlScheme()
}

func (b *BaseReconciler[T]) Ready(ctx context.Context) *Result {
	panic("unimplemented")
}

func (b *BaseReconciler[T]) Reconcile(ctx context.Context) *Result {
	panic("unimplemented")
}

func (b *BaseReconciler[T]) GetSpec() T {
	return b.Spec
}
