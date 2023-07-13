// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcases

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connectcases/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the attributes of an existing layout. If the action is successful, the
// service sends back an HTTP 200 response with an empty HTTP body. A
// ValidationException is returned when you add non-existent fieldIds to a layout.
// Title and Status fields cannot be part of layouts because they are not
// configurable.
func (c *Client) UpdateLayout(ctx context.Context, params *UpdateLayoutInput, optFns ...func(*Options)) (*UpdateLayoutOutput, error) {
	if params == nil {
		params = &UpdateLayoutInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLayout", params, optFns, c.addOperationUpdateLayoutMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLayoutOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLayoutInput struct {

	// The unique identifier of the Cases domain.
	//
	// This member is required.
	DomainId *string

	// The unique identifier of the layout.
	//
	// This member is required.
	LayoutId *string

	// Information about which fields will be present in the layout, the order of the
	// fields, and a read-only attribute of the field.
	Content types.LayoutContent

	// The name of the layout. It must be unique per domain.
	Name *string

	noSmithyDocumentSerde
}

type UpdateLayoutOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLayoutMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateLayout{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateLayout{}, middleware.After)
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
	if err = addOpUpdateLayoutValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLayout(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLayout(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cases",
		OperationName: "UpdateLayout",
	}
}
