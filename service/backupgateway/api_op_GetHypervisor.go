// Code generated by smithy-go-codegen DO NOT EDIT.

package backupgateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backupgateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action requests information about the specified hypervisor to which the
// gateway will connect. A hypervisor is hardware, software, or firmware that
// creates and manages virtual machines, and allocates resources to them.
func (c *Client) GetHypervisor(ctx context.Context, params *GetHypervisorInput, optFns ...func(*Options)) (*GetHypervisorOutput, error) {
	if params == nil {
		params = &GetHypervisorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetHypervisor", params, optFns, c.addOperationGetHypervisorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetHypervisorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetHypervisorInput struct {

	// The Amazon Resource Name (ARN) of the hypervisor.
	//
	// This member is required.
	HypervisorArn *string

	noSmithyDocumentSerde
}

type GetHypervisorOutput struct {

	// Details about the requested hypervisor.
	Hypervisor *types.HypervisorDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetHypervisorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetHypervisor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetHypervisor{}, middleware.After)
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
	if err = addOpGetHypervisorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetHypervisor(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetHypervisor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup-gateway",
		OperationName: "GetHypervisor",
	}
}
