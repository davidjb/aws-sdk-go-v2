// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// A description of edge packaging jobs.
func (c *Client) DescribeEdgePackagingJob(ctx context.Context, params *DescribeEdgePackagingJobInput, optFns ...func(*Options)) (*DescribeEdgePackagingJobOutput, error) {
	if params == nil {
		params = &DescribeEdgePackagingJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEdgePackagingJob", params, optFns, c.addOperationDescribeEdgePackagingJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEdgePackagingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEdgePackagingJobInput struct {

	// The name of the edge packaging job.
	//
	// This member is required.
	EdgePackagingJobName *string

	noSmithyDocumentSerde
}

type DescribeEdgePackagingJobOutput struct {

	// The Amazon Resource Name (ARN) of the edge packaging job.
	//
	// This member is required.
	EdgePackagingJobArn *string

	// The name of the edge packaging job.
	//
	// This member is required.
	EdgePackagingJobName *string

	// The current status of the packaging job.
	//
	// This member is required.
	EdgePackagingJobStatus types.EdgePackagingJobStatus

	// The name of the SageMaker Neo compilation job that is used to locate model
	// artifacts that are being packaged.
	CompilationJobName *string

	// The timestamp of when the packaging job was created.
	CreationTime *time.Time

	// Returns a message describing the job status and error messages.
	EdgePackagingJobStatusMessage *string

	// The timestamp of when the job was last updated.
	LastModifiedTime *time.Time

	// The Amazon Simple Storage (S3) URI where model artifacts ares stored.
	ModelArtifact *string

	// The name of the model.
	ModelName *string

	// The signature document of files in the model artifact.
	ModelSignature *string

	// The version of the model.
	ModelVersion *string

	// The output configuration for the edge packaging job.
	OutputConfig *types.EdgeOutputConfig

	// The output of a SageMaker Edge Manager deployable resource.
	PresetDeploymentOutput *types.EdgePresetDeploymentOutput

	// The Amazon Web Services KMS key to use when encrypting the EBS volume the job
	// run on.
	ResourceKey *string

	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to
	// download and upload the model, and to contact Neo.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEdgePackagingJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEdgePackagingJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEdgePackagingJob{}, middleware.After)
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
	if err = addOpDescribeEdgePackagingJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEdgePackagingJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEdgePackagingJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeEdgePackagingJob",
	}
}
