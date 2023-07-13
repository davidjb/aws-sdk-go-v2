// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the emergency calling configuration details for the specified Voice
// Connector.
func (c *Client) GetVoiceConnectorEmergencyCallingConfiguration(ctx context.Context, params *GetVoiceConnectorEmergencyCallingConfigurationInput, optFns ...func(*Options)) (*GetVoiceConnectorEmergencyCallingConfigurationOutput, error) {
	if params == nil {
		params = &GetVoiceConnectorEmergencyCallingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVoiceConnectorEmergencyCallingConfiguration", params, optFns, c.addOperationGetVoiceConnectorEmergencyCallingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVoiceConnectorEmergencyCallingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVoiceConnectorEmergencyCallingConfigurationInput struct {

	// The Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string

	noSmithyDocumentSerde
}

type GetVoiceConnectorEmergencyCallingConfigurationOutput struct {

	// The details of the emergency calling configuration.
	EmergencyCallingConfiguration *types.EmergencyCallingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetVoiceConnectorEmergencyCallingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetVoiceConnectorEmergencyCallingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetVoiceConnectorEmergencyCallingConfiguration{}, middleware.After)
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
	if err = addOpGetVoiceConnectorEmergencyCallingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetVoiceConnectorEmergencyCallingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetVoiceConnectorEmergencyCallingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetVoiceConnectorEmergencyCallingConfiguration",
	}
}
