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

package secrets

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
)

// NewSecretsChoicesListParams creates a new SecretsChoicesListParams object
// with the default values initialized.
func NewSecretsChoicesListParams() *SecretsChoicesListParams {

	return &SecretsChoicesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSecretsChoicesListParamsWithTimeout creates a new SecretsChoicesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSecretsChoicesListParamsWithTimeout(timeout time.Duration) *SecretsChoicesListParams {

	return &SecretsChoicesListParams{

		timeout: timeout,
	}
}

// NewSecretsChoicesListParamsWithContext creates a new SecretsChoicesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewSecretsChoicesListParamsWithContext(ctx context.Context) *SecretsChoicesListParams {

	return &SecretsChoicesListParams{

		Context: ctx,
	}
}

// NewSecretsChoicesListParamsWithHTTPClient creates a new SecretsChoicesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSecretsChoicesListParamsWithHTTPClient(client *http.Client) *SecretsChoicesListParams {

	return &SecretsChoicesListParams{
		HTTPClient: client,
	}
}

/*SecretsChoicesListParams contains all the parameters to send to the API endpoint
for the secrets choices list operation typically these are written to a http.Request
*/
type SecretsChoicesListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the secrets choices list params
func (o *SecretsChoicesListParams) WithTimeout(timeout time.Duration) *SecretsChoicesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secrets choices list params
func (o *SecretsChoicesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secrets choices list params
func (o *SecretsChoicesListParams) WithContext(ctx context.Context) *SecretsChoicesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secrets choices list params
func (o *SecretsChoicesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secrets choices list params
func (o *SecretsChoicesListParams) WithHTTPClient(client *http.Client) *SecretsChoicesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secrets choices list params
func (o *SecretsChoicesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SecretsChoicesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
