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
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// This action deletes an Amazon S3 on Outposts bucket's lifecycle configuration.
// To delete an S3 bucket's lifecycle configuration, see DeleteBucketLifecycle (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketLifecycle.html)
// in the Amazon S3 API Reference. Deletes the lifecycle configuration from the
// specified Outposts bucket. Amazon S3 on Outposts removes all the lifecycle
// configuration rules in the lifecycle subresource associated with the bucket.
// Your objects never expire, and Amazon S3 on Outposts no longer automatically
// deletes any objects on the basis of rules contained in the deleted lifecycle
// configuration. For more information, see Using Amazon S3 on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html)
// in Amazon S3 User Guide. To use this action, you must have permission to perform
// the s3-outposts:DeleteLifecycleConfiguration action. By default, the bucket
// owner has this permission and the Outposts bucket owner can grant this
// permission to others. All Amazon S3 on Outposts REST API requests for this
// action require an additional parameter of x-amz-outpost-id to be passed with
// the request. In addition, you must use an S3 on Outposts endpoint hostname
// prefix instead of s3-control . For an example of the request syntax for Amazon
// S3 on Outposts that uses the S3 on Outposts endpoint hostname prefix and the
// x-amz-outpost-id derived by using the access point ARN, see the Examples (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketLifecycleConfiguration.html#API_control_DeleteBucketLifecycleConfiguration_Examples)
// section. For more information about object expiration, see Elements to Describe
// Lifecycle Actions (https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html#intro-lifecycle-rules-actions)
// . Related actions include:
//   - PutBucketLifecycleConfiguration (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketLifecycleConfiguration.html)
//   - GetBucketLifecycleConfiguration (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketLifecycleConfiguration.html)
func (c *Client) DeleteBucketLifecycleConfiguration(ctx context.Context, params *DeleteBucketLifecycleConfigurationInput, optFns ...func(*Options)) (*DeleteBucketLifecycleConfigurationOutput, error) {
	if params == nil {
		params = &DeleteBucketLifecycleConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBucketLifecycleConfiguration", params, optFns, c.addOperationDeleteBucketLifecycleConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBucketLifecycleConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBucketLifecycleConfigurationInput struct {

	// The account ID of the lifecycle configuration to delete.
	//
	// This member is required.
	AccountId *string

	// Specifies the bucket. For using this parameter with Amazon S3 on Outposts with
	// the REST API, you must specify the name and the x-amz-outpost-id as well. For
	// using this parameter with S3 on Outposts with the Amazon Web Services SDK and
	// CLI, you must specify the ARN of the bucket accessed in the format
	// arn:aws:s3-outposts:::outpost//bucket/ . For example, to access the bucket
	// reports through Outpost my-outpost owned by account 123456789012 in Region
	// us-west-2 , use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports .
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string

	noSmithyDocumentSerde
}

type DeleteBucketLifecycleConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteBucketLifecycleConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteBucketLifecycleConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteBucketLifecycleConfiguration{}, middleware.After)
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
	if err = addDeleteBucketLifecycleConfigurationResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteBucketLifecycleConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBucketLifecycleConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addDeleteBucketLifecycleConfigurationUpdateEndpoint(stack, options); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opDeleteBucketLifecycleConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteBucketLifecycleConfiguration",
	}
}

func copyDeleteBucketLifecycleConfigurationInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*DeleteBucketLifecycleConfigurationInput)
	if !ok {
		return nil, fmt.Errorf("expect *DeleteBucketLifecycleConfigurationInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getDeleteBucketLifecycleConfigurationARNMember(input interface{}) (*string, bool) {
	in := input.(*DeleteBucketLifecycleConfigurationInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func setDeleteBucketLifecycleConfigurationARNMember(input interface{}, v string) error {
	in := input.(*DeleteBucketLifecycleConfigurationInput)
	in.Bucket = &v
	return nil
}
func backFillDeleteBucketLifecycleConfigurationAccountID(input interface{}, v string) error {
	in := input.(*DeleteBucketLifecycleConfigurationInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addDeleteBucketLifecycleConfigurationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getDeleteBucketLifecycleConfigurationARNMember,
			BackfillAccountID: backFillDeleteBucketLifecycleConfigurationAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setDeleteBucketLifecycleConfigurationARNMember,
			CopyInput:         copyDeleteBucketLifecycleConfigurationInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}

type opDeleteBucketLifecycleConfigurationResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  BuiltInParameterResolver
}

func (*opDeleteBucketLifecycleConfigurationResolveEndpointMiddleware) ID() string {
	return "opDeleteBucketLifecycleConfigurationResolveEndpointMiddleware"
}

func (m *opDeleteBucketLifecycleConfigurationResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*DeleteBucketLifecycleConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	params.AccountId = input.AccountId

	params.Bucket = input.Bucket

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

func addDeleteBucketLifecycleConfigurationResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDeleteBucketLifecycleConfigurationResolveEndpointMiddleware{
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
