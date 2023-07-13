// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkidentity

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkidentity/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List all AppInstanceUsers created under a single AppInstance .
func (c *Client) ListAppInstanceUsers(ctx context.Context, params *ListAppInstanceUsersInput, optFns ...func(*Options)) (*ListAppInstanceUsersOutput, error) {
	if params == nil {
		params = &ListAppInstanceUsersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAppInstanceUsers", params, optFns, c.addOperationListAppInstanceUsersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAppInstanceUsersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAppInstanceUsersInput struct {

	// The ARN of the AppInstance .
	//
	// This member is required.
	AppInstanceArn *string

	// The maximum number of requests that you want returned.
	MaxResults *int32

	// The token passed by previous API calls until all requested users are returned.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAppInstanceUsersOutput struct {

	// The ARN of the AppInstance .
	AppInstanceArn *string

	// The information for each requested AppInstanceUser .
	AppInstanceUsers []types.AppInstanceUserSummary

	// The token passed by previous API calls until all requested users are returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAppInstanceUsersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAppInstanceUsers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAppInstanceUsers{}, middleware.After)
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
	if err = addOpListAppInstanceUsersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAppInstanceUsers(options.Region), middleware.Before); err != nil {
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

// ListAppInstanceUsersAPIClient is a client that implements the
// ListAppInstanceUsers operation.
type ListAppInstanceUsersAPIClient interface {
	ListAppInstanceUsers(context.Context, *ListAppInstanceUsersInput, ...func(*Options)) (*ListAppInstanceUsersOutput, error)
}

var _ ListAppInstanceUsersAPIClient = (*Client)(nil)

// ListAppInstanceUsersPaginatorOptions is the paginator options for
// ListAppInstanceUsers
type ListAppInstanceUsersPaginatorOptions struct {
	// The maximum number of requests that you want returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAppInstanceUsersPaginator is a paginator for ListAppInstanceUsers
type ListAppInstanceUsersPaginator struct {
	options   ListAppInstanceUsersPaginatorOptions
	client    ListAppInstanceUsersAPIClient
	params    *ListAppInstanceUsersInput
	nextToken *string
	firstPage bool
}

// NewListAppInstanceUsersPaginator returns a new ListAppInstanceUsersPaginator
func NewListAppInstanceUsersPaginator(client ListAppInstanceUsersAPIClient, params *ListAppInstanceUsersInput, optFns ...func(*ListAppInstanceUsersPaginatorOptions)) *ListAppInstanceUsersPaginator {
	if params == nil {
		params = &ListAppInstanceUsersInput{}
	}

	options := ListAppInstanceUsersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAppInstanceUsersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAppInstanceUsersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAppInstanceUsers page.
func (p *ListAppInstanceUsersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAppInstanceUsersOutput, error) {
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

	result, err := p.client.ListAppInstanceUsers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAppInstanceUsers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListAppInstanceUsers",
	}
}
