// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists, in descending order, the executors that joined a session. Newer
// executors are listed first; older executors are listed later. The result can be
// optionally filtered by state.
func (c *Client) ListExecutors(ctx context.Context, params *ListExecutorsInput, optFns ...func(*Options)) (*ListExecutorsOutput, error) {
	if params == nil {
		params = &ListExecutorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListExecutors", params, optFns, c.addOperationListExecutorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListExecutorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExecutorsInput struct {

	// The session ID.
	//
	// This member is required.
	SessionId *string

	// A filter for a specific executor state. A description of each state follows.
	// CREATING - The executor is being started, including acquiring resources. CREATED
	// - The executor has been started. REGISTERED - The executor has been registered.
	// TERMINATING - The executor is in the process of shutting down. TERMINATED - The
	// executor is no longer running. FAILED - Due to a failure, the executor is no
	// longer running.
	ExecutorStateFilter types.ExecutorState

	// The maximum number of executors to return.
	MaxResults *int32

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListExecutorsOutput struct {

	// The session ID.
	//
	// This member is required.
	SessionId *string

	// Contains summary information about the executor.
	ExecutorsSummary []types.ExecutorsSummary

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListExecutorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListExecutors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListExecutors{}, middleware.After)
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
	if err = addOpListExecutorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListExecutors(options.Region), middleware.Before); err != nil {
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

// ListExecutorsAPIClient is a client that implements the ListExecutors operation.
type ListExecutorsAPIClient interface {
	ListExecutors(context.Context, *ListExecutorsInput, ...func(*Options)) (*ListExecutorsOutput, error)
}

var _ ListExecutorsAPIClient = (*Client)(nil)

// ListExecutorsPaginatorOptions is the paginator options for ListExecutors
type ListExecutorsPaginatorOptions struct {
	// The maximum number of executors to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListExecutorsPaginator is a paginator for ListExecutors
type ListExecutorsPaginator struct {
	options   ListExecutorsPaginatorOptions
	client    ListExecutorsAPIClient
	params    *ListExecutorsInput
	nextToken *string
	firstPage bool
}

// NewListExecutorsPaginator returns a new ListExecutorsPaginator
func NewListExecutorsPaginator(client ListExecutorsAPIClient, params *ListExecutorsInput, optFns ...func(*ListExecutorsPaginatorOptions)) *ListExecutorsPaginator {
	if params == nil {
		params = &ListExecutorsInput{}
	}

	options := ListExecutorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListExecutorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListExecutorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListExecutors page.
func (p *ListExecutorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListExecutorsOutput, error) {
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

	result, err := p.client.ListExecutors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListExecutors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "ListExecutors",
	}
}
