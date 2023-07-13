// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearchserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearchserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all OpenSearch Serverless collections. For more information, see Creating
// and managing Amazon OpenSearch Serverless collections (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-manage.html)
// . Make sure to include an empty request body {} if you don't include any
// collection filters in the request.
func (c *Client) ListCollections(ctx context.Context, params *ListCollectionsInput, optFns ...func(*Options)) (*ListCollectionsOutput, error) {
	if params == nil {
		params = &ListCollectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCollections", params, optFns, c.addOperationListCollectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCollectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCollectionsInput struct {

	// List of filter names and values that you can use for requests.
	CollectionFilters *types.CollectionFilters

	// The maximum number of results to return. Default is 20. You can use nextToken
	// to get the next page of results.
	MaxResults *int32

	// If your initial ListCollections operation returns a nextToken , you can include
	// the returned nextToken in subsequent ListCollections operations, which returns
	// results in the next page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCollectionsOutput struct {

	// Details about each collection.
	CollectionSummaries []types.CollectionSummary

	// When nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCollectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListCollections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListCollections{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCollections(options.Region), middleware.Before); err != nil {
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

// ListCollectionsAPIClient is a client that implements the ListCollections
// operation.
type ListCollectionsAPIClient interface {
	ListCollections(context.Context, *ListCollectionsInput, ...func(*Options)) (*ListCollectionsOutput, error)
}

var _ ListCollectionsAPIClient = (*Client)(nil)

// ListCollectionsPaginatorOptions is the paginator options for ListCollections
type ListCollectionsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCollectionsPaginator is a paginator for ListCollections
type ListCollectionsPaginator struct {
	options   ListCollectionsPaginatorOptions
	client    ListCollectionsAPIClient
	params    *ListCollectionsInput
	nextToken *string
	firstPage bool
}

// NewListCollectionsPaginator returns a new ListCollectionsPaginator
func NewListCollectionsPaginator(client ListCollectionsAPIClient, params *ListCollectionsInput, optFns ...func(*ListCollectionsPaginatorOptions)) *ListCollectionsPaginator {
	if params == nil {
		params = &ListCollectionsInput{}
	}

	options := ListCollectionsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCollectionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCollectionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCollections page.
func (p *ListCollectionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCollectionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListCollections(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListCollections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aoss",
		OperationName: "ListCollections",
	}
}
