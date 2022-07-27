// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation allows you to request changes for your entities. Within a single
// ChangeSet, you cannot start the same change type against the same entity
// multiple times. Additionally, when a ChangeSet is running, all the entities
// targeted by the different changes are locked until the ChangeSet has completed
// (either succeeded, cancelled, or failed). If you try to start a ChangeSet
// containing a change against an entity that is already locked, you will receive a
// ResourceInUseException. For example, you cannot start the ChangeSet described in
// the example
// (https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/API_StartChangeSet.html#API_StartChangeSet_Examples)
// later in this topic, because it contains two changes to execute the same change
// type (AddRevisions) against the same entity (entity-id@1). For more information
// about working with change sets, see  Working with change sets
// (https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/welcome.html#working-with-change-sets).
func (c *Client) StartChangeSet(ctx context.Context, params *StartChangeSetInput, optFns ...func(*Options)) (*StartChangeSetOutput, error) {
	if params == nil {
		params = &StartChangeSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartChangeSet", params, optFns, c.addOperationStartChangeSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartChangeSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartChangeSetInput struct {

	// The catalog related to the request. Fixed value: AWSMarketplace
	//
	// This member is required.
	Catalog *string

	// Array of change object.
	//
	// This member is required.
	ChangeSet []types.Change

	// Optional case sensitive string of up to 100 ASCII characters. The change set
	// name can be used to filter the list of change sets.
	ChangeSetName *string

	// A unique token to identify the request to ensure idempotency.
	ClientRequestToken *string

	noSmithyDocumentSerde
}

type StartChangeSetOutput struct {

	// The ARN associated to the unique identifier generated for the request.
	ChangeSetArn *string

	// Unique identifier generated for the request.
	ChangeSetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartChangeSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartChangeSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartChangeSet{}, middleware.After)
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
	if err = addIdempotencyToken_opStartChangeSetMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartChangeSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartChangeSet(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartChangeSet struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartChangeSet) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartChangeSet) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartChangeSetInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartChangeSetInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartChangeSetMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartChangeSet{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartChangeSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aws-marketplace",
		OperationName: "StartChangeSet",
	}
}
