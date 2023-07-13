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

// Simulates the effect of the mobile device access rules for the given attributes
// of a sample access event. Use this method to test the effects of the current set
// of mobile device access rules for the WorkMail organization for a particular
// user's attributes.
func (c *Client) GetMobileDeviceAccessEffect(ctx context.Context, params *GetMobileDeviceAccessEffectInput, optFns ...func(*Options)) (*GetMobileDeviceAccessEffectOutput, error) {
	if params == nil {
		params = &GetMobileDeviceAccessEffectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMobileDeviceAccessEffect", params, optFns, c.addOperationGetMobileDeviceAccessEffectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMobileDeviceAccessEffectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMobileDeviceAccessEffectInput struct {

	// The WorkMail organization to simulate the access effect for.
	//
	// This member is required.
	OrganizationId *string

	// Device model the simulated user will report.
	DeviceModel *string

	// Device operating system the simulated user will report.
	DeviceOperatingSystem *string

	// Device type the simulated user will report.
	DeviceType *string

	// Device user agent the simulated user will report.
	DeviceUserAgent *string

	noSmithyDocumentSerde
}

type GetMobileDeviceAccessEffectOutput struct {

	// The effect of the simulated access, ALLOW or DENY , after evaluating mobile
	// device access rules in the WorkMail organization for the simulated user
	// parameters.
	Effect types.MobileDeviceAccessRuleEffect

	// A list of the rules which matched the simulated user input and produced the
	// effect.
	MatchedRules []types.MobileDeviceAccessMatchedRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMobileDeviceAccessEffectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMobileDeviceAccessEffect{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMobileDeviceAccessEffect{}, middleware.After)
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
	if err = addOpGetMobileDeviceAccessEffectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMobileDeviceAccessEffect(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMobileDeviceAccessEffect(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "GetMobileDeviceAccessEffect",
	}
}
