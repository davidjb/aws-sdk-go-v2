// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates an automatic scaling policy for a core instance group or
// task instance group in an Amazon EMR cluster. The automatic scaling policy
// defines how an instance group dynamically adds and terminates Amazon EC2
// instances in response to the value of a CloudWatch metric.
func (c *Client) PutAutoScalingPolicy(ctx context.Context, params *PutAutoScalingPolicyInput, optFns ...func(*Options)) (*PutAutoScalingPolicyOutput, error) {
	if params == nil {
		params = &PutAutoScalingPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAutoScalingPolicy", params, optFns, c.addOperationPutAutoScalingPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAutoScalingPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAutoScalingPolicyInput struct {

	// Specifies the definition of the automatic scaling policy.
	//
	// This member is required.
	AutoScalingPolicy *types.AutoScalingPolicy

	// Specifies the ID of a cluster. The instance group to which the automatic
	// scaling policy is applied is within this cluster.
	//
	// This member is required.
	ClusterId *string

	// Specifies the ID of the instance group to which the automatic scaling policy is
	// applied.
	//
	// This member is required.
	InstanceGroupId *string

	noSmithyDocumentSerde
}

type PutAutoScalingPolicyOutput struct {

	// The automatic scaling policy definition.
	AutoScalingPolicy *types.AutoScalingPolicyDescription

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string

	// Specifies the ID of a cluster. The instance group to which the automatic
	// scaling policy is applied is within this cluster.
	ClusterId *string

	// Specifies the ID of the instance group to which the scaling policy is applied.
	InstanceGroupId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAutoScalingPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutAutoScalingPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAutoScalingPolicy{}, middleware.After)
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
	if err = addOpPutAutoScalingPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAutoScalingPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutAutoScalingPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "PutAutoScalingPolicy",
	}
}
