// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the name, description, and rules in a device pool given the attributes
// and the pool ARN. Rule updates are all-or-nothing, meaning they can only be
// updated as a whole (or not at all).
func (c *Client) UpdateDevicePool(ctx context.Context, params *UpdateDevicePoolInput, optFns ...func(*Options)) (*UpdateDevicePoolOutput, error) {
	if params == nil {
		params = &UpdateDevicePoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDevicePool", params, optFns, c.addOperationUpdateDevicePoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDevicePoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to the update device pool operation.
type UpdateDevicePoolInput struct {

	// The Amazon Resource Name (ARN) of the Device Farm device pool to update.
	//
	// This member is required.
	Arn *string

	// Sets whether the maxDevices parameter applies to your device pool. If you set
	// this parameter to true , the maxDevices parameter does not apply, and Device
	// Farm does not limit the number of devices that it adds to your device pool. In
	// this case, Device Farm adds all available devices that meet the criteria
	// specified in the rules parameter. If you use this parameter in your request,
	// you cannot use the maxDevices parameter in the same request.
	ClearMaxDevices *bool

	// A description of the device pool to update.
	Description *string

	// The number of devices that Device Farm can add to your device pool. Device Farm
	// adds devices that are available and that meet the criteria that you assign for
	// the rules parameter. Depending on how many devices meet these constraints, your
	// device pool might contain fewer devices than the value for this parameter. By
	// specifying the maximum number of devices, you can control the costs that you
	// incur by running tests. If you use this parameter in your request, you cannot
	// use the clearMaxDevices parameter in the same request.
	MaxDevices *int32

	// A string that represents the name of the device pool to update.
	Name *string

	// Represents the rules to modify for the device pool. Updating rules is optional.
	// If you update rules for your request, the update replaces the existing rules.
	Rules []types.Rule

	noSmithyDocumentSerde
}

// Represents the result of an update device pool request.
type UpdateDevicePoolOutput struct {

	// The device pool you just updated.
	DevicePool *types.DevicePool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDevicePoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDevicePool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDevicePool{}, middleware.After)
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
	if err = addOpUpdateDevicePoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDevicePool(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDevicePool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "UpdateDevicePool",
	}
}
