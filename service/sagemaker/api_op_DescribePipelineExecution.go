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

// Describes the details of a pipeline execution.
func (c *Client) DescribePipelineExecution(ctx context.Context, params *DescribePipelineExecutionInput, optFns ...func(*Options)) (*DescribePipelineExecutionOutput, error) {
	if params == nil {
		params = &DescribePipelineExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePipelineExecution", params, optFns, c.addOperationDescribePipelineExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePipelineExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePipelineExecutionInput struct {

	// The Amazon Resource Name (ARN) of the pipeline execution.
	//
	// This member is required.
	PipelineExecutionArn *string

	noSmithyDocumentSerde
}

type DescribePipelineExecutionOutput struct {

	// Information about the user who created or modified an experiment, trial, trial
	// component, lineage group, project, or model card.
	CreatedBy *types.UserContext

	// The time when the pipeline execution was created.
	CreationTime *time.Time

	// If the execution failed, a message describing why.
	FailureReason *string

	// Information about the user who created or modified an experiment, trial, trial
	// component, lineage group, project, or model card.
	LastModifiedBy *types.UserContext

	// The time when the pipeline execution was modified last.
	LastModifiedTime *time.Time

	// The parallelism configuration applied to the pipeline.
	ParallelismConfiguration *types.ParallelismConfiguration

	// The Amazon Resource Name (ARN) of the pipeline.
	PipelineArn *string

	// The Amazon Resource Name (ARN) of the pipeline execution.
	PipelineExecutionArn *string

	// The description of the pipeline execution.
	PipelineExecutionDescription *string

	// The display name of the pipeline execution.
	PipelineExecutionDisplayName *string

	// The status of the pipeline execution.
	PipelineExecutionStatus types.PipelineExecutionStatus

	// Specifies the names of the experiment and trial created by a pipeline.
	PipelineExperimentConfig *types.PipelineExperimentConfig

	// The selective execution configuration applied to the pipeline run.
	SelectiveExecutionConfig *types.SelectiveExecutionConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePipelineExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePipelineExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePipelineExecution{}, middleware.After)
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
	if err = addOpDescribePipelineExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePipelineExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribePipelineExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribePipelineExecution",
	}
}
