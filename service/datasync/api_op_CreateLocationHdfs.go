// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an endpoint for a Hadoop Distributed File System (HDFS).
func (c *Client) CreateLocationHdfs(ctx context.Context, params *CreateLocationHdfsInput, optFns ...func(*Options)) (*CreateLocationHdfsOutput, error) {
	if params == nil {
		params = &CreateLocationHdfsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLocationHdfs", params, optFns, c.addOperationCreateLocationHdfsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLocationHdfsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLocationHdfsInput struct {

	// The Amazon Resource Names (ARNs) of the agents that are used to connect to the
	// HDFS cluster.
	//
	// This member is required.
	AgentArns []string

	// The type of authentication used to determine the identity of the user.
	//
	// This member is required.
	AuthenticationType types.HdfsAuthenticationType

	// The NameNode that manages the HDFS namespace. The NameNode performs operations
	// such as opening, closing, and renaming files and directories. The NameNode
	// contains the information to map blocks of data to the DataNodes. You can use
	// only one NameNode.
	//
	// This member is required.
	NameNodes []types.HdfsNameNode

	// The size of data blocks to write into the HDFS cluster. The block size must be
	// a multiple of 512 bytes. The default block size is 128 mebibytes (MiB).
	BlockSize *int32

	// The Kerberos key table (keytab) that contains mappings between the defined
	// Kerberos principal and the encrypted keys. You can load the keytab from a file
	// by providing the file's address. If you're using the CLI, it performs base64
	// encoding for you. Otherwise, provide the base64-encoded text. If KERBEROS is
	// specified for AuthenticationType , this parameter is required.
	KerberosKeytab []byte

	// The krb5.conf file that contains the Kerberos configuration information. You
	// can load the krb5.conf file by providing the file's address. If you're using
	// the CLI, it performs the base64 encoding for you. Otherwise, provide the
	// base64-encoded text. If KERBEROS is specified for AuthenticationType , this
	// parameter is required.
	KerberosKrb5Conf []byte

	// The Kerberos principal with access to the files and folders on the HDFS
	// cluster. If KERBEROS is specified for AuthenticationType , this parameter is
	// required.
	KerberosPrincipal *string

	// The URI of the HDFS cluster's Key Management Server (KMS).
	KmsKeyProviderUri *string

	// The Quality of Protection (QOP) configuration specifies the Remote Procedure
	// Call (RPC) and data transfer protection settings configured on the Hadoop
	// Distributed File System (HDFS) cluster. If QopConfiguration isn't specified,
	// RpcProtection and DataTransferProtection default to PRIVACY . If you set
	// RpcProtection or DataTransferProtection , the other parameter assumes the same
	// value.
	QopConfiguration *types.QopConfiguration

	// The number of DataNodes to replicate the data to when writing to the HDFS
	// cluster. By default, data is replicated to three DataNodes.
	ReplicationFactor *int32

	// The user name used to identify the client on the host operating system. If
	// SIMPLE is specified for AuthenticationType , this parameter is required.
	SimpleUser *string

	// A subdirectory in the HDFS cluster. This subdirectory is used to read data from
	// or write data to the HDFS cluster. If the subdirectory isn't specified, it will
	// default to / .
	Subdirectory *string

	// The key-value pair that represents the tag that you want to add to the
	// location. The value can be an empty string. We recommend using tags to name your
	// resources.
	Tags []types.TagListEntry

	noSmithyDocumentSerde
}

type CreateLocationHdfsOutput struct {

	// The ARN of the source HDFS cluster location that's created.
	LocationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLocationHdfsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLocationHdfs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLocationHdfs{}, middleware.After)
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
	if err = addOpCreateLocationHdfsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocationHdfs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLocationHdfs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "CreateLocationHdfs",
	}
}
