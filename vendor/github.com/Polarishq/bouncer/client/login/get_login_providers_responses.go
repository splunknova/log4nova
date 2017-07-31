package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Polarishq/bouncer/models"
)

// GetLoginProvidersReader is a Reader for the GetLoginProviders structure.
type GetLoginProvidersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoginProvidersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLoginProvidersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetLoginProvidersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLoginProvidersOK creates a GetLoginProvidersOK with default headers values
func NewGetLoginProvidersOK() *GetLoginProvidersOK {
	return &GetLoginProvidersOK{}
}

/*GetLoginProvidersOK handles this case with default header values.

Account retrieved
*/
type GetLoginProvidersOK struct {
	Payload []*models.SocialProvider
}

func (o *GetLoginProvidersOK) Error() string {
	return fmt.Sprintf("[GET /login/providers][%d] getLoginProvidersOK  %+v", 200, o.Payload)
}

func (o *GetLoginProvidersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoginProvidersDefault creates a GetLoginProvidersDefault with default headers values
func NewGetLoginProvidersDefault(code int) *GetLoginProvidersDefault {
	return &GetLoginProvidersDefault{
		_statusCode: code,
	}
}

/*GetLoginProvidersDefault handles this case with default header values.

Unexpected error
*/
type GetLoginProvidersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get login providers default response
func (o *GetLoginProvidersDefault) Code() int {
	return o._statusCode
}

func (o *GetLoginProvidersDefault) Error() string {
	return fmt.Sprintf("[GET /login/providers][%d] GetLoginProviders default  %+v", o._statusCode, o.Payload)
}

func (o *GetLoginProvidersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
