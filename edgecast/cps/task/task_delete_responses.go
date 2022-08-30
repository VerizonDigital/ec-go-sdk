// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package task

import "github.com/EdgeCast/ec-sdk-go/edgecast/cps/models"

// This file was generated by codegen-sdk-go.
// Any changes made to this file may be overwritten.

// NewTaskDeleteNoContent creates a TaskDeleteNoContent with default headers values
func NewTaskDeleteNoContent() *TaskDeleteNoContent {
	return &TaskDeleteNoContent{}
}

/*
	TaskDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type TaskDeleteNoContent struct {
}

// NewTaskDeleteBadRequest creates a TaskDeleteBadRequest with default headers values
func NewTaskDeleteBadRequest() *TaskDeleteBadRequest {
	return &TaskDeleteBadRequest{}
}

/*
	TaskDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TaskDeleteBadRequest struct {
	models.HyperionErrorReponse
}
