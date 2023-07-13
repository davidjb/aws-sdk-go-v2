// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a batch prediction job.
func (c *Client) CreateBatchPredictionJob(ctx context.Context, params *CreateBatchPredictionJobInput, optFns ...func(*Options)) (*CreateBatchPredictionJobOutput, error) {
	if params == nil {
		params = &CreateBatchPredictionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBatchPredictionJob", params, optFns, c.addOperationCreateBatchPredictionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBatchPredictionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBatchPredictionJobInput struct {

	// The name of the detector.
	//
	// This member is required.
	DetectorName *string

	// The name of the event type.
	//
	// This member is required.
	EventTypeName *string

	// The ARN of the IAM role to use for this job request. The IAM Role must have
	// read permissions to your input S3 bucket and write permissions to your output S3
	// bucket. For more information about bucket permissions, see User policy examples (https://docs.aws.amazon.com/AmazonS3/latest/userguide/example-policies-s3.html)
	// in the Amazon S3 User Guide.
	//
	// This member is required.
	IamRoleArn *string

	// The Amazon S3 location of your training file.
	//
	// This member is required.
	InputPath *string

	// The ID of the batch prediction job.
	//
	// This member is required.
	JobId *string

	// The Amazon S3 location of your output file.
	//
	// This member is required.
	OutputPath *string

	// The detector version.
	DetectorVersion *string

	// A collection of key and value pairs.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateBatchPredictionJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBatchPredictionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateBatchPredictionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateBatchPredictionJob{}, middleware.After)
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
	if err = addOpCreateBatchPredictionJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBatchPredictionJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBatchPredictionJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "CreateBatchPredictionJob",
	}
}
