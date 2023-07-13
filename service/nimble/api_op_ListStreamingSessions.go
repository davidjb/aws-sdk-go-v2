// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the streaming sessions in a studio.
func (c *Client) ListStreamingSessions(ctx context.Context, params *ListStreamingSessionsInput, optFns ...func(*Options)) (*ListStreamingSessionsOutput, error) {
	if params == nil {
		params = &ListStreamingSessionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStreamingSessions", params, optFns, c.addOperationListStreamingSessionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStreamingSessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStreamingSessionsInput struct {

	// The studio ID.
	//
	// This member is required.
	StudioId *string

	// Filters the request to streaming sessions created by the given user.
	CreatedBy *string

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// Filters the request to streaming session owned by the given user
	OwnedBy *string

	// Filters the request to only the provided session IDs.
	SessionIds *string

	noSmithyDocumentSerde
}

type ListStreamingSessionsOutput struct {

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// A collection of streaming sessions.
	Sessions []types.StreamingSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStreamingSessionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListStreamingSessions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListStreamingSessions{}, middleware.After)
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
	if err = addOpListStreamingSessionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStreamingSessions(options.Region), middleware.Before); err != nil {
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

// ListStreamingSessionsAPIClient is a client that implements the
// ListStreamingSessions operation.
type ListStreamingSessionsAPIClient interface {
	ListStreamingSessions(context.Context, *ListStreamingSessionsInput, ...func(*Options)) (*ListStreamingSessionsOutput, error)
}

var _ ListStreamingSessionsAPIClient = (*Client)(nil)

// ListStreamingSessionsPaginatorOptions is the paginator options for
// ListStreamingSessions
type ListStreamingSessionsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStreamingSessionsPaginator is a paginator for ListStreamingSessions
type ListStreamingSessionsPaginator struct {
	options   ListStreamingSessionsPaginatorOptions
	client    ListStreamingSessionsAPIClient
	params    *ListStreamingSessionsInput
	nextToken *string
	firstPage bool
}

// NewListStreamingSessionsPaginator returns a new ListStreamingSessionsPaginator
func NewListStreamingSessionsPaginator(client ListStreamingSessionsAPIClient, params *ListStreamingSessionsInput, optFns ...func(*ListStreamingSessionsPaginatorOptions)) *ListStreamingSessionsPaginator {
	if params == nil {
		params = &ListStreamingSessionsInput{}
	}

	options := ListStreamingSessionsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStreamingSessionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStreamingSessionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStreamingSessions page.
func (p *ListStreamingSessionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStreamingSessionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListStreamingSessions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStreamingSessions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "nimble",
		OperationName: "ListStreamingSessions",
	}
}
