// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the network interfaces specified in a decoder manifest. This API
// operation uses pagination. Specify the nextToken parameter in the request to
// return more results.
func (c *Client) ListDecoderManifestNetworkInterfaces(ctx context.Context, params *ListDecoderManifestNetworkInterfacesInput, optFns ...func(*Options)) (*ListDecoderManifestNetworkInterfacesOutput, error) {
	if params == nil {
		params = &ListDecoderManifestNetworkInterfacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDecoderManifestNetworkInterfaces", params, optFns, c.addOperationListDecoderManifestNetworkInterfacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDecoderManifestNetworkInterfacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDecoderManifestNetworkInterfacesInput struct {

	// The name of the decoder manifest to list information about.
	//
	// This member is required.
	Name *string

	// The maximum number of items to return, between 1 and 100, inclusive.
	MaxResults *int32

	// A pagination token for the next set of results. If the results of a search are
	// large, only a portion of the results are returned, and a nextToken pagination
	// token is returned in the response. To retrieve the next set of results, reissue
	// the search request and include the returned token. When all results have been
	// returned, the response does not contain a pagination token value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDecoderManifestNetworkInterfacesOutput struct {

	// A list of information about network interfaces.
	NetworkInterfaces []types.NetworkInterface

	// The token to retrieve the next set of results, or null if there are no more
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDecoderManifestNetworkInterfacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListDecoderManifestNetworkInterfaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListDecoderManifestNetworkInterfaces{}, middleware.After)
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
	if err = addOpListDecoderManifestNetworkInterfacesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDecoderManifestNetworkInterfaces(options.Region), middleware.Before); err != nil {
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

// ListDecoderManifestNetworkInterfacesAPIClient is a client that implements the
// ListDecoderManifestNetworkInterfaces operation.
type ListDecoderManifestNetworkInterfacesAPIClient interface {
	ListDecoderManifestNetworkInterfaces(context.Context, *ListDecoderManifestNetworkInterfacesInput, ...func(*Options)) (*ListDecoderManifestNetworkInterfacesOutput, error)
}

var _ ListDecoderManifestNetworkInterfacesAPIClient = (*Client)(nil)

// ListDecoderManifestNetworkInterfacesPaginatorOptions is the paginator options
// for ListDecoderManifestNetworkInterfaces
type ListDecoderManifestNetworkInterfacesPaginatorOptions struct {
	// The maximum number of items to return, between 1 and 100, inclusive.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDecoderManifestNetworkInterfacesPaginator is a paginator for
// ListDecoderManifestNetworkInterfaces
type ListDecoderManifestNetworkInterfacesPaginator struct {
	options   ListDecoderManifestNetworkInterfacesPaginatorOptions
	client    ListDecoderManifestNetworkInterfacesAPIClient
	params    *ListDecoderManifestNetworkInterfacesInput
	nextToken *string
	firstPage bool
}

// NewListDecoderManifestNetworkInterfacesPaginator returns a new
// ListDecoderManifestNetworkInterfacesPaginator
func NewListDecoderManifestNetworkInterfacesPaginator(client ListDecoderManifestNetworkInterfacesAPIClient, params *ListDecoderManifestNetworkInterfacesInput, optFns ...func(*ListDecoderManifestNetworkInterfacesPaginatorOptions)) *ListDecoderManifestNetworkInterfacesPaginator {
	if params == nil {
		params = &ListDecoderManifestNetworkInterfacesInput{}
	}

	options := ListDecoderManifestNetworkInterfacesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDecoderManifestNetworkInterfacesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDecoderManifestNetworkInterfacesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDecoderManifestNetworkInterfaces page.
func (p *ListDecoderManifestNetworkInterfacesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDecoderManifestNetworkInterfacesOutput, error) {
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

	result, err := p.client.ListDecoderManifestNetworkInterfaces(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDecoderManifestNetworkInterfaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotfleetwise",
		OperationName: "ListDecoderManifestNetworkInterfaces",
	}
}
