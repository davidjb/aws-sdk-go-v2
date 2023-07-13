// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates an existing custom medical vocabulary with new values. This operation
// overwrites all existing information with your new values; you cannot append new
// terms onto an existing custom vocabulary.
func (c *Client) UpdateMedicalVocabulary(ctx context.Context, params *UpdateMedicalVocabularyInput, optFns ...func(*Options)) (*UpdateMedicalVocabularyOutput, error) {
	if params == nil {
		params = &UpdateMedicalVocabularyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMedicalVocabulary", params, optFns, c.addOperationUpdateMedicalVocabularyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMedicalVocabularyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMedicalVocabularyInput struct {

	// The language code that represents the language of the entries in the custom
	// vocabulary you want to update. US English ( en-US ) is the only language
	// supported with Amazon Transcribe Medical.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// The Amazon S3 location of the text file that contains your custom medical
	// vocabulary. The URI must be located in the same Amazon Web Services Region as
	// the resource you're calling. Here's an example URI path:
	// s3://DOC-EXAMPLE-BUCKET/my-vocab-file.txt
	//
	// This member is required.
	VocabularyFileUri *string

	// The name of the custom medical vocabulary you want to update. Custom medical
	// vocabulary names are case sensitive.
	//
	// This member is required.
	VocabularyName *string

	noSmithyDocumentSerde
}

type UpdateMedicalVocabularyOutput struct {

	// The language code you selected for your custom medical vocabulary. US English (
	// en-US ) is the only language supported with Amazon Transcribe Medical.
	LanguageCode types.LanguageCode

	// The date and time the specified custom medical vocabulary was last updated.
	// Timestamps are in the format YYYY-MM-DD'T'HH:MM:SS.SSSSSS-UTC . For example,
	// 2022-05-04T12:32:58.761000-07:00 represents 12:32 PM UTC-7 on May 4, 2022.
	LastModifiedTime *time.Time

	// The name of the updated custom medical vocabulary.
	VocabularyName *string

	// The processing state of your custom medical vocabulary. If the state is READY ,
	// you can use the custom vocabulary in a StartMedicalTranscriptionJob request.
	VocabularyState types.VocabularyState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMedicalVocabularyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateMedicalVocabulary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMedicalVocabulary{}, middleware.After)
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
	if err = addOpUpdateMedicalVocabularyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMedicalVocabulary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateMedicalVocabulary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "UpdateMedicalVocabulary",
	}
}
