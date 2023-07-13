// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the status of Service Catalog in SageMaker. Service Catalog is used to
// create SageMaker projects.
func (c *Client) GetSagemakerServicecatalogPortfolioStatus(ctx context.Context, params *GetSagemakerServicecatalogPortfolioStatusInput, optFns ...func(*Options)) (*GetSagemakerServicecatalogPortfolioStatusOutput, error) {
	if params == nil {
		params = &GetSagemakerServicecatalogPortfolioStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSagemakerServicecatalogPortfolioStatus", params, optFns, c.addOperationGetSagemakerServicecatalogPortfolioStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSagemakerServicecatalogPortfolioStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSagemakerServicecatalogPortfolioStatusInput struct {
	noSmithyDocumentSerde
}

type GetSagemakerServicecatalogPortfolioStatusOutput struct {

	// Whether Service Catalog is enabled or disabled in SageMaker.
	Status types.SagemakerServicecatalogStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSagemakerServicecatalogPortfolioStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSagemakerServicecatalogPortfolioStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSagemakerServicecatalogPortfolioStatus{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSagemakerServicecatalogPortfolioStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSagemakerServicecatalogPortfolioStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "GetSagemakerServicecatalogPortfolioStatus",
	}
}
