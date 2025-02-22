/*
Copyright The KubeEdge Authors.

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

	rulesv1 "github.com/kubeedge/kubeedge/pkg/apis/rules/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRuleEndpoints implements RuleEndpointInterface
type FakeRuleEndpoints struct {
	Fake *FakeRulesV1
	ns   string
}

var ruleendpointsResource = schema.GroupVersionResource{Group: "rules.kubeedge.io", Version: "v1", Resource: "ruleendpoints"}

var ruleendpointsKind = schema.GroupVersionKind{Group: "rules.kubeedge.io", Version: "v1", Kind: "RuleEndpoint"}

// Get takes name of the ruleEndpoint, and returns the corresponding ruleEndpoint object, and an error if there is any.
func (c *FakeRuleEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *rulesv1.RuleEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ruleendpointsResource, c.ns, name), &rulesv1.RuleEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rulesv1.RuleEndpoint), err
}

// List takes label and field selectors, and returns the list of RuleEndpoints that match those selectors.
func (c *FakeRuleEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *rulesv1.RuleEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ruleendpointsResource, ruleendpointsKind, c.ns, opts), &rulesv1.RuleEndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &rulesv1.RuleEndpointList{ListMeta: obj.(*rulesv1.RuleEndpointList).ListMeta}
	for _, item := range obj.(*rulesv1.RuleEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ruleEndpoints.
func (c *FakeRuleEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ruleendpointsResource, c.ns, opts))

}

// Create takes the representation of a ruleEndpoint and creates it.  Returns the server's representation of the ruleEndpoint, and an error, if there is any.
func (c *FakeRuleEndpoints) Create(ctx context.Context, ruleEndpoint *rulesv1.RuleEndpoint, opts v1.CreateOptions) (result *rulesv1.RuleEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ruleendpointsResource, c.ns, ruleEndpoint), &rulesv1.RuleEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rulesv1.RuleEndpoint), err
}

// Update takes the representation of a ruleEndpoint and updates it. Returns the server's representation of the ruleEndpoint, and an error, if there is any.
func (c *FakeRuleEndpoints) Update(ctx context.Context, ruleEndpoint *rulesv1.RuleEndpoint, opts v1.UpdateOptions) (result *rulesv1.RuleEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ruleendpointsResource, c.ns, ruleEndpoint), &rulesv1.RuleEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rulesv1.RuleEndpoint), err
}

// Delete takes name of the ruleEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeRuleEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ruleendpointsResource, c.ns, name), &rulesv1.RuleEndpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRuleEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ruleendpointsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &rulesv1.RuleEndpointList{})
	return err
}

// Patch applies the patch and returns the patched ruleEndpoint.
func (c *FakeRuleEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *rulesv1.RuleEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ruleendpointsResource, c.ns, name, pt, data, subresources...), &rulesv1.RuleEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*rulesv1.RuleEndpoint), err
}
