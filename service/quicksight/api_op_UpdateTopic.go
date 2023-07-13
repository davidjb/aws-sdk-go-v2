// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a topic.
func (c *Client) UpdateTopic(ctx context.Context, params *UpdateTopicInput, optFns ...func(*Options)) (*UpdateTopicOutput, error) {
	if params == nil {
		params = &UpdateTopicInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTopic", params, optFns, c.addOperationUpdateTopicMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTopicOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTopicInput struct {

	// The ID of the Amazon Web Services account that contains the topic that you want
	// to update.
	//
	// This member is required.
	AwsAccountId *string

	// The definition of the topic that you want to update.
	//
	// This member is required.
	Topic *types.TopicDetails

	// The ID of the topic that you want to modify. This ID is unique per Amazon Web
	// Services Region for each Amazon Web Services account.
	//
	// This member is required.
	TopicId *string

	noSmithyDocumentSerde
}

type UpdateTopicOutput struct {

	// The Amazon Resource Name (ARN) of the topic.
	Arn *string

	// The Amazon Resource Name (ARN) of the topic refresh.
	RefreshArn *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// The ID of the topic that you want to modify. This ID is unique per Amazon Web
	// Services Region for each Amazon Web Services account.
	TopicId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTopicMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateTopic{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateTopic{}, middleware.After)
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
	if err = addOpUpdateTopicValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTopic(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTopic(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateTopic",
	}
}
