// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhuborchestrator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhuborchestrator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List AWS Migration Hub Orchestrator plugins.
func (c *Client) ListPlugins(ctx context.Context, params *ListPluginsInput, optFns ...func(*Options)) (*ListPluginsOutput, error) {
	if params == nil {
		params = &ListPluginsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPlugins", params, optFns, c.addOperationListPluginsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPluginsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPluginsInput struct {

	// The maximum number of plugins that can be returned.
	MaxResults int32

	// The pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPluginsOutput struct {

	// The pagination token.
	NextToken *string

	// Migration Hub Orchestrator plugins.
	Plugins []types.PluginSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPluginsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPlugins{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPlugins{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPlugins(options.Region), middleware.Before); err != nil {
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

// ListPluginsAPIClient is a client that implements the ListPlugins operation.
type ListPluginsAPIClient interface {
	ListPlugins(context.Context, *ListPluginsInput, ...func(*Options)) (*ListPluginsOutput, error)
}

var _ ListPluginsAPIClient = (*Client)(nil)

// ListPluginsPaginatorOptions is the paginator options for ListPlugins
type ListPluginsPaginatorOptions struct {
	// The maximum number of plugins that can be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPluginsPaginator is a paginator for ListPlugins
type ListPluginsPaginator struct {
	options   ListPluginsPaginatorOptions
	client    ListPluginsAPIClient
	params    *ListPluginsInput
	nextToken *string
	firstPage bool
}

// NewListPluginsPaginator returns a new ListPluginsPaginator
func NewListPluginsPaginator(client ListPluginsAPIClient, params *ListPluginsInput, optFns ...func(*ListPluginsPaginatorOptions)) *ListPluginsPaginator {
	if params == nil {
		params = &ListPluginsInput{}
	}

	options := ListPluginsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPluginsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPluginsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPlugins page.
func (p *ListPluginsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPluginsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListPlugins(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPlugins(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "migrationhub-orchestrator",
		OperationName: "ListPlugins",
	}
}
