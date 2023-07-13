// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or overwrites one or more tags for the specified SageMaker resource. You
// can add tags to notebook instances, training jobs, hyperparameter tuning jobs,
// batch transform jobs, models, labeling jobs, work teams, endpoint
// configurations, and endpoints. Each tag consists of a key and an optional value.
// Tag keys must be unique per resource. For more information about tags, see For
// more information, see Amazon Web Services Tagging Strategies (https://aws.amazon.com/answers/account-management/aws-tagging-strategies/)
// . Tags that you add to a hyperparameter tuning job by calling this API are also
// added to any training jobs that the hyperparameter tuning job launches after you
// call this API, but not to training jobs that the hyperparameter tuning job
// launched before you called this API. To make sure that the tags associated with
// a hyperparameter tuning job are also added to all training jobs that the
// hyperparameter tuning job launches, add the tags when you first create the
// tuning job by specifying them in the Tags parameter of
// CreateHyperParameterTuningJob (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateHyperParameterTuningJob.html)
// Tags that you add to a SageMaker Studio Domain or User Profile by calling this
// API are also added to any Apps that the Domain or User Profile launches after
// you call this API, but not to Apps that the Domain or User Profile launched
// before you called this API. To make sure that the tags associated with a Domain
// or User Profile are also added to all Apps that the Domain or User Profile
// launches, add the tags when you first create the Domain or User Profile by
// specifying them in the Tags parameter of CreateDomain (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateDomain.html)
// or CreateUserProfile (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateUserProfile.html)
// .
func (c *Client) AddTags(ctx context.Context, params *AddTagsInput, optFns ...func(*Options)) (*AddTagsOutput, error) {
	if params == nil {
		params = &AddTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddTags", params, optFns, c.addOperationAddTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddTagsInput struct {

	// The Amazon Resource Name (ARN) of the resource that you want to tag.
	//
	// This member is required.
	ResourceArn *string

	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, for example, by purpose, owner, or
	// environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// .
	//
	// This member is required.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type AddTagsOutput struct {

	// A list of tags associated with the SageMaker resource.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddTags{}, middleware.After)
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
	if err = addOpAddTagsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddTags(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "AddTags",
	}
}
