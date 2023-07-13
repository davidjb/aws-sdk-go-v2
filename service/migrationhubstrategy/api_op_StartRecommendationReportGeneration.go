// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubstrategy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubstrategy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts generating a recommendation report.
func (c *Client) StartRecommendationReportGeneration(ctx context.Context, params *StartRecommendationReportGenerationInput, optFns ...func(*Options)) (*StartRecommendationReportGenerationOutput, error) {
	if params == nil {
		params = &StartRecommendationReportGenerationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartRecommendationReportGeneration", params, optFns, c.addOperationStartRecommendationReportGenerationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartRecommendationReportGenerationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartRecommendationReportGenerationInput struct {

	// Groups the resources in the recommendation report with a unique name.
	GroupIdFilter []types.Group

	// The output format for the recommendation report file. The default format is
	// Microsoft Excel.
	OutputFormat types.OutputFormat

	noSmithyDocumentSerde
}

type StartRecommendationReportGenerationOutput struct {

	// The ID of the recommendation report generation task.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartRecommendationReportGenerationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartRecommendationReportGeneration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartRecommendationReportGeneration{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartRecommendationReportGeneration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartRecommendationReportGeneration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "migrationhub-strategy",
		OperationName: "StartRecommendationReportGeneration",
	}
}
