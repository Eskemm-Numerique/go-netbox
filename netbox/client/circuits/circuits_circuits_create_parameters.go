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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// NewCircuitsCircuitsCreateParams creates a new CircuitsCircuitsCreateParams object
// with the default values initialized.
func NewCircuitsCircuitsCreateParams() *CircuitsCircuitsCreateParams {
	var ()
	return &CircuitsCircuitsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitsCreateParamsWithTimeout creates a new CircuitsCircuitsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCircuitsCircuitsCreateParamsWithTimeout(timeout time.Duration) *CircuitsCircuitsCreateParams {
	var ()
	return &CircuitsCircuitsCreateParams{

		timeout: timeout,
	}
}

// NewCircuitsCircuitsCreateParamsWithContext creates a new CircuitsCircuitsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCircuitsCircuitsCreateParamsWithContext(ctx context.Context) *CircuitsCircuitsCreateParams {
	var ()
	return &CircuitsCircuitsCreateParams{

		Context: ctx,
	}
}

// NewCircuitsCircuitsCreateParamsWithHTTPClient creates a new CircuitsCircuitsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCircuitsCircuitsCreateParamsWithHTTPClient(client *http.Client) *CircuitsCircuitsCreateParams {
	var ()
	return &CircuitsCircuitsCreateParams{
		HTTPClient: client,
	}
}

/*CircuitsCircuitsCreateParams contains all the parameters to send to the API endpoint
for the circuits circuits create operation typically these are written to a http.Request
*/
type CircuitsCircuitsCreateParams struct {

	/*Data*/
	Data *models.WritableCircuit

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the circuits circuits create params
func (o *CircuitsCircuitsCreateParams) WithTimeout(timeout time.Duration) *CircuitsCircuitsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuits create params
func (o *CircuitsCircuitsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuits create params
func (o *CircuitsCircuitsCreateParams) WithContext(ctx context.Context) *CircuitsCircuitsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuits create params
func (o *CircuitsCircuitsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuits create params
func (o *CircuitsCircuitsCreateParams) WithHTTPClient(client *http.Client) *CircuitsCircuitsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuits create params
func (o *CircuitsCircuitsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits circuits create params
func (o *CircuitsCircuitsCreateParams) WithData(data *models.WritableCircuit) *CircuitsCircuitsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits circuits create params
func (o *CircuitsCircuitsCreateParams) SetData(data *models.WritableCircuit) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
