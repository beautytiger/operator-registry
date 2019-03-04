// Code generated by go-swagger; DO NOT EDIT.

package package_appr

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewShowPackageManifestsParams creates a new ShowPackageManifestsParams object
// with the default values initialized.
func NewShowPackageManifestsParams() *ShowPackageManifestsParams {
	var ()
	return &ShowPackageManifestsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewShowPackageManifestsParamsWithTimeout creates a new ShowPackageManifestsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewShowPackageManifestsParamsWithTimeout(timeout time.Duration) *ShowPackageManifestsParams {
	var ()
	return &ShowPackageManifestsParams{

		timeout: timeout,
	}
}

// NewShowPackageManifestsParamsWithContext creates a new ShowPackageManifestsParams object
// with the default values initialized, and the ability to set a context for a request
func NewShowPackageManifestsParamsWithContext(ctx context.Context) *ShowPackageManifestsParams {
	var ()
	return &ShowPackageManifestsParams{

		Context: ctx,
	}
}

// NewShowPackageManifestsParamsWithHTTPClient creates a new ShowPackageManifestsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewShowPackageManifestsParamsWithHTTPClient(client *http.Client) *ShowPackageManifestsParams {
	var ()
	return &ShowPackageManifestsParams{
		HTTPClient: client,
	}
}

/*ShowPackageManifestsParams contains all the parameters to send to the API endpoint
for the show package manifests operation typically these are written to a http.Request
*/
type ShowPackageManifestsParams struct {

	/*Namespace
	  namespace

	*/
	Namespace string
	/*Package
	  package name

	*/
	Package string
	/*Release
	  release name

	*/
	Release string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the show package manifests params
func (o *ShowPackageManifestsParams) WithTimeout(timeout time.Duration) *ShowPackageManifestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the show package manifests params
func (o *ShowPackageManifestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the show package manifests params
func (o *ShowPackageManifestsParams) WithContext(ctx context.Context) *ShowPackageManifestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the show package manifests params
func (o *ShowPackageManifestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the show package manifests params
func (o *ShowPackageManifestsParams) WithHTTPClient(client *http.Client) *ShowPackageManifestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the show package manifests params
func (o *ShowPackageManifestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the show package manifests params
func (o *ShowPackageManifestsParams) WithNamespace(namespace string) *ShowPackageManifestsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the show package manifests params
func (o *ShowPackageManifestsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPackage adds the packageVar to the show package manifests params
func (o *ShowPackageManifestsParams) WithPackage(packageVar string) *ShowPackageManifestsParams {
	o.SetPackage(packageVar)
	return o
}

// SetPackage adds the package to the show package manifests params
func (o *ShowPackageManifestsParams) SetPackage(packageVar string) {
	o.Package = packageVar
}

// WithRelease adds the release to the show package manifests params
func (o *ShowPackageManifestsParams) WithRelease(release string) *ShowPackageManifestsParams {
	o.SetRelease(release)
	return o
}

// SetRelease adds the release to the show package manifests params
func (o *ShowPackageManifestsParams) SetRelease(release string) {
	o.Release = release
}

// WriteToRequest writes these params to a swagger request
func (o *ShowPackageManifestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param package
	if err := r.SetPathParam("package", o.Package); err != nil {
		return err
	}

	// path param release
	if err := r.SetPathParam("release", o.Release); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}