// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an existing replication group. By default, this operation deletes the
// entire replication group, including the primary/primaries and all of the read
// replicas. If the replication group has only one primary, you can optionally
// delete only the read replicas, while retaining the primary by setting
// RetainPrimaryCluster=true . When you receive a successful response from this
// operation, Amazon ElastiCache immediately begins deleting the selected
// resources; you cannot cancel or revert this operation. This operation is valid
// for Redis only.
func (c *Client) DeleteReplicationGroup(ctx context.Context, params *DeleteReplicationGroupInput, optFns ...func(*Options)) (*DeleteReplicationGroupOutput, error) {
	if params == nil {
		params = &DeleteReplicationGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteReplicationGroup", params, optFns, c.addOperationDeleteReplicationGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteReplicationGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DeleteReplicationGroup operation.
type DeleteReplicationGroupInput struct {

	// The identifier for the cluster to be deleted. This parameter is not case
	// sensitive.
	//
	// This member is required.
	ReplicationGroupId *string

	// The name of a final node group (shard) snapshot. ElastiCache creates the
	// snapshot from the primary node in the cluster, rather than one of the replicas;
	// this is to ensure that it captures the freshest data. After the final snapshot
	// is taken, the replication group is immediately deleted.
	FinalSnapshotIdentifier *string

	// If set to true , all of the read replicas are deleted, but the primary node is
	// retained.
	RetainPrimaryCluster *bool

	noSmithyDocumentSerde
}

type DeleteReplicationGroupOutput struct {

	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *types.ReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteReplicationGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteReplicationGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteReplicationGroup{}, middleware.After)
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
	if err = addOpDeleteReplicationGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteReplicationGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteReplicationGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DeleteReplicationGroup",
	}
}
