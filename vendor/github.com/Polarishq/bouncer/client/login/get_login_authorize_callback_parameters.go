package login

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

// NewGetLoginAuthorizeCallbackParams creates a new GetLoginAuthorizeCallbackParams object
// with the default values initialized.
func NewGetLoginAuthorizeCallbackParams() *GetLoginAuthorizeCallbackParams {
	var ()
	return &GetLoginAuthorizeCallbackParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoginAuthorizeCallbackParamsWithTimeout creates a new GetLoginAuthorizeCallbackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLoginAuthorizeCallbackParamsWithTimeout(timeout time.Duration) *GetLoginAuthorizeCallbackParams {
	var ()
	return &GetLoginAuthorizeCallbackParams{

		timeout: timeout,
	}
}

// NewGetLoginAuthorizeCallbackParamsWithContext creates a new GetLoginAuthorizeCallbackParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLoginAuthorizeCallbackParamsWithContext(ctx context.Context) *GetLoginAuthorizeCallbackParams {
	var ()
	return &GetLoginAuthorizeCallbackParams{

		Context: ctx,
	}
}

/*GetLoginAuthorizeCallbackParams contains all the parameters to send to the API endpoint
for the get login authorize callback operation typically these are written to a http.Request
*/
type GetLoginAuthorizeCallbackParams struct {

	/*Code*/
	Code *string
	/*State*/
	State string
	/*Token*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get login authorize callback params
func (o *GetLoginAuthorizeCallbackParams) WithTimeout(timeout time.Duration) *GetLoginAuthorizeCallbackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get login authorize callback params
func (o *GetLoginAuthorizeCallbackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get login authorize callback params
func (o *GetLoginAuthorizeCallbackParams) WithContext(ctx context.Context) *GetLoginAuthorizeCallbackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get login authorize callback params
func (o *GetLoginAuthorizeCallbackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithCode adds the code to the get login authorize callback params
func (o *GetLoginAuthorizeCallbackParams) WithCode(code *string) *GetLoginAuthorizeCallbackParams {
	o.SetCode(code)
	return o
}

// SetCode adds the code to the get login authorize callback params
func (o *GetLoginAuthorizeCallbackParams) SetCode(code *string) {
	o.Code = code
}

// WithState adds the state to the get login authorize callback params
func (o *GetLoginAuthorizeCallbackParams) WithState(state string) *GetLoginAuthorizeCallbackParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the get login authorize callback params
func (o *GetLoginAuthorizeCallbackParams) SetState(state string) {
	o.State = state
}

// WithToken adds the token to the get login authorize callback params
func (o *GetLoginAuthorizeCallbackParams) WithToken(token *string) *GetLoginAuthorizeCallbackParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get login authorize callback params
func (o *GetLoginAuthorizeCallbackParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoginAuthorizeCallbackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Code != nil {

		// query param code
		var qrCode string
		if o.Code != nil {
			qrCode = *o.Code
		}
		qCode := qrCode
		if qCode != "" {
			if err := r.SetQueryParam("code", qCode); err != nil {
				return err
			}
		}

	}

	// query param state
	qrState := o.State
	qState := qrState
	if qState != "" {
		if err := r.SetQueryParam("state", qState); err != nil {
			return err
		}
	}

	if o.Token != nil {

		// query param token
		var qrToken string
		if o.Token != nil {
			qrToken = *o.Token
		}
		qToken := qrToken
		if qToken != "" {
			if err := r.SetQueryParam("token", qToken); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}