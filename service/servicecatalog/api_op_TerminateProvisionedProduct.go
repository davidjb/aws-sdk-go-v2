// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Terminates the specified provisioned product. This operation does not delete any
// records associated with the provisioned product. You can check the status of
// this request using DescribeRecord.
func (c *Client) TerminateProvisionedProduct(ctx context.Context, params *TerminateProvisionedProductInput, optFns ...func(*Options)) (*TerminateProvisionedProductOutput, error) {
	if params == nil {
		params = &TerminateProvisionedProductInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TerminateProvisionedProduct", params, optFns, addOperationTerminateProvisionedProductMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TerminateProvisionedProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TerminateProvisionedProductInput struct {

	// An idempotency token that uniquely identifies the termination request. This
	// token is only valid during the termination process. After the provisioned
	// product is terminated, subsequent requests to terminate the same provisioned
	// product always return ResourceNotFound.
	//
	// This member is required.
	TerminateToken *string

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// If set to true, AWS Service Catalog stops managing the specified provisioned
	// product even if it cannot delete the underlying resources.
	IgnoreErrors bool

	// The identifier of the provisioned product. You cannot specify both
	// ProvisionedProductName and ProvisionedProductId.
	ProvisionedProductId *string

	// The name of the provisioned product. You cannot specify both
	// ProvisionedProductName and ProvisionedProductId.
	ProvisionedProductName *string

	// When this boolean parameter is set to true, the TerminateProvisionedProduct API
	// deletes the Service Catalog provisioned product. However, it does not remove the
	// CloudFormation stack, stack set, or the underlying resources of the deleted
	// provisioned product. The default value is false.
	RetainPhysicalResources bool
}

type TerminateProvisionedProductOutput struct {

	// Information about the result of this request.
	RecordDetail *types.RecordDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTerminateProvisionedProductMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTerminateProvisionedProduct{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTerminateProvisionedProduct{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
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
	if err = addIdempotencyToken_opTerminateProvisionedProductMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpTerminateProvisionedProductValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTerminateProvisionedProduct(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpTerminateProvisionedProduct struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpTerminateProvisionedProduct) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpTerminateProvisionedProduct) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*TerminateProvisionedProductInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *TerminateProvisionedProductInput ")
	}

	if input.TerminateToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.TerminateToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opTerminateProvisionedProductMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpTerminateProvisionedProduct{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opTerminateProvisionedProduct(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "TerminateProvisionedProduct",
	}
}
