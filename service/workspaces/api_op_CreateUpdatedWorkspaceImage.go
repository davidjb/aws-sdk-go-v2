// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new updated WorkSpace image based on the specified source image. The
// new updated WorkSpace image has the latest drivers and other updates required by
// the Amazon WorkSpaces components. To determine which WorkSpace images need to be
// updated with the latest Amazon WorkSpaces requirements, use
// DescribeWorkspaceImages
// (https://docs.aws.amazon.com/workspaces/latest/api/API_DescribeWorkspaceImages.html).
//
// *
// Only Windows 10, Windows Server 2016, and Windows Server 2019 WorkSpace images
// can be programmatically updated at this time.
//
// * Microsoft Windows updates and
// other application updates are not included in the update process.
//
// * The source
// WorkSpace image is not deleted. You can delete the source image after you've
// verified your new updated image and created a new bundle.
func (c *Client) CreateUpdatedWorkspaceImage(ctx context.Context, params *CreateUpdatedWorkspaceImageInput, optFns ...func(*Options)) (*CreateUpdatedWorkspaceImageOutput, error) {
	if params == nil {
		params = &CreateUpdatedWorkspaceImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUpdatedWorkspaceImage", params, optFns, c.addOperationCreateUpdatedWorkspaceImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUpdatedWorkspaceImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUpdatedWorkspaceImageInput struct {

	// A description of whether updates for the WorkSpace image are available.
	//
	// This member is required.
	Description *string

	// The name of the new updated WorkSpace image.
	//
	// This member is required.
	Name *string

	// The identifier of the source WorkSpace image.
	//
	// This member is required.
	SourceImageId *string

	// The tags that you want to add to the new updated WorkSpace image. To add tags at
	// the same time when you're creating the updated image, you must create an IAM
	// policy that grants your IAM user permissions to use workspaces:CreateTags.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateUpdatedWorkspaceImageOutput struct {

	// The identifier of the new updated WorkSpace image.
	ImageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateUpdatedWorkspaceImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUpdatedWorkspaceImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUpdatedWorkspaceImage{}, middleware.After)
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
	if err = addOpCreateUpdatedWorkspaceImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUpdatedWorkspaceImage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateUpdatedWorkspaceImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "CreateUpdatedWorkspaceImage",
	}
}
