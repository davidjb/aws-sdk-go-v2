// Code generated by smithy-go-codegen DO NOT EDIT.

package personalizeruntime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalizeruntime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Re-ranks a list of recommended items for the given user. The first item in the
// list is deemed the most likely item to be of interest to the user. The solution
// backing the campaign must have been created using a recipe of type
// PERSONALIZED_RANKING.
func (c *Client) GetPersonalizedRanking(ctx context.Context, params *GetPersonalizedRankingInput, optFns ...func(*Options)) (*GetPersonalizedRankingOutput, error) {
	if params == nil {
		params = &GetPersonalizedRankingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPersonalizedRanking", params, optFns, c.addOperationGetPersonalizedRankingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPersonalizedRankingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPersonalizedRankingInput struct {

	// The Amazon Resource Name (ARN) of the campaign to use for generating the
	// personalized ranking.
	//
	// This member is required.
	CampaignArn *string

	// A list of items (by itemId ) to rank. If an item was not included in the
	// training dataset, the item is appended to the end of the reranked list. The
	// maximum is 500.
	//
	// This member is required.
	InputList []string

	// The user for which you want the campaign to provide a personalized ranking.
	//
	// This member is required.
	UserId *string

	// The contextual metadata to use when getting recommendations. Contextual
	// metadata includes any interaction information that might be relevant when
	// getting a user's recommendations, such as the user's current location or device
	// type.
	Context map[string]string

	// The Amazon Resource Name (ARN) of a filter you created to include items or
	// exclude items from recommendations for a given user. For more information, see
	// Filtering Recommendations (https://docs.aws.amazon.com/personalize/latest/dg/filter.html)
	// .
	FilterArn *string

	// The values to use when filtering recommendations. For each placeholder
	// parameter in your filter expression, provide the parameter name (in matching
	// case) as a key and the filter value(s) as the corresponding value. Separate
	// multiple values for one parameter with a comma. For filter expressions that use
	// an INCLUDE element to include items, you must provide values for all parameters
	// that are defined in the expression. For filters with expressions that use an
	// EXCLUDE element to exclude items, you can omit the filter-values .In this case,
	// Amazon Personalize doesn't use that portion of the expression to filter
	// recommendations. For more information, see Filtering Recommendations (https://docs.aws.amazon.com/personalize/latest/dg/filter.html)
	// .
	FilterValues map[string]string

	noSmithyDocumentSerde
}

type GetPersonalizedRankingOutput struct {

	// A list of items in order of most likely interest to the user. The maximum is
	// 500.
	PersonalizedRanking []types.PredictedItem

	// The ID of the recommendation.
	RecommendationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPersonalizedRankingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPersonalizedRanking{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPersonalizedRanking{}, middleware.After)
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
	if err = addOpGetPersonalizedRankingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPersonalizedRanking(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPersonalizedRanking(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "GetPersonalizedRanking",
	}
}
