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

// Exchanges a DC1 Reserved Node for a DC2 Reserved Node with no changes to the
// configuration (term, payment type, or number of nodes) and no additional costs.
func (c *Client) AcceptReservedNodeExchange(ctx context.Context, params *AcceptReservedNodeExchangeInput, optFns ...func(*Options)) (*AcceptReservedNodeExchangeOutput, error) {
	if params == nil {
		params = &AcceptReservedNodeExchangeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AcceptReservedNodeExchange", params, optFns, c.addOperationAcceptReservedNodeExchangeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AcceptReservedNodeExchangeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AcceptReservedNodeExchangeInput struct {

	// A string representing the node identifier of the DC1 Reserved Node to be
	// exchanged.
	//
	// This member is required.
	ReservedNodeId *string

	// The unique identifier of the DC2 Reserved Node offering to be used for the
	// exchange. You can obtain the value for the parameter by calling
	// GetReservedNodeExchangeOfferings
	//
	// This member is required.
	TargetReservedNodeOfferingId *string

	noSmithyDocumentSerde
}

type AcceptReservedNodeExchangeOutput struct {

	//
	ExchangedReservedNode *types.ReservedNode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAcceptReservedNodeExchangeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAcceptReservedNodeExchange{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAcceptReservedNodeExchange{}, middleware.After)
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
	if err = addOpAcceptReservedNodeExchangeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptReservedNodeExchange(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAcceptReservedNodeExchange(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "AcceptReservedNodeExchange",
	}
}
