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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewVirtualizationInterfacesReadParams creates a new VirtualizationInterfacesReadParams object
// with the default values initialized.
func NewVirtualizationInterfacesReadParams() *VirtualizationInterfacesReadParams {
	var ()
	return &VirtualizationInterfacesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationInterfacesReadParamsWithTimeout creates a new VirtualizationInterfacesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationInterfacesReadParamsWithTimeout(timeout time.Duration) *VirtualizationInterfacesReadParams {
	var ()
	return &VirtualizationInterfacesReadParams{

		timeout: timeout,
	}
}

// NewVirtualizationInterfacesReadParamsWithContext creates a new VirtualizationInterfacesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationInterfacesReadParamsWithContext(ctx context.Context) *VirtualizationInterfacesReadParams {
	var ()
	return &VirtualizationInterfacesReadParams{

		Context: ctx,
	}
}

// NewVirtualizationInterfacesReadParamsWithHTTPClient creates a new VirtualizationInterfacesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationInterfacesReadParamsWithHTTPClient(client *http.Client) *VirtualizationInterfacesReadParams {
	var ()
	return &VirtualizationInterfacesReadParams{
		HTTPClient: client,
	}
}

/*VirtualizationInterfacesReadParams contains all the parameters to send to the API endpoint
for the virtualization interfaces read operation typically these are written to a http.Request
*/
type VirtualizationInterfacesReadParams struct {

	/*ID
	  A unique integer value identifying this interface.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) WithTimeout(timeout time.Duration) *VirtualizationInterfacesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) WithContext(ctx context.Context) *VirtualizationInterfacesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) WithHTTPClient(client *http.Client) *VirtualizationInterfacesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) WithID(id int64) *VirtualizationInterfacesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization interfaces read params
func (o *VirtualizationInterfacesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationInterfacesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
