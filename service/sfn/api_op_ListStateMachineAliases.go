// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists aliases (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html)
// for a specified state machine ARN. Results are sorted by time, with the most
// recently created aliases listed first. To list aliases that reference a state
// machine version (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html)
// , you can specify the version ARN in the stateMachineArn parameter. If nextToken
// is returned, there are more results available. The value of nextToken is a
// unique pagination token for each page. Make the call again using the returned
// token to retrieve the next page. Keep all other arguments unchanged. Each
// pagination token expires after 24 hours. Using an expired pagination token will
// return an HTTP 400 InvalidToken error. Related operations:
//   - CreateStateMachineAlias
//   - DescribeStateMachineAlias
//   - UpdateStateMachineAlias
//   - DeleteStateMachineAlias
func (c *Client) ListStateMachineAliases(ctx context.Context, params *ListStateMachineAliasesInput, optFns ...func(*Options)) (*ListStateMachineAliasesOutput, error) {
	if params == nil {
		params = &ListStateMachineAliasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStateMachineAliases", params, optFns, c.addOperationListStateMachineAliasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStateMachineAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStateMachineAliasesInput struct {

	// The Amazon Resource Name (ARN) of the state machine for which you want to list
	// aliases. If you specify a state machine version ARN, this API returns a list of
	// aliases for that version.
	//
	// This member is required.
	StateMachineArn *string

	// The maximum number of results that are returned per call. You can use nextToken
	// to obtain further pages of results. The default is 100 and the maximum allowed
	// page size is 1000. A value of 0 uses the default. This is only an upper limit.
	// The actual number of results returned per call might be fewer than the specified
	// maximum.
	MaxResults int32

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return an HTTP 400 InvalidToken error.
	NextToken *string

	noSmithyDocumentSerde
}

type ListStateMachineAliasesOutput struct {

	// Aliases for the state machine.
	//
	// This member is required.
	StateMachineAliases []types.StateMachineAliasListItem

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return an HTTP 400 InvalidToken error.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStateMachineAliasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListStateMachineAliases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListStateMachineAliases{}, middleware.After)
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
	if err = addOpListStateMachineAliasesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStateMachineAliases(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListStateMachineAliases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "ListStateMachineAliases",
	}
}
