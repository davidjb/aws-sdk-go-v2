// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalogappregistry

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates a resource from application. Both the resource and the
// application can be specified either by ID or name.
func (c *Client) DisassociateResource(ctx context.Context, params *DisassociateResourceInput, optFns ...func(*Options)) (*DisassociateResourceOutput, error) {
	if params == nil {
		params = &DisassociateResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateResource", params, optFns, c.addOperationDisassociateResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateResourceInput struct {

	// The name or ID of the application.
	//
	// This member is required.
	Application *string

	// The name or ID of the resource.
	//
	// This member is required.
	Resource *string

	// The type of the resource that is being disassociated.
	//
	// This member is required.
	ResourceType types.ResourceType

	noSmithyDocumentSerde
}

type DisassociateResourceOutput struct {

	// The Amazon resource name (ARN) that specifies the application.
	ApplicationArn *string

	// The Amazon resource name (ARN) that specifies the resource.
	ResourceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateResource{}, middleware.After)
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
	if err = addOpDisassociateResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "DisassociateResource",
	}
}
