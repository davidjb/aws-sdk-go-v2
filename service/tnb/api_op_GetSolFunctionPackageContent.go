// Code generated by smithy-go-codegen DO NOT EDIT.

package tnb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/tnb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the contents of a function package. A function package is a .zip file in
// CSAR (Cloud Service Archive) format that contains a network function (an ETSI
// standard telecommunication application) and function package descriptor that
// uses the TOSCA standard to describe how the network functions should run on your
// network.
func (c *Client) GetSolFunctionPackageContent(ctx context.Context, params *GetSolFunctionPackageContentInput, optFns ...func(*Options)) (*GetSolFunctionPackageContentOutput, error) {
	if params == nil {
		params = &GetSolFunctionPackageContentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSolFunctionPackageContent", params, optFns, c.addOperationGetSolFunctionPackageContentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSolFunctionPackageContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSolFunctionPackageContentInput struct {

	// The format of the package that you want to download from the function packages.
	//
	// This member is required.
	Accept types.PackageContentType

	// ID of the function package.
	//
	// This member is required.
	VnfPkgId *string

	noSmithyDocumentSerde
}

type GetSolFunctionPackageContentOutput struct {

	// Indicates the media type of the resource.
	ContentType types.PackageContentType

	// Contents of the function package.
	PackageContent []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSolFunctionPackageContentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSolFunctionPackageContent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSolFunctionPackageContent{}, middleware.After)
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
	if err = addOpGetSolFunctionPackageContentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSolFunctionPackageContent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSolFunctionPackageContent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tnb",
		OperationName: "GetSolFunctionPackageContent",
	}
}
