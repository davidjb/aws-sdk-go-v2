// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
)

// There is another ongoing conflicting backup control plane operation on the
// table. The backup is either being created, deleted or restored to a table.
type BackupInUseException struct {
	Message *string
}

func (e *BackupInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BackupInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BackupInUseException) ErrorCode() string             { return "BackupInUseException" }
func (e *BackupInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Backup not found for the given BackupARN.
type BackupNotFoundException struct {
	Message *string
}

func (e *BackupNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BackupNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BackupNotFoundException) ErrorCode() string             { return "BackupNotFoundException" }
func (e *BackupNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A condition specified in the operation could not be evaluated.
type ConditionalCheckFailedException struct {
	Message *string
}

func (e *ConditionalCheckFailedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConditionalCheckFailedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConditionalCheckFailedException) ErrorCode() string {
	return "ConditionalCheckFailedException"
}
func (e *ConditionalCheckFailedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Backups have not yet been enabled for this table.
type ContinuousBackupsUnavailableException struct {
	Message *string
}

func (e *ContinuousBackupsUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ContinuousBackupsUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ContinuousBackupsUnavailableException) ErrorCode() string {
	return "ContinuousBackupsUnavailableException"
}
func (e *ContinuousBackupsUnavailableException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// There was a conflict when writing to the specified S3 bucket.
type ExportConflictException struct {
	Message *string
}

func (e *ExportConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ExportConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ExportConflictException) ErrorCode() string             { return "ExportConflictException" }
func (e *ExportConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified export was not found.
type ExportNotFoundException struct {
	Message *string
}

func (e *ExportNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ExportNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ExportNotFoundException) ErrorCode() string             { return "ExportNotFoundException" }
func (e *ExportNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified global table already exists.
type GlobalTableAlreadyExistsException struct {
	Message *string
}

func (e *GlobalTableAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *GlobalTableAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *GlobalTableAlreadyExistsException) ErrorCode() string {
	return "GlobalTableAlreadyExistsException"
}
func (e *GlobalTableAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified global table does not exist.
type GlobalTableNotFoundException struct {
	Message *string
}

func (e *GlobalTableNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *GlobalTableNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *GlobalTableNotFoundException) ErrorCode() string             { return "GlobalTableNotFoundException" }
func (e *GlobalTableNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// DynamoDB rejected the request because you retried a request with a different
// payload but with an idempotent token that was already used.
type IdempotentParameterMismatchException struct {
	Message *string
}

func (e *IdempotentParameterMismatchException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IdempotentParameterMismatchException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IdempotentParameterMismatchException) ErrorCode() string {
	return "IdempotentParameterMismatchException"
}
func (e *IdempotentParameterMismatchException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The operation tried to access a nonexistent index.
type IndexNotFoundException struct {
	Message *string
}

func (e *IndexNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IndexNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IndexNotFoundException) ErrorCode() string             { return "IndexNotFoundException" }
func (e *IndexNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An error occurred on the server side.
type InternalServerError struct {
	Message *string
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerError) ErrorCode() string             { return "InternalServerError" }
func (e *InternalServerError) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

type InvalidEndpointException struct {
	Message *string
}

func (e *InvalidEndpointException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidEndpointException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidEndpointException) ErrorCode() string             { return "InvalidEndpointException" }
func (e *InvalidEndpointException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified ExportTime is outside of the point in time recovery window.
type InvalidExportTimeException struct {
	Message *string
}

func (e *InvalidExportTimeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidExportTimeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidExportTimeException) ErrorCode() string             { return "InvalidExportTimeException" }
func (e *InvalidExportTimeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An invalid restore time was specified. RestoreDateTime must be between
// EarliestRestorableDateTime and LatestRestorableDateTime.
type InvalidRestoreTimeException struct {
	Message *string
}

func (e *InvalidRestoreTimeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRestoreTimeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRestoreTimeException) ErrorCode() string             { return "InvalidRestoreTimeException" }
func (e *InvalidRestoreTimeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An item collection is too large. This exception is only returned for tables that
// have one or more local secondary indexes.
type ItemCollectionSizeLimitExceededException struct {
	Message *string
}

func (e *ItemCollectionSizeLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ItemCollectionSizeLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ItemCollectionSizeLimitExceededException) ErrorCode() string {
	return "ItemCollectionSizeLimitExceededException"
}
func (e *ItemCollectionSizeLimitExceededException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// There is no limit to the number of daily on-demand backups that can be taken. Up
// to 50 simultaneous table operations are allowed per account. These operations
// include CreateTable, UpdateTable, DeleteTable,UpdateTimeToLive,
// RestoreTableFromBackup, and RestoreTableToPointInTime. The only exception is
// when you are creating a table with one or more secondary indexes. You can have
// up to 25 such requests running at a time; however, if the table or index
// specifications are complex, DynamoDB might temporarily reduce the number of
// concurrent operations. There is a soft account quota of 256 tables.
type LimitExceededException struct {
	Message *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Point in time recovery has not yet been enabled for this source table.
type PointInTimeRecoveryUnavailableException struct {
	Message *string
}

func (e *PointInTimeRecoveryUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PointInTimeRecoveryUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PointInTimeRecoveryUnavailableException) ErrorCode() string {
	return "PointInTimeRecoveryUnavailableException"
}
func (e *PointInTimeRecoveryUnavailableException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Your request rate is too high. The AWS SDKs for DynamoDB automatically retry
// requests that receive this exception. Your request is eventually successful,
// unless your retry queue is too large to finish. Reduce the frequency of requests
// and use exponential backoff. For more information, go to Error Retries and
// Exponential Backoff
// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Programming.Errors.html#Programming.Errors.RetryAndBackoff)
// in the Amazon DynamoDB Developer Guide.
type ProvisionedThroughputExceededException struct {
	Message *string
}

func (e *ProvisionedThroughputExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ProvisionedThroughputExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ProvisionedThroughputExceededException) ErrorCode() string {
	return "ProvisionedThroughputExceededException"
}
func (e *ProvisionedThroughputExceededException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified replica is already part of the global table.
type ReplicaAlreadyExistsException struct {
	Message *string
}

func (e *ReplicaAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReplicaAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReplicaAlreadyExistsException) ErrorCode() string             { return "ReplicaAlreadyExistsException" }
func (e *ReplicaAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified replica is no longer part of the global table.
type ReplicaNotFoundException struct {
	Message *string
}

func (e *ReplicaNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReplicaNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReplicaNotFoundException) ErrorCode() string             { return "ReplicaNotFoundException" }
func (e *ReplicaNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Throughput exceeds the current throughput quota for your account. Please contact
// AWS Support at AWS Support (https://aws.amazon.com/support) to request a quota
// increase.
type RequestLimitExceeded struct {
	Message *string
}

func (e *RequestLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RequestLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RequestLimitExceeded) ErrorCode() string             { return "RequestLimitExceeded" }
func (e *RequestLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation conflicts with the resource's availability. For example, you
// attempted to recreate an existing table, or tried to delete a table currently in
// the CREATING state.
type ResourceInUseException struct {
	Message *string
}

func (e *ResourceInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceInUseException) ErrorCode() string             { return "ResourceInUseException" }
func (e *ResourceInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation tried to access a nonexistent table or index. The resource might
// not be specified correctly, or its status might not be ACTIVE.
type ResourceNotFoundException struct {
	Message *string
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A target table with the specified name already exists.
type TableAlreadyExistsException struct {
	Message *string
}

func (e *TableAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TableAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TableAlreadyExistsException) ErrorCode() string             { return "TableAlreadyExistsException" }
func (e *TableAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A target table with the specified name is either being created or deleted.
type TableInUseException struct {
	Message *string
}

func (e *TableInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TableInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TableInUseException) ErrorCode() string             { return "TableInUseException" }
func (e *TableInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A source table with the name TableName does not currently exist within the
// subscriber's account.
type TableNotFoundException struct {
	Message *string
}

func (e *TableNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TableNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TableNotFoundException) ErrorCode() string             { return "TableNotFoundException" }
func (e *TableNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The entire transaction request was canceled. DynamoDB cancels a
// TransactWriteItems request under the following circumstances:
//
// * A condition in
// one of the condition expressions is not met.
//
// * A table in the
// TransactWriteItems request is in a different account or region.
//
// * More than one
// action in the TransactWriteItems operation targets the same item.
//
// * There is
// insufficient provisioned capacity for the transaction to be completed.
//
// * An
// item size becomes too large (larger than 400 KB), or a local secondary index
// (LSI) becomes too large, or a similar validation error occurs because of changes
// made by the transaction.
//
// * There is a user error, such as an invalid data
// format.
//
// DynamoDB cancels a TransactGetItems request under the following
// circumstances:
//
// * There is an ongoing TransactGetItems operation that conflicts
// with a concurrent PutItem, UpdateItem, DeleteItem or TransactWriteItems request.
// In this case the TransactGetItems operation fails with a
// TransactionCanceledException.
//
// * A table in the TransactGetItems request is in a
// different account or region.
//
// * There is insufficient provisioned capacity for
// the transaction to be completed.
//
// * There is a user error, such as an invalid
// data format.
//
// If using Java, DynamoDB lists the cancellation reasons on the
// CancellationReasons property. This property is not set for other languages.
// Transaction cancellation reasons are ordered in the order of requested items, if
// an item has no error it will have NONE code and Null message. Cancellation
// reason codes and possible error messages:
//
// * No Errors:
//
// * Code: NONE
//
// *
// Message: null
//
// * Conditional Check Failed:
//
// * Code: ConditionalCheckFailed
//
// *
// Message: The conditional request failed.
//
// * Item Collection Size Limit
// Exceeded:
//
// * Code: ItemCollectionSizeLimitExceeded
//
// * Message: Collection size
// exceeded.
//
// * Transaction Conflict:
//
// * Code: TransactionConflict
//
// * Message:
// Transaction is ongoing for the item.
//
// * Provisioned Throughput Exceeded:
//
// *
// Code: ProvisionedThroughputExceeded
//
// * Messages:
//
// * The level of configured
// provisioned throughput for the table was exceeded. Consider increasing your
// provisioning level with the UpdateTable API. This Message is received when
// provisioned throughput is exceeded is on a provisioned DynamoDB table.
//
// * The
// level of configured provisioned throughput for one or more global secondary
// indexes of the table was exceeded. Consider increasing your provisioning level
// for the under-provisioned global secondary indexes with the UpdateTable API.
// This message is returned when provisioned throughput is exceeded is on a
// provisioned GSI.
//
// * Throttling Error:
//
// * Code: ThrottlingError
//
// * Messages:
//
// *
// Throughput exceeds the current capacity of your table or index. DynamoDB is
// automatically scaling your table or index so please try again shortly. If
// exceptions persist, check if you have a hot key:
// https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-partition-key-design.html.
// This message is returned when writes get throttled on an On-Demand table as
// DynamoDB is automatically scaling the table.
//
// * Throughput exceeds the current
// capacity for one or more global secondary indexes. DynamoDB is automatically
// scaling your index so please try again shortly. This message is returned when
// when writes get throttled on an On-Demand GSI as DynamoDB is automatically
// scaling the GSI.
//
// * Validation Error:
//
// * Code: ValidationError
//
// * Messages:
//
// *
// One or more parameter values were invalid.
//
// * The update expression attempted to
// update the secondary index key beyond allowed size limits.
//
// * The update
// expression attempted to update the secondary index key to unsupported type.
//
// *
// An operand in the update expression has an incorrect data type.
//
// * Item size to
// update has exceeded the maximum allowed size.
//
// * Number overflow. Attempting to
// store a number with magnitude larger than supported range.
//
// * Type mismatch for
// attribute to update.
//
// * Nesting Levels have exceeded supported limits.
//
// * The
// document path provided in the update expression is invalid for update.
//
// * The
// provided expression refers to an attribute that does not exist in the item.
type TransactionCanceledException struct {
	Message *string

	CancellationReasons []CancellationReason
}

func (e *TransactionCanceledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TransactionCanceledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TransactionCanceledException) ErrorCode() string             { return "TransactionCanceledException" }
func (e *TransactionCanceledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Operation was rejected because there is an ongoing transaction for the item.
type TransactionConflictException struct {
	Message *string
}

func (e *TransactionConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TransactionConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TransactionConflictException) ErrorCode() string             { return "TransactionConflictException" }
func (e *TransactionConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The transaction with the given request token is already in progress.
type TransactionInProgressException struct {
	Message *string
}

func (e *TransactionInProgressException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TransactionInProgressException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TransactionInProgressException) ErrorCode() string             { return "TransactionInProgressException" }
func (e *TransactionInProgressException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
