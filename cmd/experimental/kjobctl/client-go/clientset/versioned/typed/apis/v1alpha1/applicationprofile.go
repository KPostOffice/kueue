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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "sigs.k8s.io/kueue/cmd/experimental/kjobctl/apis/v1alpha1"
	scheme "sigs.k8s.io/kueue/cmd/experimental/kjobctl/client-go/clientset/versioned/scheme"
)

// ApplicationProfilesGetter has a method to return a ApplicationProfileInterface.
// A group's client should implement this interface.
type ApplicationProfilesGetter interface {
	ApplicationProfiles(namespace string) ApplicationProfileInterface
}

// ApplicationProfileInterface has methods to work with ApplicationProfile resources.
type ApplicationProfileInterface interface {
	Create(ctx context.Context, applicationProfile *v1alpha1.ApplicationProfile, opts v1.CreateOptions) (*v1alpha1.ApplicationProfile, error)
	Update(ctx context.Context, applicationProfile *v1alpha1.ApplicationProfile, opts v1.UpdateOptions) (*v1alpha1.ApplicationProfile, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ApplicationProfile, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ApplicationProfileList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApplicationProfile, err error)
	ApplicationProfileExpansion
}

// applicationProfiles implements ApplicationProfileInterface
type applicationProfiles struct {
	client rest.Interface
	ns     string
}

// newApplicationProfiles returns a ApplicationProfiles
func newApplicationProfiles(c *KjobctlV1alpha1Client, namespace string) *applicationProfiles {
	return &applicationProfiles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the applicationProfile, and returns the corresponding applicationProfile object, and an error if there is any.
func (c *applicationProfiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApplicationProfile, err error) {
	result = &v1alpha1.ApplicationProfile{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationprofiles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApplicationProfiles that match those selectors.
func (c *applicationProfiles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApplicationProfileList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApplicationProfileList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested applicationProfiles.
func (c *applicationProfiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("applicationprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a applicationProfile and creates it.  Returns the server's representation of the applicationProfile, and an error, if there is any.
func (c *applicationProfiles) Create(ctx context.Context, applicationProfile *v1alpha1.ApplicationProfile, opts v1.CreateOptions) (result *v1alpha1.ApplicationProfile, err error) {
	result = &v1alpha1.ApplicationProfile{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("applicationprofiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(applicationProfile).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a applicationProfile and updates it. Returns the server's representation of the applicationProfile, and an error, if there is any.
func (c *applicationProfiles) Update(ctx context.Context, applicationProfile *v1alpha1.ApplicationProfile, opts v1.UpdateOptions) (result *v1alpha1.ApplicationProfile, err error) {
	result = &v1alpha1.ApplicationProfile{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationprofiles").
		Name(applicationProfile.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(applicationProfile).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the applicationProfile and deletes it. Returns an error if one occurs.
func (c *applicationProfiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationprofiles").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *applicationProfiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationprofiles").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched applicationProfile.
func (c *applicationProfiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApplicationProfile, err error) {
	result = &v1alpha1.ApplicationProfile{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("applicationprofiles").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
