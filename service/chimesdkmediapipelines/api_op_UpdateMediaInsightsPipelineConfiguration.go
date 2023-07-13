// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmediapipelines

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the media insights pipeline's configuration settings.
func (c *Client) UpdateMediaInsightsPipelineConfiguration(ctx context.Context, params *UpdateMediaInsightsPipelineConfigurationInput, optFns ...func(*Options)) (*UpdateMediaInsightsPipelineConfigurationOutput, error) {
	if params == nil {
		params = &UpdateMediaInsightsPipelineConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMediaInsightsPipelineConfiguration", params, optFns, c.addOperationUpdateMediaInsightsPipelineConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMediaInsightsPipelineConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMediaInsightsPipelineConfigurationInput struct {

	// The elements in the request, such as a processor for Amazon Transcribe or a
	// sink for a Kinesis Data Stream..
	//
	// This member is required.
	Elements []types.MediaInsightsPipelineConfigurationElement

	// The unique identifier for the resource to be updated. Valid values include the
	// name and ARN of the media insights pipeline configuration.
	//
	// This member is required.
	Identifier *string

	// The ARN of the role used by the service to access Amazon Web Services resources.
	//
	// This member is required.
	ResourceAccessRoleArn *string

	// The configuration settings for real-time alerts for the media insights pipeline.
	RealTimeAlertConfiguration *types.RealTimeAlertConfiguration

	noSmithyDocumentSerde
}

type UpdateMediaInsightsPipelineConfigurationOutput struct {

	// The updated configuration settings.
	MediaInsightsPipelineConfiguration *types.MediaInsightsPipelineConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMediaInsightsPipelineConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateMediaInsightsPipelineConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateMediaInsightsPipelineConfiguration{}, middleware.After)
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
	if err = addOpUpdateMediaInsightsPipelineConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMediaInsightsPipelineConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateMediaInsightsPipelineConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "UpdateMediaInsightsPipelineConfiguration",
	}
}
