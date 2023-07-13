// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Upgrades a single directory in-place using the PublishedSchemaArn with schema
// updates found in MinorVersion . Backwards-compatible minor version upgrades are
// instantaneously available for readers on all objects in the directory. Note:
// This is a synchronous API call and upgrades only one schema on a given directory
// per call. To upgrade multiple directories from one schema, you would need to
// call this API on each directory.
func (c *Client) UpgradeAppliedSchema(ctx context.Context, params *UpgradeAppliedSchemaInput, optFns ...func(*Options)) (*UpgradeAppliedSchemaOutput, error) {
	if params == nil {
		params = &UpgradeAppliedSchemaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpgradeAppliedSchema", params, optFns, c.addOperationUpgradeAppliedSchemaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpgradeAppliedSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpgradeAppliedSchemaInput struct {

	// The ARN for the directory to which the upgraded schema will be applied.
	//
	// This member is required.
	DirectoryArn *string

	// The revision of the published schema to upgrade the directory to.
	//
	// This member is required.
	PublishedSchemaArn *string

	// Used for testing whether the major version schemas are backward compatible or
	// not. If schema compatibility fails, an exception would be thrown else the call
	// would succeed but no changes will be saved. This parameter is optional.
	DryRun bool

	noSmithyDocumentSerde
}

type UpgradeAppliedSchemaOutput struct {

	// The ARN of the directory that is returned as part of the response.
	DirectoryArn *string

	// The ARN of the upgraded schema that is returned as part of the response.
	UpgradedSchemaArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpgradeAppliedSchemaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpgradeAppliedSchema{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpgradeAppliedSchema{}, middleware.After)
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
	if err = addOpUpgradeAppliedSchemaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpgradeAppliedSchema(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpgradeAppliedSchema(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "UpgradeAppliedSchema",
	}
}
