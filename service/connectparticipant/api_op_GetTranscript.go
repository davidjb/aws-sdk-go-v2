// Code generated by smithy-go-codegen DO NOT EDIT.

package connectparticipant

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connectparticipant/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a transcript of the session, including details about any attachments.
// For information about accessing past chat contact transcripts for a persistent
// chat, see Enable persistent chat (https://docs.aws.amazon.com/connect/latest/adminguide/chat-persistence.html)
// . ConnectionToken is used for invoking this API instead of ParticipantToken .
// The Amazon Connect Participant Service APIs do not use Signature Version 4
// authentication (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html)
// .
func (c *Client) GetTranscript(ctx context.Context, params *GetTranscriptInput, optFns ...func(*Options)) (*GetTranscriptOutput, error) {
	if params == nil {
		params = &GetTranscriptInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTranscript", params, optFns, c.addOperationGetTranscriptMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTranscriptOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTranscriptInput struct {

	// The authentication token associated with the participant's connection.
	//
	// This member is required.
	ConnectionToken *string

	// The contactId from the current contact chain for which transcript is needed.
	ContactId *string

	// The maximum number of results to return in the page. Default: 10.
	MaxResults *int32

	// The pagination token. Use the value returned previously in the next subsequent
	// request to retrieve the next set of results.
	NextToken *string

	// The direction from StartPosition from which to retrieve message. Default:
	// BACKWARD when no StartPosition is provided, FORWARD with StartPosition.
	ScanDirection types.ScanDirection

	// The sort order for the records. Default: DESCENDING.
	SortOrder types.SortKey

	// A filtering option for where to start.
	StartPosition *types.StartPosition

	noSmithyDocumentSerde
}

type GetTranscriptOutput struct {

	// The initial contact ID for the contact.
	InitialContactId *string

	// The pagination token. Use the value returned previously in the next subsequent
	// request to retrieve the next set of results.
	NextToken *string

	// The list of messages in the session.
	Transcript []types.Item

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTranscriptMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTranscript{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTranscript{}, middleware.After)
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
	if err = addOpGetTranscriptValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTranscript(options.Region), middleware.Before); err != nil {
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

// GetTranscriptAPIClient is a client that implements the GetTranscript operation.
type GetTranscriptAPIClient interface {
	GetTranscript(context.Context, *GetTranscriptInput, ...func(*Options)) (*GetTranscriptOutput, error)
}

var _ GetTranscriptAPIClient = (*Client)(nil)

// GetTranscriptPaginatorOptions is the paginator options for GetTranscript
type GetTranscriptPaginatorOptions struct {
	// The maximum number of results to return in the page. Default: 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetTranscriptPaginator is a paginator for GetTranscript
type GetTranscriptPaginator struct {
	options   GetTranscriptPaginatorOptions
	client    GetTranscriptAPIClient
	params    *GetTranscriptInput
	nextToken *string
	firstPage bool
}

// NewGetTranscriptPaginator returns a new GetTranscriptPaginator
func NewGetTranscriptPaginator(client GetTranscriptAPIClient, params *GetTranscriptInput, optFns ...func(*GetTranscriptPaginatorOptions)) *GetTranscriptPaginator {
	if params == nil {
		params = &GetTranscriptInput{}
	}

	options := GetTranscriptPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetTranscriptPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetTranscriptPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetTranscript page.
func (p *GetTranscriptPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetTranscriptOutput, error) {
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

	result, err := p.client.GetTranscript(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetTranscript(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "GetTranscript",
	}
}
