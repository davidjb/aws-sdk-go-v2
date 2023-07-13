// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkidentity/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the name and metadata of an AppInstanceBot .
func (c *Client) UpdateAppInstanceBot(ctx context.Context, params *UpdateAppInstanceBotInput, optFns ...func(*Options)) (*UpdateAppInstanceBotOutput, error) {
	if params == nil {
		params = &UpdateAppInstanceBotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAppInstanceBot", params, optFns, c.addOperationUpdateAppInstanceBotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAppInstanceBotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAppInstanceBotInput struct {

	// The ARN of the AppInstanceBot .
	//
	// This member is required.
	AppInstanceBotArn *string

	// The metadata of the AppInstanceBot .
	//
	// This member is required.
	Metadata *string

	// The name of the AppInstanceBot .
	//
	// This member is required.
	Name *string

	// The configuration for the bot update.
	Configuration *types.Configuration

	noSmithyDocumentSerde
}

type UpdateAppInstanceBotOutput struct {

	// The ARN of the AppInstanceBot .
	AppInstanceBotArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAppInstanceBotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAppInstanceBot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAppInstanceBot{}, middleware.After)
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
	if err = addOpUpdateAppInstanceBotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAppInstanceBot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAppInstanceBot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "UpdateAppInstanceBot",
	}
}
