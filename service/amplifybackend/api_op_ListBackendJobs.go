// Code generated by smithy-go-codegen DO NOT EDIT.

package amplifybackend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amplifybackend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the jobs for the backend of an Amplify app.
func (c *Client) ListBackendJobs(ctx context.Context, params *ListBackendJobsInput, optFns ...func(*Options)) (*ListBackendJobsOutput, error) {
	if params == nil {
		params = &ListBackendJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBackendJobs", params, optFns, c.addOperationListBackendJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBackendJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request body for ListBackendJobs.
type ListBackendJobsInput struct {

	// The app ID.
	//
	// This member is required.
	AppId *string

	// The name of the backend environment.
	//
	// This member is required.
	BackendEnvironmentName *string

	// The ID for the job.
	JobId *string

	// The maximum number of results that you want in the response.
	MaxResults int32

	// The token for the next set of results.
	NextToken *string

	// Filters the list of response objects to include only those with the specified
	// operation name.
	Operation *string

	// Filters the list of response objects to include only those with the specified
	// status.
	Status *string

	noSmithyDocumentSerde
}

type ListBackendJobsOutput struct {

	// An array of jobs and their properties.
	Jobs []types.BackendJobRespObj

	// The token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBackendJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBackendJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBackendJobs{}, middleware.After)
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
	if err = addOpListBackendJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBackendJobs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListBackendJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplifybackend",
		OperationName: "ListBackendJobs",
	}
}
