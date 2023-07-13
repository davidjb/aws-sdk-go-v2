// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Return a list of App Runner VPC Ingress Connections in your Amazon Web Services
// account.
func (c *Client) ListVpcIngressConnections(ctx context.Context, params *ListVpcIngressConnectionsInput, optFns ...func(*Options)) (*ListVpcIngressConnectionsOutput, error) {
	if params == nil {
		params = &ListVpcIngressConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVpcIngressConnections", params, optFns, c.addOperationListVpcIngressConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVpcIngressConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVpcIngressConnectionsInput struct {

	// The VPC Ingress Connections to be listed based on either the Service Arn or Vpc
	// Endpoint Id, or both.
	Filter *types.ListVpcIngressConnectionsFilter

	// The maximum number of results to include in each response (result page). It's
	// used for a paginated request. If you don't specify MaxResults , the request
	// retrieves all available results in a single response.
	MaxResults *int32

	// A token from a previous result page. It's used for a paginated request. The
	// request retrieves the next result page. All other parameter values must be
	// identical to the ones that are specified in the initial request. If you don't
	// specify NextToken , the request retrieves the first result page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListVpcIngressConnectionsOutput struct {

	// A list of summary information records for VPC Ingress Connections. In a
	// paginated request, the request returns up to MaxResults records for each call.
	//
	// This member is required.
	VpcIngressConnectionSummaryList []types.VpcIngressConnectionSummary

	// The token that you can pass in a subsequent request to get the next result
	// page. It's returned in a paginated request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVpcIngressConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListVpcIngressConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListVpcIngressConnections{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVpcIngressConnections(options.Region), middleware.Before); err != nil {
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

// ListVpcIngressConnectionsAPIClient is a client that implements the
// ListVpcIngressConnections operation.
type ListVpcIngressConnectionsAPIClient interface {
	ListVpcIngressConnections(context.Context, *ListVpcIngressConnectionsInput, ...func(*Options)) (*ListVpcIngressConnectionsOutput, error)
}

var _ ListVpcIngressConnectionsAPIClient = (*Client)(nil)

// ListVpcIngressConnectionsPaginatorOptions is the paginator options for
// ListVpcIngressConnections
type ListVpcIngressConnectionsPaginatorOptions struct {
	// The maximum number of results to include in each response (result page). It's
	// used for a paginated request. If you don't specify MaxResults , the request
	// retrieves all available results in a single response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVpcIngressConnectionsPaginator is a paginator for ListVpcIngressConnections
type ListVpcIngressConnectionsPaginator struct {
	options   ListVpcIngressConnectionsPaginatorOptions
	client    ListVpcIngressConnectionsAPIClient
	params    *ListVpcIngressConnectionsInput
	nextToken *string
	firstPage bool
}

// NewListVpcIngressConnectionsPaginator returns a new
// ListVpcIngressConnectionsPaginator
func NewListVpcIngressConnectionsPaginator(client ListVpcIngressConnectionsAPIClient, params *ListVpcIngressConnectionsInput, optFns ...func(*ListVpcIngressConnectionsPaginatorOptions)) *ListVpcIngressConnectionsPaginator {
	if params == nil {
		params = &ListVpcIngressConnectionsInput{}
	}

	options := ListVpcIngressConnectionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVpcIngressConnectionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVpcIngressConnectionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListVpcIngressConnections page.
func (p *ListVpcIngressConnectionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVpcIngressConnectionsOutput, error) {
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

	result, err := p.client.ListVpcIngressConnections(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListVpcIngressConnections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apprunner",
		OperationName: "ListVpcIngressConnections",
	}
}
