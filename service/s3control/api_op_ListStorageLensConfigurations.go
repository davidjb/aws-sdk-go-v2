// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// Gets a list of Amazon S3 Storage Lens configurations. For more information
// about S3 Storage Lens, see Assessing your storage activity and usage with
// Amazon S3 Storage Lens  (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens.html)
// in the Amazon S3 User Guide. To use this action, you must have permission to
// perform the s3:ListStorageLensConfigurations action. For more information, see
// Setting permissions to use Amazon S3 Storage Lens (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens_iam_permissions.html)
// in the Amazon S3 User Guide.
func (c *Client) ListStorageLensConfigurations(ctx context.Context, params *ListStorageLensConfigurationsInput, optFns ...func(*Options)) (*ListStorageLensConfigurationsOutput, error) {
	if params == nil {
		params = &ListStorageLensConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStorageLensConfigurations", params, optFns, c.addOperationListStorageLensConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStorageLensConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStorageLensConfigurationsInput struct {

	// The account ID of the requester.
	//
	// This member is required.
	AccountId *string

	// A pagination token to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListStorageLensConfigurationsOutput struct {

	// If the request produced more than the maximum number of S3 Storage Lens
	// configuration results, you can pass this value into a subsequent request to
	// retrieve the next page of results.
	NextToken *string

	// A list of S3 Storage Lens configurations.
	StorageLensConfigurationList []types.ListStorageLensConfigurationEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStorageLensConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListStorageLensConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListStorageLensConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addLegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opListStorageLensConfigurationsMiddleware(stack); err != nil {
		return err
	}
	if err = addListStorageLensConfigurationsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListStorageLensConfigurationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStorageLensConfigurations(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addListStorageLensConfigurationsUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addEndpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opListStorageLensConfigurationsMiddleware struct {
}

func (*endpointPrefix_opListStorageLensConfigurationsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListStorageLensConfigurationsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*ListStorageLensConfigurationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListStorageLensConfigurationsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListStorageLensConfigurationsMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListStorageLensConfigurationsAPIClient is a client that implements the
// ListStorageLensConfigurations operation.
type ListStorageLensConfigurationsAPIClient interface {
	ListStorageLensConfigurations(context.Context, *ListStorageLensConfigurationsInput, ...func(*Options)) (*ListStorageLensConfigurationsOutput, error)
}

var _ ListStorageLensConfigurationsAPIClient = (*Client)(nil)

// ListStorageLensConfigurationsPaginatorOptions is the paginator options for
// ListStorageLensConfigurations
type ListStorageLensConfigurationsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStorageLensConfigurationsPaginator is a paginator for
// ListStorageLensConfigurations
type ListStorageLensConfigurationsPaginator struct {
	options   ListStorageLensConfigurationsPaginatorOptions
	client    ListStorageLensConfigurationsAPIClient
	params    *ListStorageLensConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListStorageLensConfigurationsPaginator returns a new
// ListStorageLensConfigurationsPaginator
func NewListStorageLensConfigurationsPaginator(client ListStorageLensConfigurationsAPIClient, params *ListStorageLensConfigurationsInput, optFns ...func(*ListStorageLensConfigurationsPaginatorOptions)) *ListStorageLensConfigurationsPaginator {
	if params == nil {
		params = &ListStorageLensConfigurationsInput{}
	}

	options := ListStorageLensConfigurationsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStorageLensConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStorageLensConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStorageLensConfigurations page.
func (p *ListStorageLensConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStorageLensConfigurationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListStorageLensConfigurations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListStorageLensConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListStorageLensConfigurations",
	}
}

func copyListStorageLensConfigurationsInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*ListStorageLensConfigurationsInput)
	if !ok {
		return nil, fmt.Errorf("expect *ListStorageLensConfigurationsInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func backFillListStorageLensConfigurationsAccountID(input interface{}, v string) error {
	in := input.(*ListStorageLensConfigurationsInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addListStorageLensConfigurationsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyListStorageLensConfigurationsInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}

type opListStorageLensConfigurationsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  BuiltInParameterResolver
}

func (*opListStorageLensConfigurationsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListStorageLensConfigurationsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*ListStorageLensConfigurationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	params.AccountId = input.AccountId

	params.RequiresAccountId = ptr.Bool(true)

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

	ctx = smithyhttp.DisableEndpointHostPrefix(ctx, true)

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "s3"
			signingRegion := m.BuiltInResolver.(*BuiltInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "s3"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*BuiltInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("s3")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addListStorageLensConfigurationsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListStorageLensConfigurationsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &BuiltInResolver{
			Region:       options.Region,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			Endpoint:     options.BaseEndpoint,
			UseArnRegion: options.UseARNRegion,
		},
	}, "ResolveEndpoint", middleware.After)
}
