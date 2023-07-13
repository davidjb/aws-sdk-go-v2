// Code generated by smithy-go-codegen DO NOT EDIT.

package rum

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rum/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a destination for CloudWatch RUM extended metrics, so that the
// specified app monitor stops sending extended metrics to that destination.
func (c *Client) DeleteRumMetricsDestination(ctx context.Context, params *DeleteRumMetricsDestinationInput, optFns ...func(*Options)) (*DeleteRumMetricsDestinationOutput, error) {
	if params == nil {
		params = &DeleteRumMetricsDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteRumMetricsDestination", params, optFns, c.addOperationDeleteRumMetricsDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteRumMetricsDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRumMetricsDestinationInput struct {

	// The name of the app monitor that is sending metrics to the destination that you
	// want to delete.
	//
	// This member is required.
	AppMonitorName *string

	// The type of destination to delete. Valid values are CloudWatch and Evidently .
	//
	// This member is required.
	Destination types.MetricDestination

	// This parameter is required if Destination is Evidently . If Destination is
	// CloudWatch , do not use this parameter. This parameter specifies the ARN of the
	// Evidently experiment that corresponds to the destination to delete.
	DestinationArn *string

	noSmithyDocumentSerde
}

type DeleteRumMetricsDestinationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteRumMetricsDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteRumMetricsDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteRumMetricsDestination{}, middleware.After)
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
	if err = addOpDeleteRumMetricsDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRumMetricsDestination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteRumMetricsDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rum",
		OperationName: "DeleteRumMetricsDestination",
	}
}
