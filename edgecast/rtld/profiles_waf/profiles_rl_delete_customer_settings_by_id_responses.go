// Code generated by go-swagger; DO NOT EDIT.

package profiles_waf

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ProfilesRlDeleteCustomerSettingsByIDReader is a Reader for the ProfilesRlDeleteCustomerSettingsByID structure.
type ProfilesRlDeleteCustomerSettingsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProfilesRlDeleteCustomerSettingsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewProfilesRlDeleteCustomerSettingsByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProfilesRlDeleteCustomerSettingsByIDNoContent creates a ProfilesRlDeleteCustomerSettingsByIDNoContent with default headers values
func NewProfilesRlDeleteCustomerSettingsByIDNoContent() *ProfilesRlDeleteCustomerSettingsByIDNoContent {
	return &ProfilesRlDeleteCustomerSettingsByIDNoContent{}
}

/* ProfilesRlDeleteCustomerSettingsByIDNoContent describes a response with status code 204, with default header values.

Success
*/
type ProfilesRlDeleteCustomerSettingsByIDNoContent struct {
}

func (o *ProfilesRlDeleteCustomerSettingsByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/rl/profiles/{id}][%d] profilesRlDeleteCustomerSettingsByIdNoContent ", 204)
}

func (o *ProfilesRlDeleteCustomerSettingsByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
