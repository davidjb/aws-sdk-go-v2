// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of PackageSummary
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageSummary.html)
// objects for packages in a repository that match the request parameters.
func (c *Client) ListPackages(ctx context.Context, params *ListPackagesInput, optFns ...func(*Options)) (*ListPackagesOutput, error) {
	if params == nil {
		params = &ListPackagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPackages", params, optFns, c.addOperationListPackagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPackagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPackagesInput struct {

	// The name of the domain that contains the repository that contains the requested
	// list of packages.
	//
	// This member is required.
	Domain *string

	// The name of the repository from which packages are to be listed.
	//
	// This member is required.
	Repository *string

	// The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	// The format of the packages.
	Format types.PackageFormat

	// The maximum number of results to return per page.
	MaxResults *int32

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	// * The namespace of a Maven package is its
	// groupId.
	//
	// * The namespace of an npm package is its scope.
	//
	// * A Python package
	// does not contain a corresponding component, so Python packages do not have a
	// namespace.
	Namespace *string

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// A prefix used to filter returned packages. Only packages with names that start
	// with packagePrefix are returned.
	PackagePrefix *string

	noSmithyDocumentSerde
}

type ListPackagesOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// The list of returned PackageSummary
	// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageSummary.html)
	// objects.
	Packages []types.PackageSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPackagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPackages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPackages{}, middleware.After)
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListPackagesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPackages(options.Region), middleware.Before); err != nil {
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

// ListPackagesAPIClient is a client that implements the ListPackages operation.
type ListPackagesAPIClient interface {
	ListPackages(context.Context, *ListPackagesInput, ...func(*Options)) (*ListPackagesOutput, error)
}

var _ ListPackagesAPIClient = (*Client)(nil)

// ListPackagesPaginatorOptions is the paginator options for ListPackages
type ListPackagesPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPackagesPaginator is a paginator for ListPackages
type ListPackagesPaginator struct {
	options   ListPackagesPaginatorOptions
	client    ListPackagesAPIClient
	params    *ListPackagesInput
	nextToken *string
	firstPage bool
}

// NewListPackagesPaginator returns a new ListPackagesPaginator
func NewListPackagesPaginator(client ListPackagesAPIClient, params *ListPackagesInput, optFns ...func(*ListPackagesPaginatorOptions)) *ListPackagesPaginator {
	if params == nil {
		params = &ListPackagesInput{}
	}

	options := ListPackagesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPackagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPackagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPackages page.
func (p *ListPackagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPackagesOutput, error) {
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

	result, err := p.client.ListPackages(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPackages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "ListPackages",
	}
}
