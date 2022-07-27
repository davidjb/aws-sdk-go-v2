// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create an endpoint group for the specified listener for a custom routing
// accelerator. An endpoint group is a collection of endpoints in one Amazon Web
// Services Region.
func (c *Client) CreateCustomRoutingEndpointGroup(ctx context.Context, params *CreateCustomRoutingEndpointGroupInput, optFns ...func(*Options)) (*CreateCustomRoutingEndpointGroupOutput, error) {
	if params == nil {
		params = &CreateCustomRoutingEndpointGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCustomRoutingEndpointGroup", params, optFns, c.addOperationCreateCustomRoutingEndpointGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCustomRoutingEndpointGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCustomRoutingEndpointGroupInput struct {

	// Sets the port range and protocol for all endpoints (virtual private cloud
	// subnets) in a custom routing endpoint group to accept client traffic on.
	//
	// This member is required.
	DestinationConfigurations []types.CustomRoutingDestinationConfiguration

	// The Amazon Web Services Region where the endpoint group is located. A listener
	// can have only one endpoint group in a specific Region.
	//
	// This member is required.
	EndpointGroupRegion *string

	// A unique, case-sensitive identifier that you provide to ensure the
	// idempotency—that is, the uniqueness—of the request.
	//
	// This member is required.
	IdempotencyToken *string

	// The Amazon Resource Name (ARN) of the listener for a custom routing endpoint.
	//
	// This member is required.
	ListenerArn *string

	noSmithyDocumentSerde
}

type CreateCustomRoutingEndpointGroupOutput struct {

	// The information about the endpoint group created for a custom routing
	// accelerator.
	EndpointGroup *types.CustomRoutingEndpointGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCustomRoutingEndpointGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCustomRoutingEndpointGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCustomRoutingEndpointGroup{}, middleware.After)
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateCustomRoutingEndpointGroupMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateCustomRoutingEndpointGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCustomRoutingEndpointGroup(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateCustomRoutingEndpointGroup struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateCustomRoutingEndpointGroup) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateCustomRoutingEndpointGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateCustomRoutingEndpointGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateCustomRoutingEndpointGroupInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateCustomRoutingEndpointGroupMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateCustomRoutingEndpointGroup{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateCustomRoutingEndpointGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "CreateCustomRoutingEndpointGroup",
	}
}
