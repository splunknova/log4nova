package identity

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

// NewPostAccountAPIKeysParams creates a new PostAccountAPIKeysParams object
// with the default values initialized.
func NewPostAccountAPIKeysParams() *PostAccountAPIKeysParams {
	var ()
	return &PostAccountAPIKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAccountAPIKeysParamsWithTimeout creates a new PostAccountAPIKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAccountAPIKeysParamsWithTimeout(timeout time.Duration) *PostAccountAPIKeysParams {
	var ()
	return &PostAccountAPIKeysParams{

		timeout: timeout,
	}
}

// NewPostAccountAPIKeysParamsWithContext creates a new PostAccountAPIKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAccountAPIKeysParamsWithContext(ctx context.Context) *PostAccountAPIKeysParams {
	var ()
	return &PostAccountAPIKeysParams{

		Context: ctx,
	}
}

/*PostAccountAPIKeysParams contains all the parameters to send to the API endpoint
for the post account API keys operation typically these are written to a http.Request
*/
type PostAccountAPIKeysParams struct {

	/*Body*/
	Body PostAccountAPIKeysBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post account API keys params
func (o *PostAccountAPIKeysParams) WithTimeout(timeout time.Duration) *PostAccountAPIKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post account API keys params
func (o *PostAccountAPIKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post account API keys params
func (o *PostAccountAPIKeysParams) WithContext(ctx context.Context) *PostAccountAPIKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post account API keys params
func (o *PostAccountAPIKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post account API keys params
func (o *PostAccountAPIKeysParams) WithBody(body PostAccountAPIKeysBody) *PostAccountAPIKeysParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post account API keys params
func (o *PostAccountAPIKeysParams) SetBody(body PostAccountAPIKeysBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostAccountAPIKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
