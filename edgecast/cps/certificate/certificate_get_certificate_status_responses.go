// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package certificate

import "github.com/EdgeCast/ec-sdk-go/edgecast/cps/models"

// This file was generated by codegen-sdk-go.
// Any changes made to this file may be overwritten.

// NewCertificateGetCertificateStatusOK creates a CertificateGetCertificateStatusOK with default headers values
func NewCertificateGetCertificateStatusOK() *CertificateGetCertificateStatusOK {
	return &CertificateGetCertificateStatusOK{}
}

/*
	CertificateGetCertificateStatusOK describes a response with status code 200, with default header values.

Success
*/
type CertificateGetCertificateStatusOK struct {
	models.CertificateStatus
}

// NewCertificateGetCertificateStatusBadRequest creates a CertificateGetCertificateStatusBadRequest with default headers values
func NewCertificateGetCertificateStatusBadRequest() *CertificateGetCertificateStatusBadRequest {
	return &CertificateGetCertificateStatusBadRequest{}
}

/*
	CertificateGetCertificateStatusBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CertificateGetCertificateStatusBadRequest struct {
	models.HyperionErrorReponse
}

// NewCertificateGetCertificateStatusNotFound creates a CertificateGetCertificateStatusNotFound with default headers values
func NewCertificateGetCertificateStatusNotFound() *CertificateGetCertificateStatusNotFound {
	return &CertificateGetCertificateStatusNotFound{}
}

/*
	CertificateGetCertificateStatusNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CertificateGetCertificateStatusNotFound struct {
	models.HyperionErrorReponse
}
