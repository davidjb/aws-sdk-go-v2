// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/health/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns detailed information about one or more specified events. Information
// includes standard event data (Amazon Web Services Region, service, and so on, as
// returned by DescribeEvents (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEvents.html)
// ), a detailed event description, and possible additional metadata that depends
// upon the nature of the event. Affected entities are not included. To retrieve
// the entities, use the DescribeAffectedEntities (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedEntities.html)
// operation. If a specified event can't be retrieved, an error message is returned
// for that event. This operation supports resource-level permissions. You can use
// this operation to allow or deny access to specific Health events. For more
// information, see Resource- and action-based conditions (https://docs.aws.amazon.com/health/latest/ug/security_iam_id-based-policy-examples.html#resource-action-based-conditions)
// in the Health User Guide.
func (c *Client) DescribeEventDetails(ctx context.Context, params *DescribeEventDetailsInput, optFns ...func(*Options)) (*DescribeEventDetailsOutput, error) {
	if params == nil {
		params = &DescribeEventDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventDetails", params, optFns, c.addOperationDescribeEventDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventDetailsInput struct {

	// A list of event ARNs (unique identifiers). For example:
	// "arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-CDE456",
	// "arn:aws:health:us-west-1::event/EBS/AWS_EBS_LOST_VOLUME/AWS_EBS_LOST_VOLUME_CHI789_JKL101"
	//
	// This member is required.
	EventArns []string

	// The locale (language) to return information in. English (en) is the default and
	// the only supported value at this time.
	Locale *string

	noSmithyDocumentSerde
}

type DescribeEventDetailsOutput struct {

	// Error messages for any events that could not be retrieved.
	FailedSet []types.EventDetailsErrorItem

	// Information about the events that could be retrieved.
	SuccessfulSet []types.EventDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEventDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEventDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEventDetails{}, middleware.After)
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
	if err = addOpDescribeEventDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventDetails(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEventDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "health",
		OperationName: "DescribeEventDetails",
	}
}
