// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an empty ipynb file in the specified Apache Spark enabled workgroup.
// Throws an error if a file in the workgroup with the same name already exists.
func (c *Client) CreateNotebook(ctx context.Context, params *CreateNotebookInput, optFns ...func(*Options)) (*CreateNotebookOutput, error) {
	if params == nil {
		params = &CreateNotebookInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNotebook", params, optFns, c.addOperationCreateNotebookMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNotebookOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNotebookInput struct {

	// The name of the ipynb file to be created in the Spark workgroup, without the
	// .ipynb extension.
	//
	// This member is required.
	Name *string

	// The name of the Spark enabled workgroup in which the notebook will be created.
	//
	// This member is required.
	WorkGroup *string

	// A unique case-sensitive string used to ensure the request to create the
	// notebook is idempotent (executes only once). This token is listed as not
	// required because Amazon Web Services SDKs (for example the Amazon Web Services
	// SDK for Java) auto-generate the token for you. If you are not using the Amazon
	// Web Services SDK or the Amazon Web Services CLI, you must provide this token or
	// the action will fail.
	ClientRequestToken *string

	noSmithyDocumentSerde
}

type CreateNotebookOutput struct {

	// A unique identifier for the notebook.
	NotebookId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateNotebookMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateNotebook{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateNotebook{}, middleware.After)
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
	if err = addOpCreateNotebookValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNotebook(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateNotebook(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "CreateNotebook",
	}
}
