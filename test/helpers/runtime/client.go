package runtime

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type FakeClient struct {
	Error   error
	Client  client.Client
	updated []runtime.Object
}

func NewAlreadyExistsException() *errors.StatusError {
	return errors.NewAlreadyExists(schema.GroupResource{}, "existingname")
}

func NewFakeClient(client client.Client, err error) *FakeClient {
	return &FakeClient{
		Error:   err,
		Client:  client,
		updated: []runtime.Object{},
	}
}

func (fw *FakeClient) WasUpdated(name string) bool {
	for _, o := range fw.updated {
		listkey, _ := client.ObjectKeyFromObject(o)
		if listkey.Name == name {
			return true
		}
	}
	return false
}

func (fw *FakeClient) Create(ctx context.Context, obj runtime.Object, opts ...client.CreateOption) error {
	if fw.Error != nil {
		return fw.Error
	}
	return fw.Client.Create(ctx, obj, opts...)
}

func (fw *FakeClient) Delete(ctx context.Context, obj runtime.Object, opts ...client.DeleteOption) error {
	return fw.Error
}

func (fw *FakeClient) DeleteAllOf(ctx context.Context, obj runtime.Object, opts ...client.DeleteAllOfOption) error {
	return fw.Error
}

func (fw *FakeClient) Update(ctx context.Context, obj runtime.Object, opts ...client.UpdateOption) error {
	fw.updated = append(fw.updated, obj)
	return fw.Client.Update(ctx, obj, opts...)
}

func (fw *FakeClient) Get(ctx context.Context, key client.ObjectKey, obj runtime.Object) error {
	return fw.Client.Get(ctx, key, obj)
}

func (fw *FakeClient) List(ctx context.Context, list runtime.Object, opts ...client.ListOption) error {
	return fw.Error
}

func (fw *FakeClient) Patch(ctx context.Context, obj runtime.Object, patch client.Patch, opts ...client.PatchOption) error {
	return fw.Error
}

func (fw *FakeClient) Status() client.StatusWriter {
	return fw.Status()
}
