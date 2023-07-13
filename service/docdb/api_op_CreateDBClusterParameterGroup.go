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

// Creates a new cluster parameter group. Parameters in a cluster parameter group
// apply to all of the instances in a cluster. A cluster parameter group is
// initially created with the default parameters for the database engine used by
// instances in the cluster. In Amazon DocumentDB, you cannot make modifications
// directly to the default.docdb3.6 cluster parameter group. If your Amazon
// DocumentDB cluster is using the default cluster parameter group and you want to
// modify a value in it, you must first create a new parameter group (https://docs.aws.amazon.com/documentdb/latest/developerguide/cluster_parameter_group-create.html)
// or copy an existing parameter group (https://docs.aws.amazon.com/documentdb/latest/developerguide/cluster_parameter_group-copy.html)
// , modify it, and then apply the modified parameter group to your cluster. For
// the new cluster parameter group and associated settings to take effect, you must
// then reboot the instances in the cluster without failover. For more information,
// see Modifying Amazon DocumentDB Cluster Parameter Groups (https://docs.aws.amazon.com/documentdb/latest/developerguide/cluster_parameter_group-modify.html)
// .
func (c *Client) CreateDBClusterParameterGroup(ctx context.Context, params *CreateDBClusterParameterGroupInput, optFns ...func(*Options)) (*CreateDBClusterParameterGroupOutput, error) {
	if params == nil {
		params = &CreateDBClusterParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBClusterParameterGroup", params, optFns, c.addOperationCreateDBClusterParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBClusterParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of CreateDBClusterParameterGroup .
type CreateDBClusterParameterGroupInput struct {

	// The name of the cluster parameter group. Constraints:
	//   - Must not match the name of an existing DBClusterParameterGroup .
	// This value is stored as a lowercase string.
	//
	// This member is required.
	DBClusterParameterGroupName *string

	// The cluster parameter group family name.
	//
	// This member is required.
	DBParameterGroupFamily *string

	// The description for the cluster parameter group.
	//
	// This member is required.
	Description *string

	// The tags to be assigned to the cluster parameter group.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDBClusterParameterGroupOutput struct {

	// Detailed information about a cluster parameter group.
	DBClusterParameterGroup *types.DBClusterParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDBClusterParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBClusterParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBClusterParameterGroup{}, middleware.After)
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
	if err = addOpCreateDBClusterParameterGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBClusterParameterGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDBClusterParameterGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBClusterParameterGroup",
	}
}
