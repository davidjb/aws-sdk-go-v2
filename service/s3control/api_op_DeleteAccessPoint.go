// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"strings"
)

// Deletes the specified access point. All Amazon S3 on Outposts REST API requests
// for this action require an additional parameter of x-amz-outpost-id to be passed
// with the request and an S3 on Outposts endpoint hostname prefix instead of
// s3-control. For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived using the access point ARN, see the Examples
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPoint.html#API_control_DeleteAccessPoint_Examples)
// section. The following actions are related to DeleteAccessPoint:
//
// *
// CreateAccessPoint
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html)
//
// *
// GetAccessPoint
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html)
//
// *
// ListAccessPoints
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPoints.html)
func (c *Client) DeleteAccessPoint(ctx context.Context, params *DeleteAccessPointInput, optFns ...func(*Options)) (*DeleteAccessPointOutput, error) {
	if params == nil {
		params = &DeleteAccessPointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAccessPoint", params, optFns, addOperationDeleteAccessPointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAccessPointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAccessPointInput struct {

	// The account ID for the account that owns the specified access point.
	//
	// This member is required.
	AccountId *string

	// The name of the access point you want to delete. For using this parameter with
	// Amazon S3 on Outposts with the REST API, you must specify the name and the
	// x-amz-outpost-id as well. For using this parameter with S3 on Outposts with the
	// AWS SDK and CLI, you must specify the ARN of the access point accessed in the
	// format arn:aws:s3-outposts:::outpost//accesspoint/. For example, to access the
	// access point reports-ap through outpost my-outpost owned by account 123456789012
	// in Region us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/accesspoint/reports-ap.
	// The value must be URL encoded.
	//
	// This member is required.
	Name *string
}

type DeleteAccessPointOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteAccessPointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteAccessPoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteAccessPoint{}, middleware.After)
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
	if err = addEndpointPrefix_opDeleteAccessPointMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeleteAccessPointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAccessPoint(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addDeleteAccessPointUpdateEndpoint(stack, options); err != nil {
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

type endpointPrefix_opDeleteAccessPointMiddleware struct {
}

func (*endpointPrefix_opDeleteAccessPointMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDeleteAccessPointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*DeleteAccessPointInput)
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
func addEndpointPrefix_opDeleteAccessPointMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDeleteAccessPointMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteAccessPoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteAccessPoint",
	}
}

func copyDeleteAccessPointInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*DeleteAccessPointInput)
	if !ok {
		return nil, fmt.Errorf("expect *DeleteAccessPointInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getDeleteAccessPointARNMember(input interface{}) (*string, bool) {
	in := input.(*DeleteAccessPointInput)
	if in.Name == nil {
		return nil, false
	}
	return in.Name, true
}
func setDeleteAccessPointARNMember(input interface{}, v string) error {
	in := input.(*DeleteAccessPointInput)
	in.Name = &v
	return nil
}
func backFillDeleteAccessPointAccountID(input interface{}, v string) error {
	in := input.(*DeleteAccessPointInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addDeleteAccessPointUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getDeleteAccessPointARNMember,
			BackfillAccountID: backFillDeleteAccessPointAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setDeleteAccessPointARNMember,
			CopyInput:         copyDeleteAccessPointInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}
