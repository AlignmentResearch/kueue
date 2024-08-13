/*
Copyright 2024 The Kubernetes Authors.

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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "sigs.k8s.io/kueue/cmd/experimental/kjobctl/apis/v1alpha1"
)

// FakeRayClusterTemplates implements RayClusterTemplateInterface
type FakeRayClusterTemplates struct {
	Fake *FakeKjobctlV1alpha1
	ns   string
}

var rayclustertemplatesResource = v1alpha1.SchemeGroupVersion.WithResource("rayclustertemplates")

var rayclustertemplatesKind = v1alpha1.SchemeGroupVersion.WithKind("RayClusterTemplate")

// Get takes name of the rayClusterTemplate, and returns the corresponding rayClusterTemplate object, and an error if there is any.
func (c *FakeRayClusterTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RayClusterTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rayclustertemplatesResource, c.ns, name), &v1alpha1.RayClusterTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RayClusterTemplate), err
}

// List takes label and field selectors, and returns the list of RayClusterTemplates that match those selectors.
func (c *FakeRayClusterTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RayClusterTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rayclustertemplatesResource, rayclustertemplatesKind, c.ns, opts), &v1alpha1.RayClusterTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RayClusterTemplateList{ListMeta: obj.(*v1alpha1.RayClusterTemplateList).ListMeta}
	for _, item := range obj.(*v1alpha1.RayClusterTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rayClusterTemplates.
func (c *FakeRayClusterTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rayclustertemplatesResource, c.ns, opts))

}

// Create takes the representation of a rayClusterTemplate and creates it.  Returns the server's representation of the rayClusterTemplate, and an error, if there is any.
func (c *FakeRayClusterTemplates) Create(ctx context.Context, rayClusterTemplate *v1alpha1.RayClusterTemplate, opts v1.CreateOptions) (result *v1alpha1.RayClusterTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rayclustertemplatesResource, c.ns, rayClusterTemplate), &v1alpha1.RayClusterTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RayClusterTemplate), err
}

// Update takes the representation of a rayClusterTemplate and updates it. Returns the server's representation of the rayClusterTemplate, and an error, if there is any.
func (c *FakeRayClusterTemplates) Update(ctx context.Context, rayClusterTemplate *v1alpha1.RayClusterTemplate, opts v1.UpdateOptions) (result *v1alpha1.RayClusterTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rayclustertemplatesResource, c.ns, rayClusterTemplate), &v1alpha1.RayClusterTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RayClusterTemplate), err
}

// Delete takes name of the rayClusterTemplate and deletes it. Returns an error if one occurs.
func (c *FakeRayClusterTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(rayclustertemplatesResource, c.ns, name, opts), &v1alpha1.RayClusterTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRayClusterTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rayclustertemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RayClusterTemplateList{})
	return err
}

// Patch applies the patch and returns the patched rayClusterTemplate.
func (c *FakeRayClusterTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RayClusterTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rayclustertemplatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.RayClusterTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RayClusterTemplate), err
}