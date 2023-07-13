// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a prompt.
func (c *Client) UpdatePrompt(ctx context.Context, params *UpdatePromptInput, optFns ...func(*Options)) (*UpdatePromptOutput, error) {
	if params == nil {
		params = &UpdatePromptInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePrompt", params, optFns, c.addOperationUpdatePromptMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePromptOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePromptInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// A unique identifier for the prompt.
	//
	// This member is required.
	PromptId *string

	// A description of the prompt.
	Description *string

	// The name of the prompt.
	Name *string

	// The URI for the S3 bucket where the prompt is stored.
	S3Uri *string

	noSmithyDocumentSerde
}

type UpdatePromptOutput struct {

	// The Amazon Resource Name (ARN) of the prompt.
	PromptARN *string

	// A unique identifier for the prompt.
	PromptId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePromptMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePrompt{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePrompt{}, middleware.After)
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
	if err = addOpUpdatePromptValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePrompt(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePrompt(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "UpdatePrompt",
	}
}
