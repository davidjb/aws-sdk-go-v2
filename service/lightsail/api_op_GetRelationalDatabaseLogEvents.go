// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a list of log events for a database in Amazon Lightsail.
func (c *Client) GetRelationalDatabaseLogEvents(ctx context.Context, params *GetRelationalDatabaseLogEventsInput, optFns ...func(*Options)) (*GetRelationalDatabaseLogEventsOutput, error) {
	if params == nil {
		params = &GetRelationalDatabaseLogEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRelationalDatabaseLogEvents", params, optFns, c.addOperationGetRelationalDatabaseLogEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRelationalDatabaseLogEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRelationalDatabaseLogEventsInput struct {

	// The name of the log stream. Use the get relational database log streams
	// operation to get a list of available log streams.
	//
	// This member is required.
	LogStreamName *string

	// The name of your database for which to get log events.
	//
	// This member is required.
	RelationalDatabaseName *string

	// The end of the time interval from which to get log events. Constraints:
	//   - Specified in Coordinated Universal Time (UTC).
	//   - Specified in the Unix time format. For example, if you wish to use an end
	//   time of October 1, 2018, at 8 PM UTC, then you input 1538424000 as the end
	//   time.
	EndTime *time.Time

	// The token to advance to the next or previous page of results from your request.
	// To get a page token, perform an initial GetRelationalDatabaseLogEvents request.
	// If your results are paginated, the response will return a next forward token
	// and/or next backward token that you can specify as the page token in a
	// subsequent request.
	PageToken *string

	// Parameter to specify if the log should start from head or tail. If true is
	// specified, the log event starts from the head of the log. If false is
	// specified, the log event starts from the tail of the log. For PostgreSQL, the
	// default value of false is the only option available.
	StartFromHead *bool

	// The start of the time interval from which to get log events. Constraints:
	//   - Specified in Coordinated Universal Time (UTC).
	//   - Specified in the Unix time format. For example, if you wish to use a start
	//   time of October 1, 2018, at 8 PM UTC, then you input 1538424000 as the start
	//   time.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type GetRelationalDatabaseLogEventsOutput struct {

	// A token used for advancing to the previous page of results from your get
	// relational database log events request.
	NextBackwardToken *string

	// A token used for advancing to the next page of results from your get relational
	// database log events request.
	NextForwardToken *string

	// An object describing the result of your get relational database log events
	// request.
	ResourceLogEvents []types.LogEvent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRelationalDatabaseLogEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRelationalDatabaseLogEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRelationalDatabaseLogEvents{}, middleware.After)
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
	if err = addOpGetRelationalDatabaseLogEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRelationalDatabaseLogEvents(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRelationalDatabaseLogEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetRelationalDatabaseLogEvents",
	}
}
