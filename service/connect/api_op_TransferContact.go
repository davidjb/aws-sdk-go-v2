// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Transfers contacts from one agent or queue to another agent or queue at any
// point after a contact is created. You can transfer a contact to another queue by
// providing the contact flow which orchestrates the contact to the destination
// queue. This gives you more control over contact handling and helps you adhere to
// the service level agreement (SLA) guaranteed to your customers. Note the
// following requirements:
//
// * Transfer is supported for only TASK contacts.
//
// * Do
// not use both QueueId and UserId in the same call.
//
// * The following contact flow
// types are supported: Inbound contact flow, Transfer to agent flow, and Transfer
// to queue flow.
//
// * The TransferContact API can be called only on active
// contacts.
//
// * A contact cannot be transferred more than 11 times.
func (c *Client) TransferContact(ctx context.Context, params *TransferContactInput, optFns ...func(*Options)) (*TransferContactOutput, error) {
	if params == nil {
		params = &TransferContactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TransferContact", params, optFns, c.addOperationTransferContactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TransferContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TransferContactInput struct {

	// The identifier of the contact flow.
	//
	// This member is required.
	ContactFlowId *string

	// The identifier of the contact in this instance of Amazon Connect
	//
	// This member is required.
	ContactId *string

	// The identifier of the Amazon Connect instance. You can find the instanceId in
	// the ARN of the instance.
	//
	// This member is required.
	InstanceId *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// The identifier for the queue.
	QueueId *string

	// The identifier for the user.
	UserId *string

	noSmithyDocumentSerde
}

type TransferContactOutput struct {

	// The Amazon Resource Name (ARN) of the contact.
	ContactArn *string

	// The identifier of the contact in this instance of Amazon Connect
	ContactId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTransferContactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpTransferContact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpTransferContact{}, middleware.After)
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
	if err = addIdempotencyToken_opTransferContactMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpTransferContactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTransferContact(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpTransferContact struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpTransferContact) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpTransferContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*TransferContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *TransferContactInput ")
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
func addIdempotencyToken_opTransferContactMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpTransferContact{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opTransferContact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "TransferContact",
	}
}
