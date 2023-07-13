// Code generated by smithy-go-codegen DO NOT EDIT.

package iotfleetwise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about the status of registering your Amazon Web Services
// account, IAM, and Amazon Timestream resources so that Amazon Web Services IoT
// FleetWise can transfer your vehicle data to the Amazon Web Services Cloud. For
// more information, including step-by-step procedures, see Setting up Amazon Web
// Services IoT FleetWise (https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/setting-up.html)
// . This API operation doesn't require input parameters.
func (c *Client) GetRegisterAccountStatus(ctx context.Context, params *GetRegisterAccountStatusInput, optFns ...func(*Options)) (*GetRegisterAccountStatusOutput, error) {
	if params == nil {
		params = &GetRegisterAccountStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRegisterAccountStatus", params, optFns, c.addOperationGetRegisterAccountStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRegisterAccountStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRegisterAccountStatusInput struct {
	noSmithyDocumentSerde
}

type GetRegisterAccountStatusOutput struct {

	// The status of registering your account and resources. The status can be one of:
	//   - REGISTRATION_SUCCESS - The Amazon Web Services resource is successfully
	//   registered.
	//   - REGISTRATION_PENDING - Amazon Web Services IoT FleetWise is processing the
	//   registration request. This process takes approximately five minutes to complete.
	//
	//   - REGISTRATION_FAILURE - Amazon Web Services IoT FleetWise can't register the
	//   AWS resource. Try again later.
	//
	// This member is required.
	AccountStatus types.RegistrationStatus

	// The time the account was registered, in seconds since epoch (January 1, 1970 at
	// midnight UTC time).
	//
	// This member is required.
	CreationTime *time.Time

	// The unique ID of the Amazon Web Services account, provided at account creation.
	//
	// This member is required.
	CustomerAccountId *string

	// Information about the registered IAM resources or errors, if any.
	//
	// This member is required.
	IamRegistrationResponse *types.IamRegistrationResponse

	// The time this registration was last updated, in seconds since epoch (January 1,
	// 1970 at midnight UTC time).
	//
	// This member is required.
	LastModificationTime *time.Time

	// Information about the registered Amazon Timestream resources or errors, if any.
	TimestreamRegistrationResponse *types.TimestreamRegistrationResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRegisterAccountStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetRegisterAccountStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetRegisterAccountStatus{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRegisterAccountStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRegisterAccountStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotfleetwise",
		OperationName: "GetRegisterAccountStatus",
	}
}
