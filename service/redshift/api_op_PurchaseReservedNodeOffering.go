// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows you to purchase reserved nodes. Amazon Redshift offers a predefined set
// of reserved node offerings. You can purchase one or more of the offerings. You
// can call the DescribeReservedNodeOfferings API to obtain the available reserved
// node offerings. You can call this API by providing a specific reserved node
// offering and the number of nodes you want to reserve. For more information about
// reserved node offerings, go to Purchasing Reserved Nodes (https://docs.aws.amazon.com/redshift/latest/mgmt/purchase-reserved-node-instance.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) PurchaseReservedNodeOffering(ctx context.Context, params *PurchaseReservedNodeOfferingInput, optFns ...func(*Options)) (*PurchaseReservedNodeOfferingOutput, error) {
	if params == nil {
		params = &PurchaseReservedNodeOfferingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PurchaseReservedNodeOffering", params, optFns, c.addOperationPurchaseReservedNodeOfferingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PurchaseReservedNodeOfferingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PurchaseReservedNodeOfferingInput struct {

	// The unique identifier of the reserved node offering you want to purchase.
	//
	// This member is required.
	ReservedNodeOfferingId *string

	// The number of reserved nodes that you want to purchase. Default: 1
	NodeCount *int32

	noSmithyDocumentSerde
}

type PurchaseReservedNodeOfferingOutput struct {

	// Describes a reserved node. You can call the DescribeReservedNodeOfferings API
	// to obtain the available reserved node offerings.
	ReservedNode *types.ReservedNode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPurchaseReservedNodeOfferingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPurchaseReservedNodeOffering{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPurchaseReservedNodeOffering{}, middleware.After)
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
	if err = addOpPurchaseReservedNodeOfferingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPurchaseReservedNodeOffering(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPurchaseReservedNodeOffering(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "PurchaseReservedNodeOffering",
	}
}
