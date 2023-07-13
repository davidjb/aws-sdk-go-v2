// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified origination phone number, or all the phone numbers in
// your account. If you specify phone number IDs, the output includes information
// for only the specified phone numbers. If you specify filters, the output
// includes information for only those phone numbers that meet the filter criteria.
// If you don't specify phone number IDs or filters, the output includes
// information for all phone numbers. If you specify a phone number ID that isn't
// valid, an Error is returned.
func (c *Client) DescribePhoneNumbers(ctx context.Context, params *DescribePhoneNumbersInput, optFns ...func(*Options)) (*DescribePhoneNumbersOutput, error) {
	if params == nil {
		params = &DescribePhoneNumbersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePhoneNumbers", params, optFns, c.addOperationDescribePhoneNumbersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePhoneNumbersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePhoneNumbersInput struct {

	// An array of PhoneNumberFilter objects to filter the results.
	Filters []types.PhoneNumberFilter

	// The maximum number of results to return per each request.
	MaxResults *int32

	// The token to be used for the next set of paginated results. You don't need to
	// supply a value for this field in the initial request.
	NextToken *string

	// The unique identifier of phone numbers to find information about. This is an
	// array of strings that can be either the PhoneNumberId or PhoneNumberArn.
	PhoneNumberIds []string

	noSmithyDocumentSerde
}

type DescribePhoneNumbersOutput struct {

	// The token to be used for the next set of paginated results. If this field is
	// empty then there are no more results.
	NextToken *string

	// An array of PhoneNumberInformation objects that contain the details for the
	// requested phone numbers.
	PhoneNumbers []types.PhoneNumberInformation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePhoneNumbersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribePhoneNumbers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribePhoneNumbers{}, middleware.After)
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
	if err = addOpDescribePhoneNumbersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePhoneNumbers(options.Region), middleware.Before); err != nil {
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

// DescribePhoneNumbersAPIClient is a client that implements the
// DescribePhoneNumbers operation.
type DescribePhoneNumbersAPIClient interface {
	DescribePhoneNumbers(context.Context, *DescribePhoneNumbersInput, ...func(*Options)) (*DescribePhoneNumbersOutput, error)
}

var _ DescribePhoneNumbersAPIClient = (*Client)(nil)

// DescribePhoneNumbersPaginatorOptions is the paginator options for
// DescribePhoneNumbers
type DescribePhoneNumbersPaginatorOptions struct {
	// The maximum number of results to return per each request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribePhoneNumbersPaginator is a paginator for DescribePhoneNumbers
type DescribePhoneNumbersPaginator struct {
	options   DescribePhoneNumbersPaginatorOptions
	client    DescribePhoneNumbersAPIClient
	params    *DescribePhoneNumbersInput
	nextToken *string
	firstPage bool
}

// NewDescribePhoneNumbersPaginator returns a new DescribePhoneNumbersPaginator
func NewDescribePhoneNumbersPaginator(client DescribePhoneNumbersAPIClient, params *DescribePhoneNumbersInput, optFns ...func(*DescribePhoneNumbersPaginatorOptions)) *DescribePhoneNumbersPaginator {
	if params == nil {
		params = &DescribePhoneNumbersInput{}
	}

	options := DescribePhoneNumbersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribePhoneNumbersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribePhoneNumbersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribePhoneNumbers page.
func (p *DescribePhoneNumbersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribePhoneNumbersOutput, error) {
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

	result, err := p.client.DescribePhoneNumbers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribePhoneNumbers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "DescribePhoneNumbers",
	}
}
