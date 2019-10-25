/*
Copyright 2019 The Kube Vault Authors.

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

package v1alpha1

import (
	"time"

	v1alpha1 "kubevault.dev/operator/apis/engine/v1alpha1"
	scheme "kubevault.dev/operator/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GCPRolesGetter has a method to return a GCPRoleInterface.
// A group's client should implement this interface.
type GCPRolesGetter interface {
	GCPRoles(namespace string) GCPRoleInterface
}

// GCPRoleInterface has methods to work with GCPRole resources.
type GCPRoleInterface interface {
	Create(*v1alpha1.GCPRole) (*v1alpha1.GCPRole, error)
	Update(*v1alpha1.GCPRole) (*v1alpha1.GCPRole, error)
	UpdateStatus(*v1alpha1.GCPRole) (*v1alpha1.GCPRole, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GCPRole, error)
	List(opts v1.ListOptions) (*v1alpha1.GCPRoleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GCPRole, err error)
	GCPRoleExpansion
}

// gCPRoles implements GCPRoleInterface
type gCPRoles struct {
	client rest.Interface
	ns     string
}

// newGCPRoles returns a GCPRoles
func newGCPRoles(c *EngineV1alpha1Client, namespace string) *gCPRoles {
	return &gCPRoles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gCPRole, and returns the corresponding gCPRole object, and an error if there is any.
func (c *gCPRoles) Get(name string, options v1.GetOptions) (result *v1alpha1.GCPRole, err error) {
	result = &v1alpha1.GCPRole{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gcproles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GCPRoles that match those selectors.
func (c *gCPRoles) List(opts v1.ListOptions) (result *v1alpha1.GCPRoleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GCPRoleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gcproles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gCPRoles.
func (c *gCPRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gcproles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a gCPRole and creates it.  Returns the server's representation of the gCPRole, and an error, if there is any.
func (c *gCPRoles) Create(gCPRole *v1alpha1.GCPRole) (result *v1alpha1.GCPRole, err error) {
	result = &v1alpha1.GCPRole{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gcproles").
		Body(gCPRole).
		Do().
		Into(result)
	return
}

// Update takes the representation of a gCPRole and updates it. Returns the server's representation of the gCPRole, and an error, if there is any.
func (c *gCPRoles) Update(gCPRole *v1alpha1.GCPRole) (result *v1alpha1.GCPRole, err error) {
	result = &v1alpha1.GCPRole{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gcproles").
		Name(gCPRole.Name).
		Body(gCPRole).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *gCPRoles) UpdateStatus(gCPRole *v1alpha1.GCPRole) (result *v1alpha1.GCPRole, err error) {
	result = &v1alpha1.GCPRole{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gcproles").
		Name(gCPRole.Name).
		SubResource("status").
		Body(gCPRole).
		Do().
		Into(result)
	return
}

// Delete takes name of the gCPRole and deletes it. Returns an error if one occurs.
func (c *gCPRoles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gcproles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gCPRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gcproles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched gCPRole.
func (c *gCPRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GCPRole, err error) {
	result = &v1alpha1.GCPRole{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gcproles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
