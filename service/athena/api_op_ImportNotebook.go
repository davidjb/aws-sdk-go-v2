// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Imports a single ipynb file to a Spark enabled workgroup. The maximum file size
// that can be imported is 10 megabytes. If an ipynb file with the same name
// already exists in the workgroup, throws an error.
func (c *Client) ImportNotebook(ctx context.Context, params *ImportNotebookInput, optFns ...func(*Options)) (*ImportNotebookOutput, error) {
	if params == nil {
		params = &ImportNotebookInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportNotebook", params, optFns, c.addOperationImportNotebookMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportNotebookOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportNotebookInput struct {

	// The name of the notebook to import.
	//
	// This member is required.
	Name *string

	// The notebook content to be imported.
	//
	// This member is required.
	Payload *string

	// The notebook content type. Currently, the only valid type is IPYNB .
	//
	// This member is required.
	Type types.NotebookType

	// The name of the Spark enabled workgroup to import the notebook to.
	//
	// This member is required.
	WorkGroup *string

	// A unique case-sensitive string used to ensure the request to import the
	// notebook is idempotent (executes only once). This token is listed as not
	// required because Amazon Web Services SDKs (for example the Amazon Web Services
	// SDK for Java) auto-generate the token for you. If you are not using the Amazon
	// Web Services SDK or the Amazon Web Services CLI, you must provide this token or
	// the action will fail.
	ClientRequestToken *string

	noSmithyDocumentSerde
}

type ImportNotebookOutput struct {

	// The ID assigned to the imported notebook.
	NotebookId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportNotebookMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportNotebook{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportNotebook{}, middleware.After)
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
	if err = addOpImportNotebookValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportNotebook(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportNotebook(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "ImportNotebook",
	}
}
