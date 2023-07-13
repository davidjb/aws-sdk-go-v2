// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about a stack drift detection operation. A stack drift
// detection operation detects whether a stack's actual configuration differs, or
// has drifted, from its expected configuration, as defined in the stack template
// and any values specified as template parameters. A stack is considered to have
// drifted if one or more of its resources have drifted. For more information about
// stack and resource drift, see Detecting Unregulated Configuration Changes to
// Stacks and Resources (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift.html)
// . Use DetectStackDrift to initiate a stack drift detection operation.
// DetectStackDrift returns a StackDriftDetectionId you can use to monitor the
// progress of the operation using DescribeStackDriftDetectionStatus . Once the
// drift detection operation has completed, use DescribeStackResourceDrifts to
// return drift information about the stack and its resources.
func (c *Client) DescribeStackDriftDetectionStatus(ctx context.Context, params *DescribeStackDriftDetectionStatusInput, optFns ...func(*Options)) (*DescribeStackDriftDetectionStatusOutput, error) {
	if params == nil {
		params = &DescribeStackDriftDetectionStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStackDriftDetectionStatus", params, optFns, c.addOperationDescribeStackDriftDetectionStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStackDriftDetectionStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStackDriftDetectionStatusInput struct {

	// The ID of the drift detection results of this operation. CloudFormation
	// generates new results, with a new drift detection ID, each time this operation
	// is run. However, the number of drift results CloudFormation retains for any
	// given stack, and for how long, may vary.
	//
	// This member is required.
	StackDriftDetectionId *string

	noSmithyDocumentSerde
}

type DescribeStackDriftDetectionStatusOutput struct {

	// The status of the stack drift detection operation.
	//   - DETECTION_COMPLETE : The stack drift detection operation has successfully
	//   completed for all resources in the stack that support drift detection.
	//   (Resources that don't currently support stack detection remain unchecked.) If
	//   you specified logical resource IDs for CloudFormation to use as a filter for the
	//   stack drift detection operation, only the resources with those logical IDs are
	//   checked for drift.
	//   - DETECTION_FAILED : The stack drift detection operation has failed for at
	//   least one resource in the stack. Results will be available for resources on
	//   which CloudFormation successfully completed drift detection.
	//   - DETECTION_IN_PROGRESS : The stack drift detection operation is currently in
	//   progress.
	//
	// This member is required.
	DetectionStatus types.StackDriftDetectionStatus

	// The ID of the drift detection results of this operation. CloudFormation
	// generates new results, with a new drift detection ID, each time this operation
	// is run. However, the number of reports CloudFormation retains for any given
	// stack, and for how long, may vary.
	//
	// This member is required.
	StackDriftDetectionId *string

	// The ID of the stack.
	//
	// This member is required.
	StackId *string

	// Time at which the stack drift detection operation was initiated.
	//
	// This member is required.
	Timestamp *time.Time

	// The reason the stack drift detection operation has its current status.
	DetectionStatusReason *string

	// Total number of stack resources that have drifted. This is NULL until the drift
	// detection operation reaches a status of DETECTION_COMPLETE . This value will be
	// 0 for stacks whose drift status is IN_SYNC .
	DriftedStackResourceCount *int32

	// Status of the stack's actual configuration compared to its expected
	// configuration.
	//   - DRIFTED : The stack differs from its expected template configuration. A
	//   stack is considered to have drifted if one or more of its resources have
	//   drifted.
	//   - NOT_CHECKED : CloudFormation hasn't checked if the stack differs from its
	//   expected template configuration.
	//   - IN_SYNC : The stack's actual configuration matches its expected template
	//   configuration.
	//   - UNKNOWN : This value is reserved for future use.
	StackDriftStatus types.StackDriftStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStackDriftDetectionStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeStackDriftDetectionStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeStackDriftDetectionStatus{}, middleware.After)
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
	if err = addOpDescribeStackDriftDetectionStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStackDriftDetectionStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeStackDriftDetectionStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DescribeStackDriftDetectionStatus",
	}
}
