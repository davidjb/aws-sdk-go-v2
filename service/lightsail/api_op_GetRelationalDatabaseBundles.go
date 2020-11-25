// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the list of bundles that are available in Amazon Lightsail. A bundle
// describes the performance specifications for a database. You can use a bundle ID
// to create a new database with explicit performance specifications.
func (c *Client) GetRelationalDatabaseBundles(ctx context.Context, params *GetRelationalDatabaseBundlesInput, optFns ...func(*Options)) (*GetRelationalDatabaseBundlesOutput, error) {
	if params == nil {
		params = &GetRelationalDatabaseBundlesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRelationalDatabaseBundles", params, optFns, addOperationGetRelationalDatabaseBundlesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRelationalDatabaseBundlesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRelationalDatabaseBundlesInput struct {

	// The token to advance to the next page of results from your request. To get a
	// page token, perform an initial GetRelationalDatabaseBundles request. If your
	// results are paginated, the response will return a next page token that you can
	// specify as the page token in a subsequent request.
	PageToken *string
}

type GetRelationalDatabaseBundlesOutput struct {

	// An object describing the result of your get relational database bundles request.
	Bundles []types.RelationalDatabaseBundle

	// The token to advance to the next page of results from your request. A next page
	// token is not returned if there are no more results to display. To get the next
	// page of results, perform another GetRelationalDatabaseBundles request and
	// specify the next page token using the pageToken parameter.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRelationalDatabaseBundlesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRelationalDatabaseBundles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRelationalDatabaseBundles{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRelationalDatabaseBundles(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRelationalDatabaseBundles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetRelationalDatabaseBundles",
	}
}
