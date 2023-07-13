// Code generated by smithy-go-codegen DO NOT EDIT.

package mediastoredata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediastoredata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Uploads an object to the specified path. Object sizes are limited to 25 MB for
// standard upload availability and 10 MB for streaming upload availability.
func (c *Client) PutObject(ctx context.Context, params *PutObjectInput, optFns ...func(*Options)) (*PutObjectOutput, error) {
	if params == nil {
		params = &PutObjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutObject", params, optFns, c.addOperationPutObjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutObjectInput struct {

	// The bytes to be stored.
	//
	// This member is required.
	Body io.Reader

	// The path (including the file name) where the object is stored in the container.
	// Format: // For example, to upload the file mlaw.avi to the folder path
	// premium\canada in the container movies , enter the path premium/canada/mlaw.avi
	// . Do not include the container name in this path. If the path includes any
	// folders that don't exist yet, the service creates them. For example, suppose you
	// have an existing premium/usa subfolder. If you specify premium/canada , the
	// service creates a canada subfolder in the premium folder. You then have two
	// subfolders, usa and canada , in the premium folder. There is no correlation
	// between the path to the source and the path (folders) in the container in AWS
	// Elemental MediaStore. For more information about folders and how they exist in a
	// container, see the AWS Elemental MediaStore User Guide (http://docs.aws.amazon.com/mediastore/latest/ug/)
	// . The file name is the name that is assigned to the file that you upload. The
	// file can have the same name inside and outside of AWS Elemental MediaStore, or
	// it can have the same name. The file name can include or omit an extension.
	//
	// This member is required.
	Path *string

	// An optional CacheControl header that allows the caller to control the object's
	// cache behavior. Headers can be passed in as specified in the HTTP at
	// https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9 (https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9)
	// . Headers with a custom user-defined value are also accepted.
	CacheControl *string

	// The content type of the object.
	ContentType *string

	// Indicates the storage class of a Put request. Defaults to high-performance
	// temporal storage class, and objects are persisted into durable storage shortly
	// after being received.
	StorageClass types.StorageClass

	// Indicates the availability of an object while it is still uploading. If the
	// value is set to streaming , the object is available for downloading after some
	// initial buffering but before the object is uploaded completely. If the value is
	// set to standard , the object is available for downloading only when it is
	// uploaded completely. The default value for this header is standard . To use this
	// header, you must also set the HTTP Transfer-Encoding header to chunked .
	UploadAvailability types.UploadAvailability

	noSmithyDocumentSerde
}

type PutObjectOutput struct {

	// The SHA256 digest of the object that is persisted.
	ContentSHA256 *string

	// Unique identifier of the object in the container.
	ETag *string

	// The storage class where the object was persisted. The class should be
	// “Temporal”.
	StorageClass types.StorageClass

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutObjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutObject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutObject{}, middleware.After)
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
	if err = v4.AddUnsignedPayloadMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
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
	if err = addOpPutObjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutObject(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutObject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediastore",
		OperationName: "PutObject",
	}
}
