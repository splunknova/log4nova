package applications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetApplicationsParams creates a new GetApplicationsParams object
// with the default values initialized.
func NewGetApplicationsParams() *GetApplicationsParams {
	var (
		countDefault = int32(10)
		indexDefault = int32(0)
	)
	return &GetApplicationsParams{
		Count: &countDefault,
		Index: &indexDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetApplicationsParamsWithTimeout creates a new GetApplicationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetApplicationsParamsWithTimeout(timeout time.Duration) *GetApplicationsParams {
	var (
		countDefault = int32(10)
		indexDefault = int32(0)
	)
	return &GetApplicationsParams{
		Count: &countDefault,
		Index: &indexDefault,

		timeout: timeout,
	}
}

// NewGetApplicationsParamsWithContext creates a new GetApplicationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetApplicationsParamsWithContext(ctx context.Context) *GetApplicationsParams {
	var (
		countDefault = int32(10)
		indexDefault = int32(0)
	)
	return &GetApplicationsParams{
		Count: &countDefault,
		Index: &indexDefault,

		Context: ctx,
	}
}

/*GetApplicationsParams contains all the parameters to send to the API endpoint
for the get applications operation typically these are written to a http.Request
*/
type GetApplicationsParams struct {

	/*Count
	  Number of events desired.  Will return fewer if there are fewer to return.

	*/
	Count *int32
	/*Index
	  Index of the first event in the set of returned results.  Negative index will count from the end of the results.

	*/
	Index *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get applications params
func (o *GetApplicationsParams) WithTimeout(timeout time.Duration) *GetApplicationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get applications params
func (o *GetApplicationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get applications params
func (o *GetApplicationsParams) WithContext(ctx context.Context) *GetApplicationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get applications params
func (o *GetApplicationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCount adds the count to the get applications params
func (o *GetApplicationsParams) WithCount(count *int32) *GetApplicationsParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get applications params
func (o *GetApplicationsParams) SetCount(count *int32) {
	o.Count = count
}

// WithIndex adds the index to the get applications params
func (o *GetApplicationsParams) WithIndex(index *int32) *GetApplicationsParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the get applications params
func (o *GetApplicationsParams) SetIndex(index *int32) {
	o.Index = index
}

// WriteToRequest writes these params to a swagger request
func (o *GetApplicationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int32
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt32(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	if o.Index != nil {

		// query param index
		var qrIndex int32
		if o.Index != nil {
			qrIndex = *o.Index
		}
		qIndex := swag.FormatInt32(qrIndex)
		if qIndex != "" {
			if err := r.SetQueryParam("index", qIndex); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
