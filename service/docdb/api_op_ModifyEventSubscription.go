// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies an existing Amazon DocumentDB event notification subscription.
func (c *Client) ModifyEventSubscription(ctx context.Context, params *ModifyEventSubscriptionInput, optFns ...func(*Options)) (*ModifyEventSubscriptionOutput, error) {
	if params == nil {
		params = &ModifyEventSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyEventSubscription", params, optFns, c.addOperationModifyEventSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyEventSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to ModifyEventSubscription .
type ModifyEventSubscriptionInput struct {

	// The name of the Amazon DocumentDB event notification subscription.
	//
	// This member is required.
	SubscriptionName *string

	// A Boolean value; set to true to activate the subscription.
	Enabled *bool

	// A list of event categories for a SourceType that you want to subscribe to.
	EventCategories []string

	// The Amazon Resource Name (ARN) of the SNS topic created for event notification.
	// The ARN is created by Amazon SNS when you create a topic and subscribe to it.
	SnsTopicArn *string

	// The type of source that is generating the events. For example, if you want to
	// be notified of events generated by an instance, set this parameter to
	// db-instance . If this value is not specified, all events are returned. Valid
	// values: db-instance , db-parameter-group , db-security-group
	SourceType *string

	noSmithyDocumentSerde
}

type ModifyEventSubscriptionOutput struct {

	// Detailed information about an event to which you have subscribed.
	EventSubscription *types.EventSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyEventSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyEventSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyEventSubscription{}, middleware.After)
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
	if err = addOpModifyEventSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyEventSubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyEventSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyEventSubscription",
	}
}
