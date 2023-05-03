/*
Copyright AppsCode Inc. and Contributors

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
	"context"
	"time"

	v1alpha1 "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	scheme "kubedb.dev/apimachinery/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ProxySQLVersionsGetter has a method to return a ProxySQLVersionInterface.
// A group's client should implement this interface.
type ProxySQLVersionsGetter interface {
	ProxySQLVersions() ProxySQLVersionInterface
}

// ProxySQLVersionInterface has methods to work with ProxySQLVersion resources.
type ProxySQLVersionInterface interface {
	Create(ctx context.Context, proxySQLVersion *v1alpha1.ProxySQLVersion, opts v1.CreateOptions) (*v1alpha1.ProxySQLVersion, error)
	Update(ctx context.Context, proxySQLVersion *v1alpha1.ProxySQLVersion, opts v1.UpdateOptions) (*v1alpha1.ProxySQLVersion, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ProxySQLVersion, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ProxySQLVersionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ProxySQLVersion, err error)
	ProxySQLVersionExpansion
}

// proxySQLVersions implements ProxySQLVersionInterface
type proxySQLVersions struct {
	client rest.Interface
}

// newProxySQLVersions returns a ProxySQLVersions
func newProxySQLVersions(c *CatalogV1alpha1Client) *proxySQLVersions {
	return &proxySQLVersions{
		client: c.RESTClient(),
	}
}

// Get takes name of the proxySQLVersion, and returns the corresponding proxySQLVersion object, and an error if there is any.
func (c *proxySQLVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ProxySQLVersion, err error) {
	result = &v1alpha1.ProxySQLVersion{}
	err = c.client.Get().
		Resource("proxysqlversions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ProxySQLVersions that match those selectors.
func (c *proxySQLVersions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ProxySQLVersionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ProxySQLVersionList{}
	err = c.client.Get().
		Resource("proxysqlversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested proxySQLVersions.
func (c *proxySQLVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("proxysqlversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a proxySQLVersion and creates it.  Returns the server's representation of the proxySQLVersion, and an error, if there is any.
func (c *proxySQLVersions) Create(ctx context.Context, proxySQLVersion *v1alpha1.ProxySQLVersion, opts v1.CreateOptions) (result *v1alpha1.ProxySQLVersion, err error) {
	result = &v1alpha1.ProxySQLVersion{}
	err = c.client.Post().
		Resource("proxysqlversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(proxySQLVersion).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a proxySQLVersion and updates it. Returns the server's representation of the proxySQLVersion, and an error, if there is any.
func (c *proxySQLVersions) Update(ctx context.Context, proxySQLVersion *v1alpha1.ProxySQLVersion, opts v1.UpdateOptions) (result *v1alpha1.ProxySQLVersion, err error) {
	result = &v1alpha1.ProxySQLVersion{}
	err = c.client.Put().
		Resource("proxysqlversions").
		Name(proxySQLVersion.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(proxySQLVersion).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the proxySQLVersion and deletes it. Returns an error if one occurs.
func (c *proxySQLVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("proxysqlversions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *proxySQLVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("proxysqlversions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched proxySQLVersion.
func (c *proxySQLVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ProxySQLVersion, err error) {
	result = &v1alpha1.ProxySQLVersion{}
	err = c.client.Patch(pt).
		Resource("proxysqlversions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}