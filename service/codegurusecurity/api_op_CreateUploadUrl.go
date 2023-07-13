// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurusecurity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates a pre-signed URL and request headers used to upload a code resource.
// You can upload your code resource to the URL and add the request headers using
// any HTTP client.
func (c *Client) CreateUploadUrl(ctx context.Context, params *CreateUploadUrlInput, optFns ...func(*Options)) (*CreateUploadUrlOutput, error) {
	if params == nil {
		params = &CreateUploadUrlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUploadUrl", params, optFns, c.addOperationCreateUploadUrlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUploadUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUploadUrlInput struct {

	// The name of the scan that will use the uploaded resource. CodeGuru Security
	// uses the unique scan name to track revisions across multiple scans of the same
	// resource. Use this scanName when you call CreateScan on the code resource you
	// upload to this URL.
	//
	// This member is required.
	ScanName *string

	noSmithyDocumentSerde
}

type CreateUploadUrlOutput struct {

	// The identifier for the uploaded code resource.
	//
	// This member is required.
	CodeArtifactId *string

	// A set of key-value pairs that contain the required headers when uploading your
	// resource.
	//
	// This member is required.
	RequestHeaders map[string]string

	// A pre-signed S3 URL. You can upload the code file you want to scan and add the
	// required requestHeaders using any HTTP client.
	//
	// This member is required.
	S3Url *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateUploadUrlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateUploadUrl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateUploadUrl{}, middleware.After)
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
	if err = addOpCreateUploadUrlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUploadUrl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateUploadUrl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-security",
		OperationName: "CreateUploadUrl",
	}
}
