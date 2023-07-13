// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves Protection objects for the account. You can retrieve all protections
// or you can provide filtering criteria and retrieve just the subset of
// protections that match the criteria.
func (c *Client) ListProtections(ctx context.Context, params *ListProtectionsInput, optFns ...func(*Options)) (*ListProtectionsOutput, error) {
	if params == nil {
		params = &ListProtectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProtections", params, optFns, c.addOperationListProtectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProtectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProtectionsInput struct {

	// Narrows the set of protections that the call retrieves. You can retrieve a
	// single protection by providing its name or the ARN (Amazon Resource Name) of its
	// protected resource. You can also retrieve all protections for a specific
	// resource type. You can provide up to one criteria per filter type. Shield
	// Advanced returns protections that exactly match all of the filter criteria that
	// you provide.
	InclusionFilters *types.InclusionProtectionFilters

	// The greatest number of objects that you want Shield Advanced to return to the
	// list request. Shield Advanced might return fewer objects than you indicate in
	// this setting, even if more objects are available. If there are more objects
	// remaining, Shield Advanced will always also return a NextToken value in the
	// response. The default setting is 20.
	MaxResults *int32

	// When you request a list of objects from Shield Advanced, if the response does
	// not include all of the remaining available objects, Shield Advanced includes a
	// NextToken value in the response. You can retrieve the next batch of objects by
	// requesting the list again and providing the token that was returned by the prior
	// call in your request. You can indicate the maximum number of objects that you
	// want Shield Advanced to return for a single call with the MaxResults setting.
	// Shield Advanced will not return more than MaxResults objects, but may return
	// fewer, even if more objects are still available. Whenever more objects remain
	// that Shield Advanced has not yet returned to you, the response will include a
	// NextToken value. On your first call to a list operation, leave this setting
	// empty.
	NextToken *string

	noSmithyDocumentSerde
}

type ListProtectionsOutput struct {

	// When you request a list of objects from Shield Advanced, if the response does
	// not include all of the remaining available objects, Shield Advanced includes a
	// NextToken value in the response. You can retrieve the next batch of objects by
	// requesting the list again and providing the token that was returned by the prior
	// call in your request. You can indicate the maximum number of objects that you
	// want Shield Advanced to return for a single call with the MaxResults setting.
	// Shield Advanced will not return more than MaxResults objects, but may return
	// fewer, even if more objects are still available. Whenever more objects remain
	// that Shield Advanced has not yet returned to you, the response will include a
	// NextToken value.
	NextToken *string

	// The array of enabled Protection objects.
	Protections []types.Protection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProtectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListProtections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListProtections{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProtections(options.Region), middleware.Before); err != nil {
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

// ListProtectionsAPIClient is a client that implements the ListProtections
// operation.
type ListProtectionsAPIClient interface {
	ListProtections(context.Context, *ListProtectionsInput, ...func(*Options)) (*ListProtectionsOutput, error)
}

var _ ListProtectionsAPIClient = (*Client)(nil)

// ListProtectionsPaginatorOptions is the paginator options for ListProtections
type ListProtectionsPaginatorOptions struct {
	// The greatest number of objects that you want Shield Advanced to return to the
	// list request. Shield Advanced might return fewer objects than you indicate in
	// this setting, even if more objects are available. If there are more objects
	// remaining, Shield Advanced will always also return a NextToken value in the
	// response. The default setting is 20.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProtectionsPaginator is a paginator for ListProtections
type ListProtectionsPaginator struct {
	options   ListProtectionsPaginatorOptions
	client    ListProtectionsAPIClient
	params    *ListProtectionsInput
	nextToken *string
	firstPage bool
}

// NewListProtectionsPaginator returns a new ListProtectionsPaginator
func NewListProtectionsPaginator(client ListProtectionsAPIClient, params *ListProtectionsInput, optFns ...func(*ListProtectionsPaginatorOptions)) *ListProtectionsPaginator {
	if params == nil {
		params = &ListProtectionsInput{}
	}

	options := ListProtectionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProtectionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProtectionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProtections page.
func (p *ListProtectionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProtectionsOutput, error) {
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

	result, err := p.client.ListProtections(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProtections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "ListProtections",
	}
}
