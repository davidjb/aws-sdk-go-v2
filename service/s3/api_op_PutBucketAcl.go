// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	internalChecksum "github.com/aws/aws-sdk-go-v2/service/internal/checksum"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the permissions on an existing bucket using access control lists (ACL).
// For more information, see Using ACLs (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3_ACLs_UsingACLs.html)
// . To set the ACL of a bucket, you must have WRITE_ACP permission. You can use
// one of the following two ways to set a bucket's permissions:
//   - Specify the ACL in the request body
//   - Specify permissions using request headers
//
// You cannot specify access permission using both the body and the request
// headers. Depending on your application needs, you may choose to set the ACL on a
// bucket using either the request body or the headers. For example, if you have an
// existing application that updates a bucket ACL using the request body, then you
// can continue to use that approach. If your bucket uses the bucket owner enforced
// setting for S3 Object Ownership, ACLs are disabled and no longer affect
// permissions. You must use policies to grant access to your bucket and the
// objects in it. Requests to set ACLs or update ACLs fail and return the
// AccessControlListNotSupported error code. Requests to read ACLs are still
// supported. For more information, see Controlling object ownership (https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html)
// in the Amazon S3 User Guide. Permissions You can set access permissions using
// one of the following methods:
//   - Specify a canned ACL with the x-amz-acl request header. Amazon S3 supports a
//     set of predefined ACLs, known as canned ACLs. Each canned ACL has a predefined
//     set of grantees and permissions. Specify the canned ACL name as the value of
//     x-amz-acl . If you use this header, you cannot use other access
//     control-specific headers in your request. For more information, see Canned ACL (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL)
//     .
//   - Specify access permissions explicitly with the x-amz-grant-read ,
//     x-amz-grant-read-acp , x-amz-grant-write-acp , and x-amz-grant-full-control
//     headers. When using these headers, you specify explicit access permissions and
//     grantees (Amazon Web Services accounts or Amazon S3 groups) who will receive the
//     permission. If you use these ACL-specific headers, you cannot use the
//     x-amz-acl header to set a canned ACL. These parameters map to the set of
//     permissions that Amazon S3 supports in an ACL. For more information, see
//     Access Control List (ACL) Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html)
//     . You specify each grantee as a type=value pair, where the type is one of the
//     following:
//   - id – if the value specified is the canonical user ID of an Amazon Web
//     Services account
//   - uri – if you are granting permissions to a predefined group
//   - emailAddress – if the value specified is the email address of an Amazon Web
//     Services account Using email addresses to specify a grantee is only supported in
//     the following Amazon Web Services Regions:
//   - US East (N. Virginia)
//   - US West (N. California)
//   - US West (Oregon)
//   - Asia Pacific (Singapore)
//   - Asia Pacific (Sydney)
//   - Asia Pacific (Tokyo)
//   - Europe (Ireland)
//   - South America (São Paulo) For a list of all the Amazon S3 supported Regions
//     and endpoints, see Regions and Endpoints (https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region)
//     in the Amazon Web Services General Reference. For example, the following
//     x-amz-grant-write header grants create, overwrite, and delete objects
//     permission to LogDelivery group predefined by Amazon S3 and two Amazon Web
//     Services accounts identified by their email addresses. x-amz-grant-write:
//     uri="http://acs.amazonaws.com/groups/s3/LogDelivery", id="111122223333",
//     id="555566667777"
//
// You can use either a canned ACL or specify access permissions explicitly. You
// cannot do both. Grantee Values You can specify the person (grantee) to whom
// you're assigning access rights (using request elements) in the following ways:
//   - By the person's ID: <>ID<><>GranteesEmail<> DisplayName is optional and
//     ignored in the request
//   - By URI: <>http://acs.amazonaws.com/groups/global/AuthenticatedUsers<>
//   - By Email address: <>Grantees@email.com<>& The grantee is resolved to the
//     CanonicalUser and, in a response to a GET Object acl request, appears as the
//     CanonicalUser. Using email addresses to specify a grantee is only supported in
//     the following Amazon Web Services Regions:
//   - US East (N. Virginia)
//   - US West (N. California)
//   - US West (Oregon)
//   - Asia Pacific (Singapore)
//   - Asia Pacific (Sydney)
//   - Asia Pacific (Tokyo)
//   - Europe (Ireland)
//   - South America (São Paulo) For a list of all the Amazon S3 supported Regions
//     and endpoints, see Regions and Endpoints (https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region)
//     in the Amazon Web Services General Reference.
//
// The following operations are related to PutBucketAcl :
//   - CreateBucket (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html)
//   - DeleteBucket (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html)
//   - GetObjectAcl (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAcl.html)
func (c *Client) PutBucketAcl(ctx context.Context, params *PutBucketAclInput, optFns ...func(*Options)) (*PutBucketAclOutput, error) {
	if params == nil {
		params = &PutBucketAclInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketAcl", params, optFns, c.addOperationPutBucketAclMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketAclOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketAclInput struct {

	// The bucket to which to apply the ACL.
	//
	// This member is required.
	Bucket *string

	// The canned ACL to apply to the bucket.
	ACL types.BucketCannedACL

	// Contains the elements that set the ACL permissions for an object per grantee.
	AccessControlPolicy *types.AccessControlPolicy

	// Indicates the algorithm used to create the checksum for the object when using
	// the SDK. This header will not provide any additional functionality if not using
	// the SDK. When sending this header, there must be a corresponding x-amz-checksum
	// or x-amz-trailer header sent. Otherwise, Amazon S3 fails the request with the
	// HTTP status code 400 Bad Request . For more information, see Checking object
	// integrity (https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html)
	// in the Amazon S3 User Guide. If you provide an individual checksum, Amazon S3
	// ignores any provided ChecksumAlgorithm parameter.
	ChecksumAlgorithm types.ChecksumAlgorithm

	// The base64-encoded 128-bit MD5 digest of the data. This header must be used as
	// a message integrity check to verify that the request body was not corrupted in
	// transit. For more information, go to RFC 1864. (http://www.ietf.org/rfc/rfc1864.txt)
	// For requests made using the Amazon Web Services Command Line Interface (CLI) or
	// Amazon Web Services SDKs, this field is calculated automatically.
	ContentMD5 *string

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	ExpectedBucketOwner *string

	// Allows grantee the read, write, read ACP, and write ACP permissions on the
	// bucket.
	GrantFullControl *string

	// Allows grantee to list the objects in the bucket.
	GrantRead *string

	// Allows grantee to read the bucket ACL.
	GrantReadACP *string

	// Allows grantee to create new objects in the bucket. For the bucket and object
	// owners of existing objects, also allows deletions and overwrites of those
	// objects.
	GrantWrite *string

	// Allows grantee to write the ACL for the applicable bucket.
	GrantWriteACP *string

	noSmithyDocumentSerde
}

type PutBucketAclOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBucketAclMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketAcl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketAcl{}, middleware.After)
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
	if err = swapWithCustomHTTPSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addPutBucketAclResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutBucketAclValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketAcl(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addPutBucketAclInputChecksumMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addEndpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSerializeImmutableHostnameBucketMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func (v *PutBucketAclInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opPutBucketAcl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketAcl",
	}
}

// getPutBucketAclRequestAlgorithmMember gets the request checksum algorithm value
// provided as input.
func getPutBucketAclRequestAlgorithmMember(input interface{}) (string, bool) {
	in := input.(*PutBucketAclInput)
	if len(in.ChecksumAlgorithm) == 0 {
		return "", false
	}
	return string(in.ChecksumAlgorithm), true
}

func addPutBucketAclInputChecksumMiddlewares(stack *middleware.Stack, options Options) error {
	return internalChecksum.AddInputMiddleware(stack, internalChecksum.InputMiddlewareOptions{
		GetAlgorithm:                     getPutBucketAclRequestAlgorithmMember,
		RequireChecksum:                  true,
		EnableTrailingChecksum:           false,
		EnableComputeSHA256PayloadHash:   true,
		EnableDecodedContentLengthHeader: true,
	})
}

// getPutBucketAclBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getPutBucketAclBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketAclInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutBucketAclUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutBucketAclBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}

type opPutBucketAclResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  BuiltInParameterResolver
}

func (*opPutBucketAclResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opPutBucketAclResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*PutBucketAclInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	params.Bucket = input.Bucket

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
			ctx = s3cust.SetSignerVersion(ctx, internalauth.SigV4)
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
			ctx = s3cust.SetSignerVersion(ctx, v4Scheme.Name)
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
			ctx = s3cust.SetSignerVersion(ctx, v4aScheme.Name)
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addPutBucketAclResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opPutBucketAclResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &BuiltInResolver{
			Region:                         options.Region,
			UseFIPS:                        options.EndpointOptions.UseFIPSEndpoint,
			UseDualStack:                   options.EndpointOptions.UseDualStackEndpoint,
			Endpoint:                       options.BaseEndpoint,
			ForcePathStyle:                 options.UsePathStyle,
			Accelerate:                     options.UseAccelerate,
			DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
			UseArnRegion:                   options.UseARNRegion,
		},
	}, "ResolveEndpoint", middleware.After)
}
