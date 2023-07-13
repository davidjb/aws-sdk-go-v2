// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Used to enable finding aggregation. Must be called from the aggregation Region.
// For more details about cross-Region replication, see Configuring finding
// aggregation (https://docs.aws.amazon.com/securityhub/latest/userguide/finding-aggregation.html)
// in the Security Hub User Guide.
func (c *Client) CreateFindingAggregator(ctx context.Context, params *CreateFindingAggregatorInput, optFns ...func(*Options)) (*CreateFindingAggregatorOutput, error) {
	if params == nil {
		params = &CreateFindingAggregatorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFindingAggregator", params, optFns, c.addOperationCreateFindingAggregatorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFindingAggregatorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFindingAggregatorInput struct {

	// Indicates whether to aggregate findings from all of the available Regions in
	// the current partition. Also determines whether to automatically aggregate
	// findings from new Regions as Security Hub supports them and you opt into them.
	// The selected option also determines how to use the Regions provided in the
	// Regions list. The options are as follows:
	//   - ALL_REGIONS - Indicates to aggregate findings from all of the Regions where
	//   Security Hub is enabled. When you choose this option, Security Hub also
	//   automatically aggregates findings from new Regions as Security Hub supports them
	//   and you opt into them.
	//   - ALL_REGIONS_EXCEPT_SPECIFIED - Indicates to aggregate findings from all of
	//   the Regions where Security Hub is enabled, except for the Regions listed in the
	//   Regions parameter. When you choose this option, Security Hub also
	//   automatically aggregates findings from new Regions as Security Hub supports them
	//   and you opt into them.
	//   - SPECIFIED_REGIONS - Indicates to aggregate findings only from the Regions
	//   listed in the Regions parameter. Security Hub does not automatically aggregate
	//   findings from new Regions.
	//
	// This member is required.
	RegionLinkingMode *string

	// If RegionLinkingMode is ALL_REGIONS_EXCEPT_SPECIFIED , then this is a
	// space-separated list of Regions that do not aggregate findings to the
	// aggregation Region. If RegionLinkingMode is SPECIFIED_REGIONS , then this is a
	// space-separated list of Regions that do aggregate findings to the aggregation
	// Region.
	Regions []string

	noSmithyDocumentSerde
}

type CreateFindingAggregatorOutput struct {

	// The aggregation Region.
	FindingAggregationRegion *string

	// The ARN of the finding aggregator. You use the finding aggregator ARN to
	// retrieve details for, update, and stop finding aggregation.
	FindingAggregatorArn *string

	// Indicates whether to link all Regions, all Regions except for a list of
	// excluded Regions, or a list of included Regions.
	RegionLinkingMode *string

	// The list of excluded Regions or included Regions.
	Regions []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFindingAggregatorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateFindingAggregator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFindingAggregator{}, middleware.After)
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
	if err = addOpCreateFindingAggregatorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFindingAggregator(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFindingAggregator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "CreateFindingAggregator",
	}
}
