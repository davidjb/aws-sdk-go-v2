// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts the reprocessing of raw message data through the pipeline.
func (c *Client) StartPipelineReprocessing(ctx context.Context, params *StartPipelineReprocessingInput, optFns ...func(*Options)) (*StartPipelineReprocessingOutput, error) {
	if params == nil {
		params = &StartPipelineReprocessingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartPipelineReprocessing", params, optFns, c.addOperationStartPipelineReprocessingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartPipelineReprocessingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartPipelineReprocessingInput struct {

	// The name of the pipeline on which to start reprocessing.
	//
	// This member is required.
	PipelineName *string

	// Specifies one or more sets of channel messages that you want to reprocess. If
	// you use the channelMessages object, you must not specify a value for startTime
	// and endTime .
	ChannelMessages *types.ChannelMessages

	// The end time (exclusive) of raw message data that is reprocessed. If you
	// specify a value for the endTime parameter, you must not use the channelMessages
	// object.
	EndTime *time.Time

	// The start time (inclusive) of raw message data that is reprocessed. If you
	// specify a value for the startTime parameter, you must not use the
	// channelMessages object.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type StartPipelineReprocessingOutput struct {

	// The ID of the pipeline reprocessing activity that was started.
	ReprocessingId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartPipelineReprocessingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartPipelineReprocessing{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartPipelineReprocessing{}, middleware.After)
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
	if err = addOpStartPipelineReprocessingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartPipelineReprocessing(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartPipelineReprocessing(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "StartPipelineReprocessing",
	}
}
