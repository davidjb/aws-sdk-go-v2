// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the permission policy for a version of an Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
// . For more information, see AddLayerVersionPermission .
func (c *Client) GetLayerVersionPolicy(ctx context.Context, params *GetLayerVersionPolicyInput, optFns ...func(*Options)) (*GetLayerVersionPolicyOutput, error) {
	if params == nil {
		params = &GetLayerVersionPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLayerVersionPolicy", params, optFns, c.addOperationGetLayerVersionPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLayerVersionPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLayerVersionPolicyInput struct {

	// The name or Amazon Resource Name (ARN) of the layer.
	//
	// This member is required.
	LayerName *string

	// The version number.
	//
	// This member is required.
	VersionNumber int64

	noSmithyDocumentSerde
}

type GetLayerVersionPolicyOutput struct {

	// The policy document.
	Policy *string

	// A unique identifier for the current revision of the policy.
	RevisionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLayerVersionPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetLayerVersionPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLayerVersionPolicy{}, middleware.After)
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
	if err = addGetLayerVersionPolicyResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetLayerVersionPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLayerVersionPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLayerVersionPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "GetLayerVersionPolicy",
	}
}

type opGetLayerVersionPolicyResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  BuiltInParameterResolver
}

func (*opGetLayerVersionPolicyResolveEndpointMiddleware) ID() string {
	return "opGetLayerVersionPolicyResolveEndpointMiddleware"
}

func (m *opGetLayerVersionPolicyResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	return next.HandleSerialize(ctx, in)
}

func addGetLayerVersionPolicyResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetLayerVersionPolicyResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &BuiltInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
