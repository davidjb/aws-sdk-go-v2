// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the provisioned concurrency configuration for a function's alias or
// version.
func (c *Client) GetProvisionedConcurrencyConfig(ctx context.Context, params *GetProvisionedConcurrencyConfigInput, optFns ...func(*Options)) (*GetProvisionedConcurrencyConfigOutput, error) {
	if params == nil {
		params = &GetProvisionedConcurrencyConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetProvisionedConcurrencyConfig", params, optFns, c.addOperationGetProvisionedConcurrencyConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetProvisionedConcurrencyConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetProvisionedConcurrencyConfigInput struct {

	// The name of the Lambda function. Name formats
	//   - Function name – my-function .
	//   - Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function .
	//   - Partial ARN – 123456789012:function:my-function .
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// The version number or alias name.
	//
	// This member is required.
	Qualifier *string

	noSmithyDocumentSerde
}

type GetProvisionedConcurrencyConfigOutput struct {

	// The amount of provisioned concurrency allocated. When a weighted alias is used
	// during linear and canary deployments, this value fluctuates depending on the
	// amount of concurrency that is provisioned for the function versions.
	AllocatedProvisionedConcurrentExecutions *int32

	// The amount of provisioned concurrency available.
	AvailableProvisionedConcurrentExecutions *int32

	// The date and time that a user last updated the configuration, in ISO 8601 format (https://www.iso.org/iso-8601-date-and-time-format.html)
	// .
	LastModified *string

	// The amount of provisioned concurrency requested.
	RequestedProvisionedConcurrentExecutions *int32

	// The status of the allocation process.
	Status types.ProvisionedConcurrencyStatusEnum

	// For failed allocations, the reason that provisioned concurrency could not be
	// allocated.
	StatusReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetProvisionedConcurrencyConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetProvisionedConcurrencyConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetProvisionedConcurrencyConfig{}, middleware.After)
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
	if err = addGetProvisionedConcurrencyConfigResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetProvisionedConcurrencyConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetProvisionedConcurrencyConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetProvisionedConcurrencyConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "GetProvisionedConcurrencyConfig",
	}
}

type opGetProvisionedConcurrencyConfigResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  BuiltInParameterResolver
}

func (*opGetProvisionedConcurrencyConfigResolveEndpointMiddleware) ID() string {
	return "opGetProvisionedConcurrencyConfigResolveEndpointMiddleware"
}

func (m *opGetProvisionedConcurrencyConfigResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addGetProvisionedConcurrencyConfigResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetProvisionedConcurrencyConfigResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &BuiltInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
