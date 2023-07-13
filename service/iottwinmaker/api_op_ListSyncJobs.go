// Code generated by smithy-go-codegen DO NOT EDIT.

package iottwinmaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iottwinmaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List all SyncJobs.
func (c *Client) ListSyncJobs(ctx context.Context, params *ListSyncJobsInput, optFns ...func(*Options)) (*ListSyncJobsOutput, error) {
	if params == nil {
		params = &ListSyncJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSyncJobs", params, optFns, c.addOperationListSyncJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSyncJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSyncJobsInput struct {

	// The ID of the workspace that contains the sync job.
	//
	// This member is required.
	WorkspaceId *string

	// The maximum number of results to return at one time. The default is 50. Valid
	// Range: Minimum value of 0. Maximum value of 200.
	MaxResults *int32

	// The string that specifies the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSyncJobsOutput struct {

	// The string that specifies the next page of results.
	NextToken *string

	// The listed SyncJob summaries.
	SyncJobSummaries []types.SyncJobSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSyncJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSyncJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSyncJobs{}, middleware.After)
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
	if err = addEndpointPrefix_opListSyncJobsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListSyncJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSyncJobs(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListSyncJobsMiddleware struct {
}

func (*endpointPrefix_opListSyncJobsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListSyncJobsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListSyncJobsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListSyncJobsMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListSyncJobsAPIClient is a client that implements the ListSyncJobs operation.
type ListSyncJobsAPIClient interface {
	ListSyncJobs(context.Context, *ListSyncJobsInput, ...func(*Options)) (*ListSyncJobsOutput, error)
}

var _ ListSyncJobsAPIClient = (*Client)(nil)

// ListSyncJobsPaginatorOptions is the paginator options for ListSyncJobs
type ListSyncJobsPaginatorOptions struct {
	// The maximum number of results to return at one time. The default is 50. Valid
	// Range: Minimum value of 0. Maximum value of 200.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSyncJobsPaginator is a paginator for ListSyncJobs
type ListSyncJobsPaginator struct {
	options   ListSyncJobsPaginatorOptions
	client    ListSyncJobsAPIClient
	params    *ListSyncJobsInput
	nextToken *string
	firstPage bool
}

// NewListSyncJobsPaginator returns a new ListSyncJobsPaginator
func NewListSyncJobsPaginator(client ListSyncJobsAPIClient, params *ListSyncJobsInput, optFns ...func(*ListSyncJobsPaginatorOptions)) *ListSyncJobsPaginator {
	if params == nil {
		params = &ListSyncJobsInput{}
	}

	options := ListSyncJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSyncJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSyncJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSyncJobs page.
func (p *ListSyncJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSyncJobsOutput, error) {
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

	result, err := p.client.ListSyncJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSyncJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iottwinmaker",
		OperationName: "ListSyncJobs",
	}
}
