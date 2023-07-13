// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates a scaling policy for a fleet. Scaling policies are used to
// automatically scale a fleet's hosting capacity to meet player demand. An active
// scaling policy instructs Amazon GameLift to track a fleet metric and
// automatically change the fleet's capacity when a certain threshold is reached.
// There are two types of scaling policies: target-based and rule-based. Use a
// target-based policy to quickly and efficiently manage fleet scaling; this option
// is the most commonly used. Use rule-based policies when you need to exert
// fine-grained control over auto-scaling. Fleets can have multiple scaling
// policies of each type in force at the same time; you can have one target-based
// policy, one or multiple rule-based scaling policies, or both. We recommend
// caution, however, because multiple auto-scaling policies can have unintended
// consequences. Learn more about how to work with auto-scaling in Set Up Fleet
// Automatic Scaling (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-autoscaling.html)
// . Target-based policy A target-based policy tracks a single metric:
// PercentAvailableGameSessions. This metric tells us how much of a fleet's hosting
// capacity is ready to host game sessions but is not currently in use. This is the
// fleet's buffer; it measures the additional player demand that the fleet could
// handle at current capacity. With a target-based policy, you set your ideal
// buffer size and leave it to Amazon GameLift to take whatever action is needed to
// maintain that target. For example, you might choose to maintain a 10% buffer for
// a fleet that has the capacity to host 100 simultaneous game sessions. This
// policy tells Amazon GameLift to take action whenever the fleet's available
// capacity falls below or rises above 10 game sessions. Amazon GameLift will start
// new instances or stop unused instances in order to return to the 10% buffer. To
// create or update a target-based policy, specify a fleet ID and name, and set the
// policy type to "TargetBased". Specify the metric to track
// (PercentAvailableGameSessions) and reference a TargetConfiguration object with
// your desired buffer value. Exclude all other parameters. On a successful
// request, the policy name is returned. The scaling policy is automatically in
// force as soon as it's successfully created. If the fleet's auto-scaling actions
// are temporarily suspended, the new policy will be in force once the fleet
// actions are restarted. Rule-based policy A rule-based policy tracks specified
// fleet metric, sets a threshold value, and specifies the type of action to
// initiate when triggered. With a rule-based policy, you can select from several
// available fleet metrics. Each policy specifies whether to scale up or scale down
// (and by how much), so you need one policy for each type of action. For example,
// a policy may make the following statement: "If the percentage of idle instances
// is greater than 20% for more than 15 minutes, then reduce the fleet capacity by
// 10%." A policy's rule statement has the following structure: If [MetricName] is
// [ComparisonOperator] [Threshold] for [EvaluationPeriods] minutes, then
// [ScalingAdjustmentType] to/by [ScalingAdjustment] . To implement the example,
// the rule statement would look like this: If [PercentIdleInstances] is
// [GreaterThanThreshold] [20] for [15] minutes, then [PercentChangeInCapacity]
// to/by [10] . To create or update a scaling policy, specify a unique combination
// of name and fleet ID, and set the policy type to "RuleBased". Specify the
// parameter values for a policy rule statement. On a successful request, the
// policy name is returned. Scaling policies are automatically in force as soon as
// they're successfully created. If the fleet's auto-scaling actions are
// temporarily suspended, the new policy will be in force once the fleet actions
// are restarted.
func (c *Client) PutScalingPolicy(ctx context.Context, params *PutScalingPolicyInput, optFns ...func(*Options)) (*PutScalingPolicyOutput, error) {
	if params == nil {
		params = &PutScalingPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutScalingPolicy", params, optFns, c.addOperationPutScalingPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutScalingPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutScalingPolicyInput struct {

	// A unique identifier for the fleet to apply this policy to. You can use either
	// the fleet ID or ARN value. The fleet cannot be in any of the following statuses:
	// ERROR or DELETING.
	//
	// This member is required.
	FleetId *string

	// Name of the Amazon GameLift-defined metric that is used to trigger a scaling
	// adjustment. For detailed descriptions of fleet metrics, see Monitor Amazon
	// GameLift with Amazon CloudWatch (https://docs.aws.amazon.com/gamelift/latest/developerguide/monitoring-cloudwatch.html)
	// .
	//   - ActivatingGameSessions -- Game sessions in the process of being created.
	//   - ActiveGameSessions -- Game sessions that are currently running.
	//   - ActiveInstances -- Fleet instances that are currently running at least one
	//   game session.
	//   - AvailableGameSessions -- Additional game sessions that fleet could host
	//   simultaneously, given current capacity.
	//   - AvailablePlayerSessions -- Empty player slots in currently active game
	//   sessions. This includes game sessions that are not currently accepting players.
	//   Reserved player slots are not included.
	//   - CurrentPlayerSessions -- Player slots in active game sessions that are
	//   being used by a player or are reserved for a player.
	//   - IdleInstances -- Active instances that are currently hosting zero game
	//   sessions.
	//   - PercentAvailableGameSessions -- Unused percentage of the total number of
	//   game sessions that a fleet could host simultaneously, given current capacity.
	//   Use this metric for a target-based scaling policy.
	//   - PercentIdleInstances -- Percentage of the total number of active instances
	//   that are hosting zero game sessions.
	//   - QueueDepth -- Pending game session placement requests, in any queue, where
	//   the current fleet is the top-priority destination.
	//   - WaitTime -- Current wait time for pending game session placement requests,
	//   in any queue, where the current fleet is the top-priority destination.
	//
	// This member is required.
	MetricName types.MetricName

	// A descriptive label that is associated with a fleet's scaling policy. Policy
	// names do not need to be unique. A fleet can have only one scaling policy with
	// the same name.
	//
	// This member is required.
	Name *string

	// Comparison operator to use when measuring the metric against the threshold
	// value.
	ComparisonOperator types.ComparisonOperatorType

	// Length of time (in minutes) the metric must be at or beyond the threshold
	// before a scaling event is triggered.
	EvaluationPeriods *int32

	// The type of scaling policy to create. For a target-based policy, set the
	// parameter MetricName to 'PercentAvailableGameSessions' and specify a
	// TargetConfiguration. For a rule-based policy set the following parameters:
	// MetricName, ComparisonOperator, Threshold, EvaluationPeriods,
	// ScalingAdjustmentType, and ScalingAdjustment.
	PolicyType types.PolicyType

	// Amount of adjustment to make, based on the scaling adjustment type.
	ScalingAdjustment int32

	// The type of adjustment to make to a fleet's instance count:
	//   - ChangeInCapacity -- add (or subtract) the scaling adjustment value from the
	//   current instance count. Positive values scale up while negative values scale
	//   down.
	//   - ExactCapacity -- set the instance count to the scaling adjustment value.
	//   - PercentChangeInCapacity -- increase or reduce the current instance count by
	//   the scaling adjustment, read as a percentage. Positive values scale up while
	//   negative values scale down; for example, a value of "-10" scales the fleet down
	//   by 10%.
	ScalingAdjustmentType types.ScalingAdjustmentType

	// An object that contains settings for a target-based scaling policy.
	TargetConfiguration *types.TargetConfiguration

	// Metric value used to trigger a scaling event.
	Threshold float64

	noSmithyDocumentSerde
}

type PutScalingPolicyOutput struct {

	// A descriptive label that is associated with a fleet's scaling policy. Policy
	// names do not need to be unique.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutScalingPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutScalingPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutScalingPolicy{}, middleware.After)
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
	if err = addOpPutScalingPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutScalingPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutScalingPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "PutScalingPolicy",
	}
}
