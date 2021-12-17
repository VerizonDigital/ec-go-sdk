// Code generated by go-swagger; DO NOT EDIT.

package lookups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/EdgeCast/ec-sdk-go/edgecast/rtldmodels"
)

// LookupsGetFieldRlReader is a Reader for the LookupsGetFieldRl structure.
type LookupsGetFieldRlReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LookupsGetFieldRlReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLookupsGetFieldRlOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLookupsGetFieldRlOK creates a LookupsGetFieldRlOK with default headers values
func NewLookupsGetFieldRlOK() *LookupsGetFieldRlOK {
	return &LookupsGetFieldRlOK{}
}

/* LookupsGetFieldRlOK describes a response with status code 200, with default header values.

Success
*/
type LookupsGetFieldRlOK struct {
	Payload *rtldmodels.HyperionCollectionLogFieldDto
}

func (o *LookupsGetFieldRlOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/rl/fields][%d] lookupsGetFieldRlOK  %+v", 200, o.Payload)
}
func (o *LookupsGetFieldRlOK) GetPayload() *rtldmodels.HyperionCollectionLogFieldDto {
	return o.Payload
}

func (o *LookupsGetFieldRlOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rtldmodels.HyperionCollectionLogFieldDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
