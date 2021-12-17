// Code generated by go-swagger; DO NOT EDIT.

package lookups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/EdgeCast/ec-sdk-go/rtldmodels"
)

// LookupsGetAwsRegionsReader is a Reader for the LookupsGetAwsRegions structure.
type LookupsGetAwsRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LookupsGetAwsRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLookupsGetAwsRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLookupsGetAwsRegionsOK creates a LookupsGetAwsRegionsOK with default headers values
func NewLookupsGetAwsRegionsOK() *LookupsGetAwsRegionsOK {
	return &LookupsGetAwsRegionsOK{}
}

/* LookupsGetAwsRegionsOK describes a response with status code 200, with default header values.

Success
*/
type LookupsGetAwsRegionsOK struct {
	Payload *rtldmodels.HyperionCollectionCodeName
}

func (o *LookupsGetAwsRegionsOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/aws-regions][%d] lookupsGetAwsRegionsOK  %+v", 200, o.Payload)
}
func (o *LookupsGetAwsRegionsOK) GetPayload() *rtldmodels.HyperionCollectionCodeName {
	return o.Payload
}

func (o *LookupsGetAwsRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rtldmodels.HyperionCollectionCodeName)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
