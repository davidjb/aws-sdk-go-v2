// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates the connector, which captures the parameters for an outbound connection
// for the AS2 protocol. The connector is required for sending files to an
// externally hosted AS2 server. For more details about connectors, see Create AS2
// connectors (https://docs.aws.amazon.com/transfer/latest/userguide/create-b2b-server.html#configure-as2-connector)
// .
func (c *Client) CreateConnector(ctx context.Context, params *CreateConnectorInput, optFns ...func(*Options)) (*CreateConnectorOutput, error) {
	if params == nil {
		params = &CreateConnectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConnector", params, optFns, c.addOperationCreateConnectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConnectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConnectorInput struct {

	// With AS2, you can send files by calling StartFileTransfer and specifying the
	// file paths in the request parameter, SendFilePaths . We use the file’s parent
	// directory (for example, for --send-file-paths /bucket/dir/file.txt , parent
	// directory is /bucket/dir/ ) to temporarily store a processed AS2 message file,
	// store the MDN when we receive them from the partner, and write a final JSON file
	// containing relevant metadata of the transmission. So, the AccessRole needs to
	// provide read and write access to the parent directory of the file location used
	// in the StartFileTransfer request. Additionally, you need to provide read and
	// write access to the parent directory of the files that you intend to send with
	// StartFileTransfer . If you are using Basic authentication for your AS2
	// connector, the access role requires the secretsmanager:GetSecretValue
	// permission for the secret. If the secret is encrypted using a customer-managed
	// key instead of the Amazon Web Services managed key in Secrets Manager, then the
	// role also needs the kms:Decrypt permission for that key.
	//
	// This member is required.
	AccessRole *string

	// A structure that contains the parameters for a connector object.
	//
	// This member is required.
	As2Config *types.As2ConnectorConfig

	// The URL of the partner's AS2 endpoint.
	//
	// This member is required.
	Url *string

	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role
	// that allows a connector to turn on CloudWatch logging for Amazon S3 events. When
	// set, you can view connector activity in your CloudWatch logs.
	LoggingRole *string

	// Key-value pairs that can be used to group and search for connectors. Tags are
	// metadata attached to connectors for any purpose.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateConnectorOutput struct {

	// The unique identifier for the connector, returned after the API call succeeds.
	//
	// This member is required.
	ConnectorId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConnectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateConnector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateConnector{}, middleware.After)
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
	if err = addOpCreateConnectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConnector(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateConnector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "CreateConnector",
	}
}
