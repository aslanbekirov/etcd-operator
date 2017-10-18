/*
Copyright 2017 The etcd-operator Authors

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
package fake

import (
	v1beta2 "github.com/coreos/etcd-operator/pkg/apis/etcd/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEtcdRestores implements EtcdRestoreInterface
type FakeEtcdRestores struct {
	Fake *FakeEtcdV1beta2
	ns   string
}

var etcdrestoresResource = schema.GroupVersionResource{Group: "etcd.database.coreos.com", Version: "v1beta2", Resource: "etcdrestores"}

var etcdrestoresKind = schema.GroupVersionKind{Group: "etcd.database.coreos.com", Version: "v1beta2", Kind: "EtcdRestore"}

// Get takes name of the etcdRestore, and returns the corresponding etcdRestore object, and an error if there is any.
func (c *FakeEtcdRestores) Get(name string, options v1.GetOptions) (result *v1beta2.EtcdRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(etcdrestoresResource, c.ns, name), &v1beta2.EtcdRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.EtcdRestore), err
}

// List takes label and field selectors, and returns the list of EtcdRestores that match those selectors.
func (c *FakeEtcdRestores) List(opts v1.ListOptions) (result *v1beta2.EtcdRestoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(etcdrestoresResource, etcdrestoresKind, c.ns, opts), &v1beta2.EtcdRestoreList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.EtcdRestoreList{}
	for _, item := range obj.(*v1beta2.EtcdRestoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested etcdRestores.
func (c *FakeEtcdRestores) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(etcdrestoresResource, c.ns, opts))

}

// Create takes the representation of a etcdRestore and creates it.  Returns the server's representation of the etcdRestore, and an error, if there is any.
func (c *FakeEtcdRestores) Create(etcdRestore *v1beta2.EtcdRestore) (result *v1beta2.EtcdRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(etcdrestoresResource, c.ns, etcdRestore), &v1beta2.EtcdRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.EtcdRestore), err
}

// Update takes the representation of a etcdRestore and updates it. Returns the server's representation of the etcdRestore, and an error, if there is any.
func (c *FakeEtcdRestores) Update(etcdRestore *v1beta2.EtcdRestore) (result *v1beta2.EtcdRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(etcdrestoresResource, c.ns, etcdRestore), &v1beta2.EtcdRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.EtcdRestore), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEtcdRestores) UpdateStatus(etcdRestore *v1beta2.EtcdRestore) (*v1beta2.EtcdRestore, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(etcdrestoresResource, "status", c.ns, etcdRestore), &v1beta2.EtcdRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.EtcdRestore), err
}

// Delete takes name of the etcdRestore and deletes it. Returns an error if one occurs.
func (c *FakeEtcdRestores) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(etcdrestoresResource, c.ns, name), &v1beta2.EtcdRestore{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEtcdRestores) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(etcdrestoresResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta2.EtcdRestoreList{})
	return err
}

// Patch applies the patch and returns the patched etcdRestore.
func (c *FakeEtcdRestores) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta2.EtcdRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(etcdrestoresResource, c.ns, name, data, subresources...), &v1beta2.EtcdRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.EtcdRestore), err
}