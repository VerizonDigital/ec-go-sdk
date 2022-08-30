// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package customer

import "github.com/EdgeCast/ec-sdk-go/edgecast/cps/models"

// This file was generated by codegen-sdk-go.
// Any changes made to this file may be overwritten.

// NewCustomerGetCustomerCommitsOK creates a CustomerGetCustomerCommitsOK with default headers values
func NewCustomerGetCustomerCommitsOK() *CustomerGetCustomerCommitsOK {
	return &CustomerGetCustomerCommitsOK{}
}

/*
	CustomerGetCustomerCommitsOK describes a response with status code 200, with default header values.

Success
*/
type CustomerGetCustomerCommitsOK struct {
	models.CustomerCommit
}

// NewCustomerGetCustomerCommitsNotFound creates a CustomerGetCustomerCommitsNotFound with default headers values
func NewCustomerGetCustomerCommitsNotFound() *CustomerGetCustomerCommitsNotFound {
	return &CustomerGetCustomerCommitsNotFound{}
}

/*
	CustomerGetCustomerCommitsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CustomerGetCustomerCommitsNotFound struct {
	models.HyperionErrorReponse
}
