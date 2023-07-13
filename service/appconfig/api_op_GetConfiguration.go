// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// (Deprecated) Retrieves the latest deployed configuration. Note the following
// important information.
//   - This API action is deprecated. Calls to receive configuration data should
//     use the StartConfigurationSession (https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_appconfigdata_StartConfigurationSession.html)
//     and GetLatestConfiguration (https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_appconfigdata_GetLatestConfiguration.html)
//     APIs instead.
//   - GetConfiguration is a priced call. For more information, see Pricing (https://aws.amazon.com/systems-manager/pricing/)
//     .
//
// Deprecated: This API has been deprecated in favor of the GetLatestConfiguration
// API used in conjunction with StartConfigurationSession.
func (c *Client) GetConfiguration(ctx context.Context, params *GetConfigurationInput, optFns ...func(*Options)) (*GetConfigurationOutput, error) {
	if params == nil {
		params = &GetConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConfiguration", params, optFns, c.addOperationGetConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConfigurationInput struct {

	// The application to get. Specify either the application name or the application
	// ID.
	//
	// This member is required.
	Application *string

	// The clientId parameter in the following command is a unique, user-specified ID
	// to identify the client for the configuration. This ID enables AppConfig to
	// deploy the configuration in intervals, as defined in the deployment strategy.
	//
	// This member is required.
	ClientId *string

	// The configuration to get. Specify either the configuration name or the
	// configuration ID.
	//
	// This member is required.
	Configuration *string

	// The environment to get. Specify either the environment name or the environment
	// ID.
	//
	// This member is required.
	Environment *string

	// The configuration version returned in the most recent GetConfiguration
	// response. AppConfig uses the value of the ClientConfigurationVersion parameter
	// to identify the configuration version on your clients. If you don’t send
	// ClientConfigurationVersion with each call to GetConfiguration , your clients
	// receive the current configuration. You are charged each time your clients
	// receive a configuration. To avoid excess charges, we recommend you use the
	// StartConfigurationSession (https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/StartConfigurationSession.html)
	// and GetLatestConfiguration (https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/GetLatestConfiguration.html)
	// APIs, which track the client configuration version on your behalf. If you choose
	// to continue using GetConfiguration , we recommend that you include the
	// ClientConfigurationVersion value with every call to GetConfiguration . The value
	// to use for ClientConfigurationVersion comes from the ConfigurationVersion
	// attribute returned by GetConfiguration when there is new or updated data, and
	// should be saved for subsequent calls to GetConfiguration . For more information
	// about working with configurations, see Retrieving the Configuration (http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-retrieving-the-configuration.html)
	// in the AppConfig User Guide.
	ClientConfigurationVersion *string

	noSmithyDocumentSerde
}

type GetConfigurationOutput struct {

	// The configuration version.
	ConfigurationVersion *string

	// The content of the configuration or the configuration data. The Content
	// attribute only contains data if the system finds new or updated configuration
	// data. If there is no new or updated data and ClientConfigurationVersion matches
	// the version of the current configuration, AppConfig returns a 204 No Content
	// HTTP response code and the Content value will be empty.
	Content []byte

	// A standard MIME type describing the format of the configuration content. For
	// more information, see Content-Type (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17)
	// .
	ContentType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetConfiguration{}, middleware.After)
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
	if err = addOpGetConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "GetConfiguration",
	}
}
