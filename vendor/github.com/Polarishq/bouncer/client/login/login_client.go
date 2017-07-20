package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new login API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for login API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetLoginAuthorize endpoints to begin the oauth2 authorization flow

Does the initial social auth redirect
*/
func (a *Client) GetLoginAuthorize(params *GetLoginAuthorizeParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoginAuthorizeParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLoginAuthorize",
		Method:             "GET",
		PathPattern:        "/login/authorize",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLoginAuthorizeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetLoginAuthorizeCallback endpoints call via a social login oauth2 callback

Callback for OAuth2 Social logins
*/
func (a *Client) GetLoginAuthorizeCallback(params *GetLoginAuthorizeCallbackParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoginAuthorizeCallbackParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLoginAuthorizeCallback",
		Method:             "GET",
		PathPattern:        "/login/authorize/callback",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLoginAuthorizeCallbackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetLoginProviders returns a list of social providers

Retrieve social providers
*/
func (a *Client) GetLoginProviders(params *GetLoginProvidersParams) (*GetLoginProvidersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoginProvidersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLoginProviders",
		Method:             "GET",
		PathPattern:        "/login/providers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLoginProvidersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLoginProvidersOK), nil

}

/*
PostLoginOauthRevoke revokes a refresh token

Revoke a refresh token
*/
func (a *Client) PostLoginOauthRevoke(params *PostLoginOauthRevokeParams) (*PostLoginOauthRevokeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLoginOauthRevokeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLoginOauthRevoke",
		Method:             "POST",
		PathPattern:        "/login/oauth/revoke",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLoginOauthRevokeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLoginOauthRevokeOK), nil

}

/*
PostLoginOauthToken gets j w t for API keys

Issue a token for an api key
*/
func (a *Client) PostLoginOauthToken(params *PostLoginOauthTokenParams, authInfo runtime.ClientAuthInfoWriter) (*PostLoginOauthTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLoginOauthTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLoginOauthToken",
		Method:             "POST",
		PathPattern:        "/login/oauth/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLoginOauthTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLoginOauthTokenOK), nil

}

/*
PostLoginToken endpoints to get a new refresh token

Gets the token for a given refresh token
*/
func (a *Client) PostLoginToken(params *PostLoginTokenParams) (*PostLoginTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLoginTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLoginToken",
		Method:             "POST",
		PathPattern:        "/login/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLoginTokenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLoginTokenOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
