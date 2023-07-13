// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the alarms for the specified metric. To filter the results, specify a
// statistic, period, or unit. This operation retrieves only standard alarms that
// are based on the specified metric. It does not return alarms based on math
// expressions that use the specified metric, or composite alarms that use the
// specified metric.
func (c *Client) DescribeAlarmsForMetric(ctx context.Context, params *DescribeAlarmsForMetricInput, optFns ...func(*Options)) (*DescribeAlarmsForMetricOutput, error) {
	if params == nil {
		params = &DescribeAlarmsForMetricInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAlarmsForMetric", params, optFns, c.addOperationDescribeAlarmsForMetricMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAlarmsForMetricOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAlarmsForMetricInput struct {

	// The name of the metric.
	//
	// This member is required.
	MetricName *string

	// The namespace of the metric.
	//
	// This member is required.
	Namespace *string

	// The dimensions associated with the metric. If the metric has any associated
	// dimensions, you must specify them in order for the call to succeed.
	Dimensions []types.Dimension

	// The percentile statistic for the metric. Specify a value between p0.0 and p100.
	ExtendedStatistic *string

	// The period, in seconds, over which the statistic is applied.
	Period *int32

	// The statistic for the metric, other than percentiles. For percentile
	// statistics, use ExtendedStatistics .
	Statistic types.Statistic

	// The unit for the metric.
	Unit types.StandardUnit

	noSmithyDocumentSerde
}

type DescribeAlarmsForMetricOutput struct {

	// The information for each alarm with the specified metric.
	MetricAlarms []types.MetricAlarm

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAlarmsForMetricMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeAlarmsForMetric{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeAlarmsForMetric{}, middleware.After)
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
	if err = addOpDescribeAlarmsForMetricValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAlarmsForMetric(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAlarmsForMetric(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "DescribeAlarmsForMetric",
	}
}
