// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns detailed metadata about a task that is being executed.
func (c *Client) DescribeTaskExecution(ctx context.Context, params *DescribeTaskExecutionInput, optFns ...func(*Options)) (*DescribeTaskExecutionOutput, error) {
	if params == nil {
		params = &DescribeTaskExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTaskExecution", params, optFns, c.addOperationDescribeTaskExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTaskExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeTaskExecutionRequest
type DescribeTaskExecutionInput struct {

	// The Amazon Resource Name (ARN) of the task that is being executed.
	//
	// This member is required.
	TaskExecutionArn *string

	noSmithyDocumentSerde
}

// DescribeTaskExecutionResponse
type DescribeTaskExecutionOutput struct {

	// The physical number of bytes transferred over the network after compression was
	// applied. In most cases, this number is less than BytesTransferred unless the
	// data isn't compressible.
	BytesCompressed int64

	// The total number of bytes that are involved in the transfer. For the number of
	// bytes sent over the network, see BytesCompressed .
	BytesTransferred int64

	// The number of logical bytes written to the destination Amazon Web Services
	// storage resource.
	BytesWritten int64

	// The estimated physical number of bytes that is to be transferred over the
	// network.
	EstimatedBytesToTransfer int64

	// The expected number of files that is to be transferred over the network. This
	// value is calculated during the PREPARING phase before the TRANSFERRING phase of
	// the task execution. This value is the expected number of files to be
	// transferred. It's calculated based on comparing the content of the source and
	// destination locations and finding the delta that needs to be transferred.
	EstimatedFilesToTransfer int64

	// A list of filter rules that exclude specific data during your transfer. For
	// more information and examples, see Filtering data transferred by DataSync (https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html)
	// .
	Excludes []types.FilterRule

	// The actual number of files that was transferred over the network. This value is
	// calculated and updated on an ongoing basis during the TRANSFERRING phase of the
	// task execution. It's updated periodically when each file is read from the source
	// and sent over the network. If failures occur during a transfer, this value can
	// be less than EstimatedFilesToTransfer . In some cases, this value can also be
	// greater than EstimatedFilesToTransfer . This element is implementation-specific
	// for some location types, so don't use it as an indicator for a correct file
	// number or to monitor your task execution.
	FilesTransferred int64

	// A list of filter rules that include specific data during your transfer. For
	// more information and examples, see Filtering data transferred by DataSync (https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html)
	// .
	Includes []types.FilterRule

	// Configures your DataSync task settings. These options include how DataSync
	// handles files, objects, and their associated metadata. You also can specify how
	// DataSync verifies data integrity, set bandwidth limits for your task, among
	// other options. Each task setting has a default value. Unless you need to, you
	// don't have to configure any of these Options before starting your task.
	Options *types.Options

	// The result of the task execution.
	Result *types.TaskExecutionResultDetail

	// The time that the task execution was started.
	StartTime *time.Time

	// The status of the task execution. For detailed information about task execution
	// statuses, see Understanding Task Statuses in the DataSync User Guide.
	Status types.TaskExecutionStatus

	// The Amazon Resource Name (ARN) of the task execution that was described.
	// TaskExecutionArn is hierarchical and includes TaskArn for the task that was
	// executed. For example, a TaskExecution value with the ARN
	// arn:aws:datasync:us-east-1:111222333444:task/task-0208075f79cedf4a2/execution/exec-08ef1e88ec491019b
	// executed the task with the ARN
	// arn:aws:datasync:us-east-1:111222333444:task/task-0208075f79cedf4a2 .
	TaskExecutionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTaskExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTaskExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTaskExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeTaskExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTaskExecution(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeTaskExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DescribeTaskExecution",
	}
}
