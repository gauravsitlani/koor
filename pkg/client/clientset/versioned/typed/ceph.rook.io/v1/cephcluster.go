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

package v1

import (
	"context"
	"time"

	v1 "github.com/koor-tech/koor/pkg/apis/ceph.rook.io/v1"
	scheme "github.com/koor-tech/koor/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CephClustersGetter has a method to return a CephClusterInterface.
// A group's client should implement this interface.
type CephClustersGetter interface {
	CephClusters(namespace string) CephClusterInterface
}

// CephClusterInterface has methods to work with CephCluster resources.
type CephClusterInterface interface {
	Create(ctx context.Context, cephCluster *v1.CephCluster, opts metav1.CreateOptions) (*v1.CephCluster, error)
	Update(ctx context.Context, cephCluster *v1.CephCluster, opts metav1.UpdateOptions) (*v1.CephCluster, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.CephCluster, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CephClusterList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CephCluster, err error)
	CephClusterExpansion
}

// cephClusters implements CephClusterInterface
type cephClusters struct {
	client rest.Interface
	ns     string
}

// newCephClusters returns a CephClusters
func newCephClusters(c *CephV1Client, namespace string) *cephClusters {
	return &cephClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cephCluster, and returns the corresponding cephCluster object, and an error if there is any.
func (c *cephClusters) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.CephCluster, err error) {
	result = &v1.CephCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cephclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CephClusters that match those selectors.
func (c *cephClusters) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CephClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CephClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cephclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cephClusters.
func (c *cephClusters) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cephclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cephCluster and creates it.  Returns the server's representation of the cephCluster, and an error, if there is any.
func (c *cephClusters) Create(ctx context.Context, cephCluster *v1.CephCluster, opts metav1.CreateOptions) (result *v1.CephCluster, err error) {
	result = &v1.CephCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cephclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cephCluster).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cephCluster and updates it. Returns the server's representation of the cephCluster, and an error, if there is any.
func (c *cephClusters) Update(ctx context.Context, cephCluster *v1.CephCluster, opts metav1.UpdateOptions) (result *v1.CephCluster, err error) {
	result = &v1.CephCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cephclusters").
		Name(cephCluster.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cephCluster).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cephCluster and deletes it. Returns an error if one occurs.
func (c *cephClusters) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cephclusters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cephClusters) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cephclusters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cephCluster.
func (c *cephClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CephCluster, err error) {
	result = &v1.CephCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cephclusters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
