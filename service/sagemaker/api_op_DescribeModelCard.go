// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the content, creation time, and security configuration of an Amazon
// SageMaker Model Card.
func (c *Client) DescribeModelCard(ctx context.Context, params *DescribeModelCardInput, optFns ...func(*Options)) (*DescribeModelCardOutput, error) {
	if params == nil {
		params = &DescribeModelCardInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeModelCard", params, optFns, c.addOperationDescribeModelCardMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeModelCardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeModelCardInput struct {

	// The name of the model card to describe.
	//
	// This member is required.
	ModelCardName *string

	// The version of the model card to describe. If a version is not provided, then
	// the latest version of the model card is described.
	ModelCardVersion int32

	noSmithyDocumentSerde
}

type DescribeModelCardOutput struct {

	// The content of the model card.
	//
	// This member is required.
	Content *string

	// Information about the user who created or modified an experiment, trial, trial
	// component, lineage group, project, or model card.
	//
	// This member is required.
	CreatedBy *types.UserContext

	// The date and time the model card was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the model card.
	//
	// This member is required.
	ModelCardArn *string

	// The name of the model card.
	//
	// This member is required.
	ModelCardName *string

	// The approval status of the model card within your organization. Different
	// organizations might have different criteria for model card review and approval.
	//   - Draft : The model card is a work in progress.
	//   - PendingReview : The model card is pending review.
	//   - Approved : The model card is approved.
	//   - Archived : The model card is archived. No more updates should be made to the
	//   model card, but it can still be exported.
	//
	// This member is required.
	ModelCardStatus types.ModelCardStatus

	// The version of the model card.
	//
	// This member is required.
	ModelCardVersion int32

	// Information about the user who created or modified an experiment, trial, trial
	// component, lineage group, project, or model card.
	LastModifiedBy *types.UserContext

	// The date and time the model card was last modified.
	LastModifiedTime *time.Time

	// The processing status of model card deletion. The ModelCardProcessingStatus
	// updates throughout the different deletion steps.
	//   - DeletePending : Model card deletion request received.
	//   - DeleteInProgress : Model card deletion is in progress.
	//   - ContentDeleted : Deleted model card content.
	//   - ExportJobsDeleted : Deleted all export jobs associated with the model card.
	//   - DeleteCompleted : Successfully deleted the model card.
	//   - DeleteFailed : The model card failed to delete.
	ModelCardProcessingStatus types.ModelCardProcessingStatus

	// The security configuration used to protect model card content.
	SecurityConfig *types.ModelCardSecurityConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeModelCardMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeModelCard{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeModelCard{}, middleware.After)
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
	if err = addOpDescribeModelCardValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeModelCard(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeModelCard(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeModelCard",
	}
}
