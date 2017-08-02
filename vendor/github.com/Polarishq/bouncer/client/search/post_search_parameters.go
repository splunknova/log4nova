package search

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

	"github.com/Polarishq/bouncer/models"
)

// NewPostSearchParams creates a new PostSearchParams object
// with the default values initialized.
func NewPostSearchParams() *PostSearchParams {
	var ()
	return &PostSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSearchParamsWithTimeout creates a new PostSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSearchParamsWithTimeout(timeout time.Duration) *PostSearchParams {
	var ()
	return &PostSearchParams{

		timeout: timeout,
	}
}

// NewPostSearchParamsWithContext creates a new PostSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSearchParamsWithContext(ctx context.Context) *PostSearchParams {
	var ()
	return &PostSearchParams{

		Context: ctx,
	}
}

/*PostSearchParams contains all the parameters to send to the API endpoint
for the post search operation typically these are written to a http.Request
*/
type PostSearchParams struct {

	/*Query*/
	Query *models.QueryInfo
	/*TenantID
	  ID of the tenant

	*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post search params
func (o *PostSearchParams) WithTimeout(timeout time.Duration) *PostSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post search params
func (o *PostSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post search params
func (o *PostSearchParams) WithContext(ctx context.Context) *PostSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post search params
func (o *PostSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithQuery adds the query to the post search params
func (o *PostSearchParams) WithQuery(query *models.QueryInfo) *PostSearchParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the post search params
func (o *PostSearchParams) SetQuery(query *models.QueryInfo) {
	o.Query = query
}

// WithTenantID adds the tenantID to the post search params
func (o *PostSearchParams) WithTenantID(tenantID string) *PostSearchParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the post search params
func (o *PostSearchParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *PostSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Query == nil {
		o.Query = new(models.QueryInfo)
	}

	if err := r.SetBodyParam(o.Query); err != nil {
		return err
	}

	// header param tenantID
	if err := r.SetHeaderParam("tenantID", o.TenantID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
