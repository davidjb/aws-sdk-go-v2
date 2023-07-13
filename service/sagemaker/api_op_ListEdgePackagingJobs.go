// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a list of edge packaging jobs.
func (c *Client) ListEdgePackagingJobs(ctx context.Context, params *ListEdgePackagingJobsInput, optFns ...func(*Options)) (*ListEdgePackagingJobsOutput, error) {
	if params == nil {
		params = &ListEdgePackagingJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEdgePackagingJobs", params, optFns, c.addOperationListEdgePackagingJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEdgePackagingJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEdgePackagingJobsInput struct {

	// Select jobs where the job was created after specified time.
	CreationTimeAfter *time.Time

	// Select jobs where the job was created before specified time.
	CreationTimeBefore *time.Time

	// Select jobs where the job was updated after specified time.
	LastModifiedTimeAfter *time.Time

	// Select jobs where the job was updated before specified time.
	LastModifiedTimeBefore *time.Time

	// Maximum number of results to select.
	MaxResults *int32

	// Filter for jobs where the model name contains this string.
	ModelNameContains *string

	// Filter for jobs containing this name in their packaging job name.
	NameContains *string

	// The response from the last list when returning a list large enough to need
	// tokening.
	NextToken *string

	// Use to specify what column to sort by.
	SortBy types.ListEdgePackagingJobsSortBy

	// What direction to sort by.
	SortOrder types.SortOrder

	// The job status to filter for.
	StatusEquals types.EdgePackagingJobStatus

	noSmithyDocumentSerde
}

type ListEdgePackagingJobsOutput struct {

	// Summaries of edge packaging jobs.
	//
	// This member is required.
	EdgePackagingJobSummaries []types.EdgePackagingJobSummary

	// Token to use when calling the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEdgePackagingJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEdgePackagingJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEdgePackagingJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEdgePackagingJobs(options.Region), middleware.Before); err != nil {
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

// ListEdgePackagingJobsAPIClient is a client that implements the
// ListEdgePackagingJobs operation.
type ListEdgePackagingJobsAPIClient interface {
	ListEdgePackagingJobs(context.Context, *ListEdgePackagingJobsInput, ...func(*Options)) (*ListEdgePackagingJobsOutput, error)
}

var _ ListEdgePackagingJobsAPIClient = (*Client)(nil)

// ListEdgePackagingJobsPaginatorOptions is the paginator options for
// ListEdgePackagingJobs
type ListEdgePackagingJobsPaginatorOptions struct {
	// Maximum number of results to select.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEdgePackagingJobsPaginator is a paginator for ListEdgePackagingJobs
type ListEdgePackagingJobsPaginator struct {
	options   ListEdgePackagingJobsPaginatorOptions
	client    ListEdgePackagingJobsAPIClient
	params    *ListEdgePackagingJobsInput
	nextToken *string
	firstPage bool
}

// NewListEdgePackagingJobsPaginator returns a new ListEdgePackagingJobsPaginator
func NewListEdgePackagingJobsPaginator(client ListEdgePackagingJobsAPIClient, params *ListEdgePackagingJobsInput, optFns ...func(*ListEdgePackagingJobsPaginatorOptions)) *ListEdgePackagingJobsPaginator {
	if params == nil {
		params = &ListEdgePackagingJobsInput{}
	}

	options := ListEdgePackagingJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEdgePackagingJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEdgePackagingJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEdgePackagingJobs page.
func (p *ListEdgePackagingJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEdgePackagingJobsOutput, error) {
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

	result, err := p.client.ListEdgePackagingJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEdgePackagingJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListEdgePackagingJobs",
	}
}
