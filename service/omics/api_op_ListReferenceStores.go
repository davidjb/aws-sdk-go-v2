// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of reference stores.
func (c *Client) ListReferenceStores(ctx context.Context, params *ListReferenceStoresInput, optFns ...func(*Options)) (*ListReferenceStoresOutput, error) {
	if params == nil {
		params = &ListReferenceStoresInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReferenceStores", params, optFns, c.addOperationListReferenceStoresMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReferenceStoresOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReferenceStoresInput struct {

	// A filter to apply to the list.
	Filter *types.ReferenceStoreFilter

	// The maximum number of stores to return in one page of results.
	MaxResults *int32

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListReferenceStoresOutput struct {

	// A list of reference stores.
	//
	// This member is required.
	ReferenceStores []types.ReferenceStoreDetail

	// A pagination token that's included if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListReferenceStoresMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListReferenceStores{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListReferenceStores{}, middleware.After)
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
	if err = addEndpointPrefix_opListReferenceStoresMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReferenceStores(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListReferenceStoresMiddleware struct {
}

func (*endpointPrefix_opListReferenceStoresMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListReferenceStoresMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "control-storage-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListReferenceStoresMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListReferenceStoresMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListReferenceStoresAPIClient is a client that implements the
// ListReferenceStores operation.
type ListReferenceStoresAPIClient interface {
	ListReferenceStores(context.Context, *ListReferenceStoresInput, ...func(*Options)) (*ListReferenceStoresOutput, error)
}

var _ ListReferenceStoresAPIClient = (*Client)(nil)

// ListReferenceStoresPaginatorOptions is the paginator options for
// ListReferenceStores
type ListReferenceStoresPaginatorOptions struct {
	// The maximum number of stores to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListReferenceStoresPaginator is a paginator for ListReferenceStores
type ListReferenceStoresPaginator struct {
	options   ListReferenceStoresPaginatorOptions
	client    ListReferenceStoresAPIClient
	params    *ListReferenceStoresInput
	nextToken *string
	firstPage bool
}

// NewListReferenceStoresPaginator returns a new ListReferenceStoresPaginator
func NewListReferenceStoresPaginator(client ListReferenceStoresAPIClient, params *ListReferenceStoresInput, optFns ...func(*ListReferenceStoresPaginatorOptions)) *ListReferenceStoresPaginator {
	if params == nil {
		params = &ListReferenceStoresInput{}
	}

	options := ListReferenceStoresPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListReferenceStoresPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListReferenceStoresPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListReferenceStores page.
func (p *ListReferenceStoresPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListReferenceStoresOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListReferenceStores(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListReferenceStores(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "omics",
		OperationName: "ListReferenceStores",
	}
}
