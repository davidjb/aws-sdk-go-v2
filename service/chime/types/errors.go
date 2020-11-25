// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
)

// You don't have permissions to perform the requested operation.
type AccessDeniedException struct {
	Message *string

	Code ErrorCode
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string             { return "AccessDeniedException" }
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The input parameters don't match the service's restrictions.
type BadRequestException struct {
	Message *string

	Code ErrorCode
}

func (e *BadRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadRequestException) ErrorCode() string             { return "BadRequestException" }
func (e *BadRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request could not be processed because of conflict in the current state of
// the resource.
type ConflictException struct {
	Message *string

	Code ErrorCode
}

func (e *ConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictException) ErrorCode() string             { return "ConflictException" }
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The client is permanently forbidden from making the request.
type ForbiddenException struct {
	Message *string

	Code ErrorCode
}

func (e *ForbiddenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ForbiddenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ForbiddenException) ErrorCode() string             { return "ForbiddenException" }
func (e *ForbiddenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more of the resources in the request does not exist in the system.
type NotFoundException struct {
	Message *string

	Code ErrorCode
}

func (e *NotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotFoundException) ErrorCode() string             { return "NotFoundException" }
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request exceeds the resource limit.
type ResourceLimitExceededException struct {
	Message *string

	Code ErrorCode
}

func (e *ResourceLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceLimitExceededException) ErrorCode() string             { return "ResourceLimitExceededException" }
func (e *ResourceLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service encountered an unexpected error.
type ServiceFailureException struct {
	Message *string

	Code ErrorCode
}

func (e *ServiceFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceFailureException) ErrorCode() string             { return "ServiceFailureException" }
func (e *ServiceFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The service is currently unavailable.
type ServiceUnavailableException struct {
	Message *string

	Code ErrorCode
}

func (e *ServiceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailableException) ErrorCode() string             { return "ServiceUnavailableException" }
func (e *ServiceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The client exceeded its request rate limit.
type ThrottledClientException struct {
	Message *string

	Code ErrorCode
}

func (e *ThrottledClientException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottledClientException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottledClientException) ErrorCode() string             { return "ThrottledClientException" }
func (e *ThrottledClientException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The client is not currently authorized to make the request.
type UnauthorizedClientException struct {
	Message *string

	Code ErrorCode
}

func (e *UnauthorizedClientException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnauthorizedClientException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnauthorizedClientException) ErrorCode() string             { return "UnauthorizedClientException" }
func (e *UnauthorizedClientException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was well-formed but was unable to be followed due to semantic
// errors.
type UnprocessableEntityException struct {
	Message *string

	Code ErrorCode
}

func (e *UnprocessableEntityException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnprocessableEntityException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnprocessableEntityException) ErrorCode() string             { return "UnprocessableEntityException" }
func (e *UnprocessableEntityException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
