// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates a mobile device access override for the given WorkMail
// organization, user, and device.
func (c *Client) PutMobileDeviceAccessOverride(ctx context.Context, params *PutMobileDeviceAccessOverrideInput, optFns ...func(*Options)) (*PutMobileDeviceAccessOverrideOutput, error) {
	if params == nil {
		params = &PutMobileDeviceAccessOverrideInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutMobileDeviceAccessOverride", params, optFns, c.addOperationPutMobileDeviceAccessOverrideMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutMobileDeviceAccessOverrideOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutMobileDeviceAccessOverrideInput struct {

	// The mobile device for which you create the override. DeviceId is case
	// insensitive.
	//
	// This member is required.
	DeviceId *string

	// The effect of the override, ALLOW or DENY .
	//
	// This member is required.
	Effect types.MobileDeviceAccessRuleEffect

	// Identifies the WorkMail organization for which you create the override.
	//
	// This member is required.
	OrganizationId *string

	// The WorkMail user for which you create the override. Accepts the following
	// types of user identities:
	//   - User ID: 12345678-1234-1234-1234-123456789012 or
	//   S-1-1-12-1234567890-123456789-123456789-1234
	//   - Email address: user@domain.tld
	//   - User name: user
	//
	// This member is required.
	UserId *string

	// A description of the override.
	Description *string

	noSmithyDocumentSerde
}

type PutMobileDeviceAccessOverrideOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutMobileDeviceAccessOverrideMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutMobileDeviceAccessOverride{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutMobileDeviceAccessOverride{}, middleware.After)
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
	if err = addOpPutMobileDeviceAccessOverrideValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutMobileDeviceAccessOverride(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutMobileDeviceAccessOverride(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "PutMobileDeviceAccessOverride",
	}
}
