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
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action creates an Amazon S3 on Outposts bucket. To create an S3 bucket,
// see Create Bucket (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html)
// in the Amazon S3 API Reference. Creates a new Outposts bucket. By creating the
// bucket, you become the bucket owner. To create an Outposts bucket, you must have
// S3 on Outposts. For more information, see Using Amazon S3 on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html)
// in Amazon S3 User Guide. Not every string is an acceptable bucket name. For
// information on bucket naming restrictions, see Working with Amazon S3 Buckets (https://docs.aws.amazon.com/AmazonS3/latest/userguide/BucketRestrictions.html#bucketnamingrules)
// . S3 on Outposts buckets support:
//   - Tags
//   - LifecycleConfigurations for deleting expired objects
//
// For a complete list of restrictions and Amazon S3 feature limitations on S3 on
// Outposts, see Amazon S3 on Outposts Restrictions and Limitations (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OnOutpostsRestrictionsLimitations.html)
// . For an example of the request syntax for Amazon S3 on Outposts that uses the
// S3 on Outposts endpoint hostname prefix and x-amz-outpost-id in your API
// request, see the Examples (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html#API_control_CreateBucket_Examples)
// section. The following actions are related to CreateBucket for Amazon S3 on
// Outposts:
//   - PutObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html)
//   - GetBucket (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucket.html)
//   - DeleteBucket (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucket.html)
//   - CreateAccessPoint (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html)
//   - PutAccessPointPolicy (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicy.html)
func (c *Client) CreateBucket(ctx context.Context, params *CreateBucketInput, optFns ...func(*Options)) (*CreateBucketOutput, error) {
	if params == nil {
		params = &CreateBucketInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBucket", params, optFns, c.addOperationCreateBucketMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBucketOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBucketInput struct {

	// The name of the bucket.
	//
	// This member is required.
	Bucket *string

	// The canned ACL to apply to the bucket. This is not supported by Amazon S3 on
	// Outposts buckets.
	ACL types.BucketCannedACL

	// The configuration information for the bucket. This is not supported by Amazon
	// S3 on Outposts buckets.
	CreateBucketConfiguration *types.CreateBucketConfiguration

	// Allows grantee the read, write, read ACP, and write ACP permissions on the
	// bucket. This is not supported by Amazon S3 on Outposts buckets.
	GrantFullControl *string

	// Allows grantee to list the objects in the bucket. This is not supported by
	// Amazon S3 on Outposts buckets.
	GrantRead *string

	// Allows grantee to read the bucket ACL. This is not supported by Amazon S3 on
	// Outposts buckets.
	GrantReadACP *string

	// Allows grantee to create, overwrite, and delete any object in the bucket. This
	// is not supported by Amazon S3 on Outposts buckets.
	GrantWrite *string

	// Allows grantee to write the ACL for the applicable bucket. This is not
	// supported by Amazon S3 on Outposts buckets.
	GrantWriteACP *string

	// Specifies whether you want S3 Object Lock to be enabled for the new bucket.
	// This is not supported by Amazon S3 on Outposts buckets.
	ObjectLockEnabledForBucket bool

	// The ID of the Outposts where the bucket is being created. This ID is required
	// by Amazon S3 on Outposts buckets.
	OutpostId *string

	noSmithyDocumentSerde
}

type CreateBucketOutput struct {

	// The Amazon Resource Name (ARN) of the bucket. For using this parameter with
	// Amazon S3 on Outposts with the REST API, you must specify the name and the
	// x-amz-outpost-id as well. For using this parameter with S3 on Outposts with the
	// Amazon Web Services SDK and CLI, you must specify the ARN of the bucket accessed
	// in the format arn:aws:s3-outposts:::outpost//bucket/ . For example, to access
	// the bucket reports through Outpost my-outpost owned by account 123456789012 in
	// Region us-west-2 , use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports .
	// The value must be URL encoded.
	BucketArn *string

	// The location of the bucket.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBucketMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateBucket{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateBucket{}, middleware.After)
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
	if err = addCreateBucketResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateBucketValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBucket(options.Region), middleware.Before); err != nil {
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
	if err = addCreateBucketEndpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateBucket(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "CreateBucket",
	}
}

func copyCreateBucketInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*CreateBucketInput)
	if !ok {
		return nil, fmt.Errorf("expect *CreateBucketInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}

// getCreateBucketOutpostIDMember returns a pointer to string denoting a provided
// outpost-id member value and a boolean indicating if the input has a modeled
// outpost-id,
func getCreateBucketOutpostIDMember(input interface{}) (*string, bool) {
	in := input.(*CreateBucketInput)
	if in.OutpostId == nil {
		return nil, false
	}
	return in.OutpostId, true
}
func addCreateBucketUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: getCreateBucketOutpostIDMember,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyCreateBucketInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}

type opCreateBucketEndpointDisableHTTPSMiddleware struct {
	EndpointDisableHTTPS bool
}

func (*opCreateBucketEndpointDisableHTTPSMiddleware) ID() string {
	return "opCreateBucketEndpointDisableHTTPSMiddleware"
}

func (m *opCreateBucketEndpointDisableHTTPSMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
func addCreateBucketEndpointDisableHTTPSMiddleware(stack *middleware.Stack, o Options) error {
	return stack.Serialize.Insert(&opCreateBucketEndpointDisableHTTPSMiddleware{
		EndpointDisableHTTPS: o.EndpointOptions.DisableHTTPS,
	}, "opCreateBucketResolveEndpointMiddleware", middleware.After)
}

type opCreateBucketResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  BuiltInParameterResolver
}

func (*opCreateBucketResolveEndpointMiddleware) ID() string {
	return "opCreateBucketResolveEndpointMiddleware"
}

func (m *opCreateBucketResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*CreateBucketInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	params.Bucket = input.Bucket

	params.OutpostId = input.OutpostId

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

func addCreateBucketResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateBucketResolveEndpointMiddleware{
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
