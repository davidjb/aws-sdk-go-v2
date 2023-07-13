// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the functions that use the specified code signing configuration. You can
// use this method prior to deleting a code signing configuration, to verify that
// no functions are using it.
func (c *Client) ListFunctionsByCodeSigningConfig(ctx context.Context, params *ListFunctionsByCodeSigningConfigInput, optFns ...func(*Options)) (*ListFunctionsByCodeSigningConfigOutput, error) {
	if params == nil {
		params = &ListFunctionsByCodeSigningConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFunctionsByCodeSigningConfig", params, optFns, c.addOperationListFunctionsByCodeSigningConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFunctionsByCodeSigningConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFunctionsByCodeSigningConfigInput struct {

	// The The Amazon Resource Name (ARN) of the code signing configuration.
	//
	// This member is required.
	CodeSigningConfigArn *string

	// Specify the pagination token that's returned by a previous request to retrieve
	// the next page of results.
	Marker *string

	// Maximum number of items to return.
	MaxItems *int32

	noSmithyDocumentSerde
}

type ListFunctionsByCodeSigningConfigOutput struct {

	// The function ARNs.
	FunctionArns []string

	// The pagination token that's included if more results are available.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFunctionsByCodeSigningConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFunctionsByCodeSigningConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFunctionsByCodeSigningConfig{}, middleware.After)
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
	if err = addOpListFunctionsByCodeSigningConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFunctionsByCodeSigningConfig(options.Region), middleware.Before); err != nil {
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

// ListFunctionsByCodeSigningConfigAPIClient is a client that implements the
// ListFunctionsByCodeSigningConfig operation.
type ListFunctionsByCodeSigningConfigAPIClient interface {
	ListFunctionsByCodeSigningConfig(context.Context, *ListFunctionsByCodeSigningConfigInput, ...func(*Options)) (*ListFunctionsByCodeSigningConfigOutput, error)
}

var _ ListFunctionsByCodeSigningConfigAPIClient = (*Client)(nil)

// ListFunctionsByCodeSigningConfigPaginatorOptions is the paginator options for
// ListFunctionsByCodeSigningConfig
type ListFunctionsByCodeSigningConfigPaginatorOptions struct {
	// Maximum number of items to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFunctionsByCodeSigningConfigPaginator is a paginator for
// ListFunctionsByCodeSigningConfig
type ListFunctionsByCodeSigningConfigPaginator struct {
	options   ListFunctionsByCodeSigningConfigPaginatorOptions
	client    ListFunctionsByCodeSigningConfigAPIClient
	params    *ListFunctionsByCodeSigningConfigInput
	nextToken *string
	firstPage bool
}

// NewListFunctionsByCodeSigningConfigPaginator returns a new
// ListFunctionsByCodeSigningConfigPaginator
func NewListFunctionsByCodeSigningConfigPaginator(client ListFunctionsByCodeSigningConfigAPIClient, params *ListFunctionsByCodeSigningConfigInput, optFns ...func(*ListFunctionsByCodeSigningConfigPaginatorOptions)) *ListFunctionsByCodeSigningConfigPaginator {
	if params == nil {
		params = &ListFunctionsByCodeSigningConfigInput{}
	}

	options := ListFunctionsByCodeSigningConfigPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFunctionsByCodeSigningConfigPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFunctionsByCodeSigningConfigPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFunctionsByCodeSigningConfig page.
func (p *ListFunctionsByCodeSigningConfigPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFunctionsByCodeSigningConfigOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListFunctionsByCodeSigningConfig(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListFunctionsByCodeSigningConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "ListFunctionsByCodeSigningConfig",
	}
}
