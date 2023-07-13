// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutvision

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutvision/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an Amazon Lookout for Vision model packaging job. A model packaging job
// creates an AWS IoT Greengrass component for a Lookout for Vision model. You can
// use the component to deploy your model to an edge device managed by Greengrass.
// Use the DescribeModelPackagingJob API to determine the current status of the
// job. The model packaging job is complete if the value of Status is SUCCEEDED .
// To deploy the component to the target device, use the component name and
// component version with the AWS IoT Greengrass CreateDeployment (https://docs.aws.amazon.com/greengrass/v2/APIReference/API_CreateDeployment.html)
// API. This operation requires the following permissions:
//   - lookoutvision:StartModelPackagingJob
//   - s3:PutObject
//   - s3:GetBucketLocation
//   - kms:GenerateDataKey
//   - greengrass:CreateComponentVersion
//   - greengrass:DescribeComponent
//   - (Optional) greengrass:TagResource . Only required if you want to tag the
//     component.
//
// For more information, see Using your Amazon Lookout for Vision model on an edge
// device in the Amazon Lookout for Vision Developer Guide.
func (c *Client) StartModelPackagingJob(ctx context.Context, params *StartModelPackagingJobInput, optFns ...func(*Options)) (*StartModelPackagingJobOutput, error) {
	if params == nil {
		params = &StartModelPackagingJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartModelPackagingJob", params, optFns, c.addOperationStartModelPackagingJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartModelPackagingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartModelPackagingJobInput struct {

	// The configuration for the model packaging job.
	//
	// This member is required.
	Configuration *types.ModelPackagingConfiguration

	// The version of the model within the project that you want to package.
	//
	// This member is required.
	ModelVersion *string

	// The name of the project which contains the version of the model that you want
	// to package.
	//
	// This member is required.
	ProjectName *string

	// ClientToken is an idempotency token that ensures a call to
	// StartModelPackagingJob completes only once. You choose the value to pass. For
	// example, An issue might prevent you from getting a response from
	// StartModelPackagingJob . In this case, safely retry your call to
	// StartModelPackagingJob by using the same ClientToken parameter value. If you
	// don't supply a value for ClientToken , the AWS SDK you are using inserts a value
	// for you. This prevents retries after a network error from making multiple
	// dataset creation requests. You'll need to provide your own value for other use
	// cases. An error occurs if the other input parameters are not the same as in the
	// first request. Using a different value for ClientToken is considered a new call
	// to StartModelPackagingJob . An idempotency token is active for 8 hours.
	ClientToken *string

	// A description for the model packaging job.
	Description *string

	// A name for the model packaging job. If you don't supply a value, the service
	// creates a job name for you.
	JobName *string

	noSmithyDocumentSerde
}

type StartModelPackagingJobOutput struct {

	// The job name for the model packaging job. If you don't supply a job name in the
	// JobName input parameter, the service creates a job name for you.
	JobName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartModelPackagingJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartModelPackagingJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartModelPackagingJob{}, middleware.After)
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
	if err = addIdempotencyToken_opStartModelPackagingJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartModelPackagingJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartModelPackagingJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartModelPackagingJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartModelPackagingJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartModelPackagingJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartModelPackagingJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartModelPackagingJobInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartModelPackagingJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartModelPackagingJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartModelPackagingJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutvision",
		OperationName: "StartModelPackagingJob",
	}
}
