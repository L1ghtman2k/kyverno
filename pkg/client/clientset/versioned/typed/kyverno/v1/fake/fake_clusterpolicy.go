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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	kyvernov1 "github.com/kyverno/kyverno/api/kyverno/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterPolicies implements ClusterPolicyInterface
type FakeClusterPolicies struct {
	Fake *FakeKyvernoV1
}

var clusterpoliciesResource = schema.GroupVersionResource{Group: "kyverno.io", Version: "v1", Resource: "clusterpolicies"}

var clusterpoliciesKind = schema.GroupVersionKind{Group: "kyverno.io", Version: "v1", Kind: "ClusterPolicy"}

// Get takes name of the clusterPolicy, and returns the corresponding clusterPolicy object, and an error if there is any.
func (c *FakeClusterPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *kyvernov1.ClusterPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterpoliciesResource, name), &kyvernov1.ClusterPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kyvernov1.ClusterPolicy), err
}

// List takes label and field selectors, and returns the list of ClusterPolicies that match those selectors.
func (c *FakeClusterPolicies) List(ctx context.Context, opts v1.ListOptions) (result *kyvernov1.ClusterPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterpoliciesResource, clusterpoliciesKind, opts), &kyvernov1.ClusterPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kyvernov1.ClusterPolicyList{ListMeta: obj.(*kyvernov1.ClusterPolicyList).ListMeta}
	for _, item := range obj.(*kyvernov1.ClusterPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterPolicies.
func (c *FakeClusterPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterpoliciesResource, opts))
}

// Create takes the representation of a clusterPolicy and creates it.  Returns the server's representation of the clusterPolicy, and an error, if there is any.
func (c *FakeClusterPolicies) Create(ctx context.Context, clusterPolicy *kyvernov1.ClusterPolicy, opts v1.CreateOptions) (result *kyvernov1.ClusterPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterpoliciesResource, clusterPolicy), &kyvernov1.ClusterPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kyvernov1.ClusterPolicy), err
}

// Update takes the representation of a clusterPolicy and updates it. Returns the server's representation of the clusterPolicy, and an error, if there is any.
func (c *FakeClusterPolicies) Update(ctx context.Context, clusterPolicy *kyvernov1.ClusterPolicy, opts v1.UpdateOptions) (result *kyvernov1.ClusterPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterpoliciesResource, clusterPolicy), &kyvernov1.ClusterPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kyvernov1.ClusterPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterPolicies) UpdateStatus(ctx context.Context, clusterPolicy *kyvernov1.ClusterPolicy, opts v1.UpdateOptions) (*kyvernov1.ClusterPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterpoliciesResource, "status", clusterPolicy), &kyvernov1.ClusterPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kyvernov1.ClusterPolicy), err
}

// Delete takes name of the clusterPolicy and deletes it. Returns an error if one occurs.
func (c *FakeClusterPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterpoliciesResource, name), &kyvernov1.ClusterPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterpoliciesResource, listOpts)

	_, err := c.Fake.Invokes(action, &kyvernov1.ClusterPolicyList{})
	return err
}

// Patch applies the patch and returns the patched clusterPolicy.
func (c *FakeClusterPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kyvernov1.ClusterPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterpoliciesResource, name, pt, data, subresources...), &kyvernov1.ClusterPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kyvernov1.ClusterPolicy), err
}