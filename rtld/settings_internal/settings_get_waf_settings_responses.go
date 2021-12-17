// Code generated by go-swagger; DO NOT EDIT.

package settings_internal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/EdgeCast/ec-sdk-go/rtldmodels"
)

// SettingsGetWafSettingsReader is a Reader for the SettingsGetWafSettings structure.
type SettingsGetWafSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SettingsGetWafSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSettingsGetWafSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSettingsGetWafSettingsOK creates a SettingsGetWafSettingsOK with default headers values
func NewSettingsGetWafSettingsOK() *SettingsGetWafSettingsOK {
	return &SettingsGetWafSettingsOK{}
}

/* SettingsGetWafSettingsOK describes a response with status code 200, with default header values.

Success
*/
type SettingsGetWafSettingsOK struct {
	Payload *rtldmodels.HyperionCollectionRtldPlatformSettingDto
}

func (o *SettingsGetWafSettingsOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/waf/settings][%d] settingsGetWafSettingsOK  %+v", 200, o.Payload)
}
func (o *SettingsGetWafSettingsOK) GetPayload() *rtldmodels.HyperionCollectionRtldPlatformSettingDto {
	return o.Payload
}

func (o *SettingsGetWafSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rtldmodels.HyperionCollectionRtldPlatformSettingDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
