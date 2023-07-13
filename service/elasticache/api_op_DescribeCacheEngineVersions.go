// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the available cache engines and their versions.
func (c *Client) DescribeCacheEngineVersions(ctx context.Context, params *DescribeCacheEngineVersionsInput, optFns ...func(*Options)) (*DescribeCacheEngineVersionsOutput, error) {
	if params == nil {
		params = &DescribeCacheEngineVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCacheEngineVersions", params, optFns, c.addOperationDescribeCacheEngineVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCacheEngineVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DescribeCacheEngineVersions operation.
type DescribeCacheEngineVersionsInput struct {

	// The name of a specific cache parameter group family to return details for.
	// Valid values are: memcached1.4 | memcached1.5 | memcached1.6 | redis2.6 |
	// redis2.8 | redis3.2 | redis4.0 | redis5.0 | redis6.x | redis6.2 | redis7
	// Constraints:
	//   - Must be 1 to 255 alphanumeric characters
	//   - First character must be a letter
	//   - Cannot end with a hyphen or contain two consecutive hyphens
	CacheParameterGroupFamily *string

	// If true , specifies that only the default version of the specified engine or
	// engine and major version combination is to be returned.
	DefaultOnly bool

	// The cache engine to return. Valid values: memcached | redis
	Engine *string

	// The cache engine version to return. Example: 1.4.14
	EngineVersion *string

	// An optional marker returned from a prior request. Use this marker for
	// pagination of results from this operation. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved. Default: 100 Constraints: minimum
	// 20; maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

// Represents the output of a DescribeCacheEngineVersions operation.
type DescribeCacheEngineVersionsOutput struct {

	// A list of cache engine version details. Each element in the list contains
	// detailed information about one cache engine version.
	CacheEngineVersions []types.CacheEngineVersion

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCacheEngineVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeCacheEngineVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeCacheEngineVersions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCacheEngineVersions(options.Region), middleware.Before); err != nil {
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

// DescribeCacheEngineVersionsAPIClient is a client that implements the
// DescribeCacheEngineVersions operation.
type DescribeCacheEngineVersionsAPIClient interface {
	DescribeCacheEngineVersions(context.Context, *DescribeCacheEngineVersionsInput, ...func(*Options)) (*DescribeCacheEngineVersionsOutput, error)
}

var _ DescribeCacheEngineVersionsAPIClient = (*Client)(nil)

// DescribeCacheEngineVersionsPaginatorOptions is the paginator options for
// DescribeCacheEngineVersions
type DescribeCacheEngineVersionsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved. Default: 100 Constraints: minimum
	// 20; maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeCacheEngineVersionsPaginator is a paginator for
// DescribeCacheEngineVersions
type DescribeCacheEngineVersionsPaginator struct {
	options   DescribeCacheEngineVersionsPaginatorOptions
	client    DescribeCacheEngineVersionsAPIClient
	params    *DescribeCacheEngineVersionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeCacheEngineVersionsPaginator returns a new
// DescribeCacheEngineVersionsPaginator
func NewDescribeCacheEngineVersionsPaginator(client DescribeCacheEngineVersionsAPIClient, params *DescribeCacheEngineVersionsInput, optFns ...func(*DescribeCacheEngineVersionsPaginatorOptions)) *DescribeCacheEngineVersionsPaginator {
	if params == nil {
		params = &DescribeCacheEngineVersionsInput{}
	}

	options := DescribeCacheEngineVersionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeCacheEngineVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeCacheEngineVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeCacheEngineVersions page.
func (p *DescribeCacheEngineVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeCacheEngineVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeCacheEngineVersions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeCacheEngineVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DescribeCacheEngineVersions",
	}
}
