// Code generated by smithy-go-codegen DO NOT EDIT.

package memorydb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/memorydb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the settings for a cluster. You can use this operation to change one
// or more cluster configuration settings by specifying the settings and the new
// values.
func (c *Client) UpdateCluster(ctx context.Context, params *UpdateClusterInput, optFns ...func(*Options)) (*UpdateClusterOutput, error) {
	if params == nil {
		params = &UpdateClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCluster", params, optFns, c.addOperationUpdateClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateClusterInput struct {

	// The name of the cluster to update
	//
	// This member is required.
	ClusterName *string

	// The Access Control List that is associated with the cluster
	ACLName *string

	// The description of the cluster to update
	Description *string

	// The upgraded version of the engine to be run on the nodes. You can upgrade to a
	// newer engine version, but you cannot downgrade to an earlier engine version. If
	// you want to use an earlier engine version, you must delete the existing cluster
	// and create it anew with the earlier engine version.
	EngineVersion *string

	// Specifies the weekly time range during which maintenance on the cluster is
	// performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H
	// Clock UTC). The minimum maintenance window is a 60 minute period. Valid values
	// for ddd are:
	//   - sun
	//   - mon
	//   - tue
	//   - wed
	//   - thu
	//   - fri
	//   - sat
	// Example: sun:23:00-mon:01:30
	MaintenanceWindow *string

	// A valid node type that you want to scale this cluster up or down to.
	NodeType *string

	// The name of the parameter group to update
	ParameterGroupName *string

	// The number of replicas that will reside in each shard
	ReplicaConfiguration *types.ReplicaConfigurationRequest

	// The SecurityGroupIds to update
	SecurityGroupIds []string

	// The number of shards in the cluster
	ShardConfiguration *types.ShardConfigurationRequest

	// The number of days for which MemoryDB retains automatic cluster snapshots
	// before deleting them. For example, if you set SnapshotRetentionLimit to 5, a
	// snapshot that was taken today is retained for 5 days before being deleted.
	SnapshotRetentionLimit *int32

	// The daily time range (in UTC) during which MemoryDB begins taking a daily
	// snapshot of your cluster.
	SnapshotWindow *string

	// The SNS topic ARN to update
	SnsTopicArn *string

	// The status of the Amazon SNS notification topic. Notifications are sent only if
	// the status is active.
	SnsTopicStatus *string

	noSmithyDocumentSerde
}

type UpdateClusterOutput struct {

	// The updated cluster
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCluster{}, middleware.After)
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
	if err = addOpUpdateClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "memorydb",
		OperationName: "UpdateCluster",
	}
}
