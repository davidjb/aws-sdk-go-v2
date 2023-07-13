// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Aurora global database spread across multiple Amazon Web Services
// Regions. The global database contains a single primary cluster with read-write
// capability, and a read-only secondary cluster that receives data from the
// primary cluster through high-speed replication performed by the Aurora storage
// subsystem. You can create a global database that is initially empty, and then
// add a primary cluster and a secondary cluster to it. Or you can specify an
// existing Aurora cluster during the create operation, and this cluster becomes
// the primary cluster of the global database. This operation applies only to
// Aurora DB clusters.
func (c *Client) CreateGlobalCluster(ctx context.Context, params *CreateGlobalClusterInput, optFns ...func(*Options)) (*CreateGlobalClusterOutput, error) {
	if params == nil {
		params = &CreateGlobalClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGlobalCluster", params, optFns, c.addOperationCreateGlobalClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGlobalClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGlobalClusterInput struct {

	// The name for your database of up to 64 alphanumeric characters. If you don't
	// specify a name, Amazon Aurora doesn't create a database in the global database
	// cluster. Constraints:
	//   - Can't be specified if SourceDBClusterIdentifier is specified. In this case,
	//   Amazon Aurora uses the database name from the source DB cluster.
	DatabaseName *string

	// Specifies whether to enable deletion protection for the new global database
	// cluster. The global database can't be deleted when deletion protection is
	// enabled.
	DeletionProtection *bool

	// The database engine to use for this global database cluster. Valid Values:
	// aurora-mysql | aurora-postgresql Constraints:
	//   - Can't be specified if SourceDBClusterIdentifier is specified. In this case,
	//   Amazon Aurora uses the engine of the source DB cluster.
	Engine *string

	// The engine version to use for this global database cluster. Constraints:
	//   - Can't be specified if SourceDBClusterIdentifier is specified. In this case,
	//   Amazon Aurora uses the engine version of the source DB cluster.
	EngineVersion *string

	// The cluster identifier for this global database cluster. This parameter is
	// stored as a lowercase string.
	GlobalClusterIdentifier *string

	// The Amazon Resource Name (ARN) to use as the primary cluster of the global
	// database. If you provide a value for this parameter, don't specify values for
	// the following settings because Amazon Aurora uses the values from the specified
	// source DB cluster:
	//   - DatabaseName
	//   - Engine
	//   - EngineVersion
	//   - StorageEncrypted
	SourceDBClusterIdentifier *string

	// Specifies whether to enable storage encryption for the new global database
	// cluster. Constraints:
	//   - Can't be specified if SourceDBClusterIdentifier is specified. In this case,
	//   Amazon Aurora uses the setting from the source DB cluster.
	StorageEncrypted *bool

	noSmithyDocumentSerde
}

type CreateGlobalClusterOutput struct {

	// A data type representing an Aurora global database.
	GlobalCluster *types.GlobalCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateGlobalClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateGlobalCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateGlobalCluster{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGlobalCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateGlobalCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateGlobalCluster",
	}
}
