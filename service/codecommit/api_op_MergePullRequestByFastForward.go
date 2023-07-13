// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attempts to merge the source commit of a pull request into the specified
// destination branch for that pull request at the specified commit using the
// fast-forward merge strategy. If the merge is successful, it closes the pull
// request.
func (c *Client) MergePullRequestByFastForward(ctx context.Context, params *MergePullRequestByFastForwardInput, optFns ...func(*Options)) (*MergePullRequestByFastForwardOutput, error) {
	if params == nil {
		params = &MergePullRequestByFastForwardInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "MergePullRequestByFastForward", params, optFns, c.addOperationMergePullRequestByFastForwardMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*MergePullRequestByFastForwardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type MergePullRequestByFastForwardInput struct {

	// The system-generated ID of the pull request. To get this ID, use
	// ListPullRequests .
	//
	// This member is required.
	PullRequestId *string

	// The name of the repository where the pull request was created.
	//
	// This member is required.
	RepositoryName *string

	// The full commit ID of the original or updated commit in the pull request source
	// branch. Pass this value if you want an exception thrown if the current commit ID
	// of the tip of the source branch does not match this commit ID.
	SourceCommitId *string

	noSmithyDocumentSerde
}

type MergePullRequestByFastForwardOutput struct {

	// Information about the specified pull request, including the merge.
	PullRequest *types.PullRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationMergePullRequestByFastForwardMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpMergePullRequestByFastForward{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpMergePullRequestByFastForward{}, middleware.After)
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
	if err = addOpMergePullRequestByFastForwardValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opMergePullRequestByFastForward(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opMergePullRequestByFastForward(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "MergePullRequestByFastForward",
	}
}
