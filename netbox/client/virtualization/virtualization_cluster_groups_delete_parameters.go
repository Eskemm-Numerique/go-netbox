// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewVirtualizationClusterGroupsDeleteParams creates a new VirtualizationClusterGroupsDeleteParams object
// with the default values initialized.
func NewVirtualizationClusterGroupsDeleteParams() *VirtualizationClusterGroupsDeleteParams {
	var ()
	return &VirtualizationClusterGroupsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClusterGroupsDeleteParamsWithTimeout creates a new VirtualizationClusterGroupsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationClusterGroupsDeleteParamsWithTimeout(timeout time.Duration) *VirtualizationClusterGroupsDeleteParams {
	var ()
	return &VirtualizationClusterGroupsDeleteParams{

		timeout: timeout,
	}
}

// NewVirtualizationClusterGroupsDeleteParamsWithContext creates a new VirtualizationClusterGroupsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationClusterGroupsDeleteParamsWithContext(ctx context.Context) *VirtualizationClusterGroupsDeleteParams {
	var ()
	return &VirtualizationClusterGroupsDeleteParams{

		Context: ctx,
	}
}

// NewVirtualizationClusterGroupsDeleteParamsWithHTTPClient creates a new VirtualizationClusterGroupsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationClusterGroupsDeleteParamsWithHTTPClient(client *http.Client) *VirtualizationClusterGroupsDeleteParams {
	var ()
	return &VirtualizationClusterGroupsDeleteParams{
		HTTPClient: client,
	}
}

/*VirtualizationClusterGroupsDeleteParams contains all the parameters to send to the API endpoint
for the virtualization cluster groups delete operation typically these are written to a http.Request
*/
type VirtualizationClusterGroupsDeleteParams struct {

	/*ID
	  A unique integer value identifying this cluster group.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization cluster groups delete params
func (o *VirtualizationClusterGroupsDeleteParams) WithTimeout(timeout time.Duration) *VirtualizationClusterGroupsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization cluster groups delete params
func (o *VirtualizationClusterGroupsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization cluster groups delete params
func (o *VirtualizationClusterGroupsDeleteParams) WithContext(ctx context.Context) *VirtualizationClusterGroupsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization cluster groups delete params
func (o *VirtualizationClusterGroupsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization cluster groups delete params
func (o *VirtualizationClusterGroupsDeleteParams) WithHTTPClient(client *http.Client) *VirtualizationClusterGroupsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization cluster groups delete params
func (o *VirtualizationClusterGroupsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the virtualization cluster groups delete params
func (o *VirtualizationClusterGroupsDeleteParams) WithID(id int64) *VirtualizationClusterGroupsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization cluster groups delete params
func (o *VirtualizationClusterGroupsDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClusterGroupsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
