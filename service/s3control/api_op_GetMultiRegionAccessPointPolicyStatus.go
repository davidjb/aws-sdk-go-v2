// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// Indicates whether the specified Multi-Region Access Point has an access control
// policy that allows public access. This action will always be routed to the US
// West (Oregon) Region. For more information about the restrictions around
// managing Multi-Region Access Points, see Managing Multi-Region Access Points (https://docs.aws.amazon.com/AmazonS3/latest/userguide/ManagingMultiRegionAccessPoints.html)
// in the Amazon S3 User Guide. The following actions are related to
// GetMultiRegionAccessPointPolicyStatus :
//   - GetMultiRegionAccessPointPolicy (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetMultiRegionAccessPointPolicy.html)
//   - PutMultiRegionAccessPointPolicy (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutMultiRegionAccessPointPolicy.html)
func (c *Client) GetMultiRegionAccessPointPolicyStatus(ctx context.Context, params *GetMultiRegionAccessPointPolicyStatusInput, optFns ...func(*Options)) (*GetMultiRegionAccessPointPolicyStatusOutput, error) {
	if params == nil {
		params = &GetMultiRegionAccessPointPolicyStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMultiRegionAccessPointPolicyStatus", params, optFns, c.addOperationGetMultiRegionAccessPointPolicyStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMultiRegionAccessPointPolicyStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMultiRegionAccessPointPolicyStatusInput struct {

	// The Amazon Web Services account ID for the owner of the Multi-Region Access
	// Point.
	//
	// This member is required.
	AccountId *string

	// Specifies the Multi-Region Access Point. The name of the Multi-Region Access
	// Point is different from the alias. For more information about the distinction
	// between the name and the alias of an Multi-Region Access Point, see Managing
	// Multi-Region Access Points (https://docs.aws.amazon.com/AmazonS3/latest/userguide/CreatingMultiRegionAccessPoints.html#multi-region-access-point-naming)
	// in the Amazon S3 User Guide.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type GetMultiRegionAccessPointPolicyStatusOutput struct {

	// Indicates whether this access point policy is public. For more information
	// about how Amazon S3 evaluates policies to determine whether they are public, see
	// The Meaning of "Public" (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status)
	// in the Amazon S3 User Guide.
	Established *types.PolicyStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMultiRegionAccessPointPolicyStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetMultiRegionAccessPointPolicyStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetMultiRegionAccessPointPolicyStatus{}, middleware.After)
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
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addGetMultiRegionAccessPointPolicyStatusResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetMultiRegionAccessPointPolicyStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMultiRegionAccessPointPolicyStatus(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addGetMultiRegionAccessPointPolicyStatusEndpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetMultiRegionAccessPointPolicyStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetMultiRegionAccessPointPolicyStatus",
	}
}

func copyGetMultiRegionAccessPointPolicyStatusInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*GetMultiRegionAccessPointPolicyStatusInput)
	if !ok {
		return nil, fmt.Errorf("expect *GetMultiRegionAccessPointPolicyStatusInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func backFillGetMultiRegionAccessPointPolicyStatusAccountID(input interface{}, v string) error {
	in := input.(*GetMultiRegionAccessPointPolicyStatusInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addGetMultiRegionAccessPointPolicyStatusUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyGetMultiRegionAccessPointPolicyStatusInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}

type opGetMultiRegionAccessPointPolicyStatusEndpointDisableHTTPSMiddleware struct {
	EndpointDisableHTTPS bool
}

func (*opGetMultiRegionAccessPointPolicyStatusEndpointDisableHTTPSMiddleware) ID() string {
	return "opGetMultiRegionAccessPointPolicyStatusEndpointDisableHTTPSMiddleware"
}

func (m *opGetMultiRegionAccessPointPolicyStatusEndpointDisableHTTPSMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointDisableHTTPS {
		req.URL.Scheme = "http"
	}

	return next.HandleSerialize(ctx, in)

}
func addGetMultiRegionAccessPointPolicyStatusEndpointDisableHTTPSMiddleware(stack *middleware.Stack, o Options) error {
	return stack.Serialize.Insert(&opGetMultiRegionAccessPointPolicyStatusEndpointDisableHTTPSMiddleware{
		EndpointDisableHTTPS: o.EndpointOptions.DisableHTTPS,
	}, "opGetMultiRegionAccessPointPolicyStatusResolveEndpointMiddleware", middleware.After)
}

type opGetMultiRegionAccessPointPolicyStatusResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  BuiltInParameterResolver
}

func (*opGetMultiRegionAccessPointPolicyStatusResolveEndpointMiddleware) ID() string {
	return "opGetMultiRegionAccessPointPolicyStatusResolveEndpointMiddleware"
}

func (m *opGetMultiRegionAccessPointPolicyStatusResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*GetMultiRegionAccessPointPolicyStatusInput)
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
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*BuiltInResolver).Region
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
			var signingName string
			if v4aScheme.SigningName == nil {
				signingName = "s3"
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addGetMultiRegionAccessPointPolicyStatusResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetMultiRegionAccessPointPolicyStatusResolveEndpointMiddleware{
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
