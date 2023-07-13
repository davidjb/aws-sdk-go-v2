// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists Amazon Inspector coverage details for a specific vulnerability.
func (c *Client) SearchVulnerabilities(ctx context.Context, params *SearchVulnerabilitiesInput, optFns ...func(*Options)) (*SearchVulnerabilitiesOutput, error) {
	if params == nil {
		params = &SearchVulnerabilitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchVulnerabilities", params, optFns, c.addOperationSearchVulnerabilitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchVulnerabilitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchVulnerabilitiesInput struct {

	// The criteria used to filter the results of a vulnerability search.
	//
	// This member is required.
	FilterCriteria *types.SearchVulnerabilitiesFilterCriteria

	// A token to use for paginating results that are returned in the response. Set
	// the value of this parameter to null for the first request to a list action. For
	// subsequent calls, use the NextToken value returned from the previous request to
	// continue listing results after the first page.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchVulnerabilitiesOutput struct {

	// Details about the listed vulnerability.
	//
	// This member is required.
	Vulnerabilities []types.Vulnerability

	// The pagination parameter to be used on the next list operation to retrieve more
	// items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchVulnerabilitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchVulnerabilities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchVulnerabilities{}, middleware.After)
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
	if err = addOpSearchVulnerabilitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchVulnerabilities(options.Region), middleware.Before); err != nil {
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

// SearchVulnerabilitiesAPIClient is a client that implements the
// SearchVulnerabilities operation.
type SearchVulnerabilitiesAPIClient interface {
	SearchVulnerabilities(context.Context, *SearchVulnerabilitiesInput, ...func(*Options)) (*SearchVulnerabilitiesOutput, error)
}

var _ SearchVulnerabilitiesAPIClient = (*Client)(nil)

// SearchVulnerabilitiesPaginatorOptions is the paginator options for
// SearchVulnerabilities
type SearchVulnerabilitiesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchVulnerabilitiesPaginator is a paginator for SearchVulnerabilities
type SearchVulnerabilitiesPaginator struct {
	options   SearchVulnerabilitiesPaginatorOptions
	client    SearchVulnerabilitiesAPIClient
	params    *SearchVulnerabilitiesInput
	nextToken *string
	firstPage bool
}

// NewSearchVulnerabilitiesPaginator returns a new SearchVulnerabilitiesPaginator
func NewSearchVulnerabilitiesPaginator(client SearchVulnerabilitiesAPIClient, params *SearchVulnerabilitiesInput, optFns ...func(*SearchVulnerabilitiesPaginatorOptions)) *SearchVulnerabilitiesPaginator {
	if params == nil {
		params = &SearchVulnerabilitiesInput{}
	}

	options := SearchVulnerabilitiesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchVulnerabilitiesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchVulnerabilitiesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchVulnerabilities page.
func (p *SearchVulnerabilitiesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchVulnerabilitiesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.SearchVulnerabilities(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchVulnerabilities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector2",
		OperationName: "SearchVulnerabilities",
	}
}
