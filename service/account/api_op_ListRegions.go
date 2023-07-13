// Code generated by smithy-go-codegen DO NOT EDIT.

package account

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the Regions for a given account and their respective opt-in statuses.
// Optionally, this list can be filtered by the region-opt-status-contains
// parameter.
func (c *Client) ListRegions(ctx context.Context, params *ListRegionsInput, optFns ...func(*Options)) (*ListRegionsOutput, error) {
	if params == nil {
		params = &ListRegionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRegions", params, optFns, c.addOperationListRegionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRegionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRegionsInput struct {

	// Specifies the 12-digit account ID number of the Amazon Web Services account
	// that you want to access or modify with this operation. If you don't specify this
	// parameter, it defaults to the Amazon Web Services account of the identity used
	// to call the operation. To use this parameter, the caller must be an identity in
	// the organization's management account (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account)
	// or a delegated administrator account. The specified account ID must also be a
	// member account in the same organization. The organization must have all
	// features enabled (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html)
	// , and the organization must have trusted access (https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-trusted-access.html)
	// enabled for the Account Management service, and optionally a delegated admin (https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-delegated-admin.html)
	// account assigned. The management account can't specify its own AccountId . It
	// must call the operation in standalone context by not including the AccountId
	// parameter. To call this operation on an account that is not a member of an
	// organization, don't specify this parameter. Instead, call the operation using an
	// identity belonging to the account whose contacts you wish to retrieve or modify.
	AccountId *string

	// The total number of items to return in the command’s output. If the total
	// number of items available is more than the value specified, a NextToken is
	// provided in the command’s output. To resume pagination, provide the NextToken
	// value in the starting-token argument of a subsequent command. Do not use the
	// NextToken response element directly outside of the Amazon Web Services CLI. For
	// usage examples, see Pagination (http://docs.aws.amazon.com/cli/latest/userguide/pagination.html)
	// in the Amazon Web Services Command Line Interface User Guide.
	MaxResults *int32

	// A token used to specify where to start paginating. This is the NextToken from a
	// previously truncated response. For usage examples, see Pagination (http://docs.aws.amazon.com/cli/latest/userguide/pagination.html)
	// in the Amazon Web Services Command Line Interface User Guide.
	NextToken *string

	// A list of Region statuses (Enabling, Enabled, Disabling, Disabled,
	// Enabled_by_default) to use to filter the list of Regions for a given account.
	// For example, passing in a value of ENABLING will only return a list of Regions
	// with a Region status of ENABLING.
	RegionOptStatusContains []types.RegionOptStatus

	noSmithyDocumentSerde
}

type ListRegionsOutput struct {

	// If there is more data to be returned, this will be populated. It should be
	// passed into the next-token request parameter of list-regions .
	NextToken *string

	// This is a list of Regions for a given account, or if the filtered parameter was
	// used, a list of Regions that match the filter criteria set in the filter
	// parameter.
	Regions []types.Region

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRegionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRegions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRegions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRegions(options.Region), middleware.Before); err != nil {
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

// ListRegionsAPIClient is a client that implements the ListRegions operation.
type ListRegionsAPIClient interface {
	ListRegions(context.Context, *ListRegionsInput, ...func(*Options)) (*ListRegionsOutput, error)
}

var _ ListRegionsAPIClient = (*Client)(nil)

// ListRegionsPaginatorOptions is the paginator options for ListRegions
type ListRegionsPaginatorOptions struct {
	// The total number of items to return in the command’s output. If the total
	// number of items available is more than the value specified, a NextToken is
	// provided in the command’s output. To resume pagination, provide the NextToken
	// value in the starting-token argument of a subsequent command. Do not use the
	// NextToken response element directly outside of the Amazon Web Services CLI. For
	// usage examples, see Pagination (http://docs.aws.amazon.com/cli/latest/userguide/pagination.html)
	// in the Amazon Web Services Command Line Interface User Guide.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRegionsPaginator is a paginator for ListRegions
type ListRegionsPaginator struct {
	options   ListRegionsPaginatorOptions
	client    ListRegionsAPIClient
	params    *ListRegionsInput
	nextToken *string
	firstPage bool
}

// NewListRegionsPaginator returns a new ListRegionsPaginator
func NewListRegionsPaginator(client ListRegionsAPIClient, params *ListRegionsInput, optFns ...func(*ListRegionsPaginatorOptions)) *ListRegionsPaginator {
	if params == nil {
		params = &ListRegionsInput{}
	}

	options := ListRegionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRegionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRegionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRegions page.
func (p *ListRegionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRegionsOutput, error) {
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

	result, err := p.client.ListRegions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRegions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "account",
		OperationName: "ListRegions",
	}
}
