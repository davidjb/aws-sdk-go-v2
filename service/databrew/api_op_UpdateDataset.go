// Code generated by smithy-go-codegen DO NOT EDIT.

package databrew

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databrew/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the definition of an existing DataBrew dataset.
func (c *Client) UpdateDataset(ctx context.Context, params *UpdateDatasetInput, optFns ...func(*Options)) (*UpdateDatasetOutput, error) {
	if params == nil {
		params = &UpdateDatasetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDataset", params, optFns, c.addOperationUpdateDatasetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDatasetInput struct {

	// Represents information on how DataBrew can find data, in either the Glue Data
	// Catalog or Amazon S3.
	//
	// This member is required.
	Input *types.Input

	// The name of the dataset to be updated.
	//
	// This member is required.
	Name *string

	// The file format of a dataset that is created from an Amazon S3 file or folder.
	Format types.InputFormat

	// Represents a set of options that define the structure of either comma-separated
	// value (CSV), Excel, or JSON input.
	FormatOptions *types.FormatOptions

	// A set of options that defines how DataBrew interprets an Amazon S3 path of the
	// dataset.
	PathOptions *types.PathOptions

	noSmithyDocumentSerde
}

type UpdateDatasetOutput struct {

	// The name of the dataset that you updated.
	//
	// This member is required.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDatasetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDataset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDataset{}, middleware.After)
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
	if err = addOpUpdateDatasetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDataset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "databrew",
		OperationName: "UpdateDataset",
	}
}
