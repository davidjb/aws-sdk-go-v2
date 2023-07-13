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

// Creates a new DB proxy.
func (c *Client) CreateDBProxy(ctx context.Context, params *CreateDBProxyInput, optFns ...func(*Options)) (*CreateDBProxyOutput, error) {
	if params == nil {
		params = &CreateDBProxyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBProxy", params, optFns, c.addOperationCreateDBProxyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBProxyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDBProxyInput struct {

	// The authorization mechanism that the proxy uses.
	//
	// This member is required.
	Auth []types.UserAuthConfig

	// The identifier for the proxy. This name must be unique for all proxies owned by
	// your Amazon Web Services account in the specified Amazon Web Services Region. An
	// identifier must begin with a letter and must contain only ASCII letters, digits,
	// and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	//
	// This member is required.
	DBProxyName *string

	// The kinds of databases that the proxy can connect to. This value determines
	// which database network protocol the proxy recognizes when it interprets network
	// traffic to and from the database. For Aurora MySQL, RDS for MariaDB, and RDS for
	// MySQL databases, specify MYSQL . For Aurora PostgreSQL and RDS for PostgreSQL
	// databases, specify POSTGRESQL . For RDS for Microsoft SQL Server, specify
	// SQLSERVER .
	//
	// This member is required.
	EngineFamily types.EngineFamily

	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access
	// secrets in Amazon Web Services Secrets Manager.
	//
	// This member is required.
	RoleArn *string

	// One or more VPC subnet IDs to associate with the new proxy.
	//
	// This member is required.
	VpcSubnetIds []string

	// Whether the proxy includes detailed information about SQL statements in its
	// logs. This information helps you to debug issues involving SQL behavior or the
	// performance and scalability of the proxy connections. The debug information
	// includes the text of SQL statements that you submit through the proxy. Thus,
	// only enable this setting when needed for debugging, and only when you have
	// security measures in place to safeguard any sensitive information that appears
	// in the logs.
	DebugLogging bool

	// The number of seconds that a connection to the proxy can be inactive before the
	// proxy disconnects it. You can set this value higher or lower than the connection
	// timeout limit for the associated database.
	IdleClientTimeout *int32

	// A Boolean parameter that specifies whether Transport Layer Security (TLS)
	// encryption is required for connections to the proxy. By enabling this setting,
	// you can enforce encrypted TLS connections to the proxy.
	RequireTLS bool

	// An optional set of key-value pairs to associate arbitrary data of your choosing
	// with the proxy.
	Tags []types.Tag

	// One or more VPC security group IDs to associate with the new proxy.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type CreateDBProxyOutput struct {

	// The DBProxy structure corresponding to the new proxy.
	DBProxy *types.DBProxy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDBProxyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBProxy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBProxy{}, middleware.After)
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
	if err = addOpCreateDBProxyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBProxy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDBProxy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBProxy",
	}
}
