// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// List hub content versions. Hub APIs are only callable through SageMaker Studio.
func (c *Client) ListHubContentVersions(ctx context.Context, params *ListHubContentVersionsInput, optFns ...func(*Options)) (*ListHubContentVersionsOutput, error) {
	if params == nil {
		params = &ListHubContentVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHubContentVersions", params, optFns, c.addOperationListHubContentVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHubContentVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHubContentVersionsInput struct {

	// The name of the hub content.
	//
	// This member is required.
	HubContentName *string

	// The type of hub content to list versions of.
	//
	// This member is required.
	HubContentType types.HubContentType

	// The name of the hub to list the content versions of.
	//
	// This member is required.
	HubName *string

	// Only list hub content versions that were created after the time specified.
	CreationTimeAfter *time.Time

	// Only list hub content versions that were created before the time specified.
	CreationTimeBefore *time.Time

	// The maximum number of hub content versions to list.
	MaxResults *int32

	// The upper bound of the hub content schema version.
	MaxSchemaVersion *string

	// The lower bound of the hub content versions to list.
	MinVersion *string

	// If the response to a previous ListHubContentVersions request was truncated, the
	// response includes a NextToken . To retrieve the next set of hub content
	// versions, use the token in the next request.
	NextToken *string

	// Sort hub content versions by either name or creation time.
	SortBy types.HubContentSortBy

	// Sort hub content versions by ascending or descending order.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListHubContentVersionsOutput struct {

	// The summaries of the listed hub content versions.
	//
	// This member is required.
	HubContentSummaries []types.HubContentInfo

	// If the response is truncated, SageMaker returns this token. To retrieve the
	// next set of hub content versions, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListHubContentVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListHubContentVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListHubContentVersions{}, middleware.After)
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
	if err = addOpListHubContentVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHubContentVersions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListHubContentVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListHubContentVersions",
	}
}
