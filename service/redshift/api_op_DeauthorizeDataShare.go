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

// From a datashare producer account, removes authorization from the specified
// datashare.
func (c *Client) DeauthorizeDataShare(ctx context.Context, params *DeauthorizeDataShareInput, optFns ...func(*Options)) (*DeauthorizeDataShareOutput, error) {
	if params == nil {
		params = &DeauthorizeDataShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeauthorizeDataShare", params, optFns, c.addOperationDeauthorizeDataShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeauthorizeDataShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeauthorizeDataShareInput struct {

	// The identifier of the data consumer that is to have authorization removed from
	// the datashare. This identifier is an Amazon Web Services account ID or a
	// keyword, such as ADX.
	//
	// This member is required.
	ConsumerIdentifier *string

	// The Amazon Resource Name (ARN) of the datashare to remove authorization from.
	//
	// This member is required.
	DataShareArn *string

	noSmithyDocumentSerde
}

type DeauthorizeDataShareOutput struct {

	// A value that specifies whether the datashare can be shared to a publicly
	// accessible cluster.
	AllowPubliclyAccessibleConsumers bool

	// An Amazon Resource Name (ARN) that references the datashare that is owned by a
	// specific namespace of the producer cluster. A datashare ARN is in the
	// arn:aws:redshift:{region}:{account-id}:{datashare}:{namespace-guid}/{datashare-name}
	// format.
	DataShareArn *string

	// A value that specifies when the datashare has an association between producer
	// and data consumers.
	DataShareAssociations []types.DataShareAssociation

	// The identifier of a datashare to show its managing entity.
	ManagedBy *string

	// The Amazon Resource Name (ARN) of the producer.
	ProducerArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeauthorizeDataShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeauthorizeDataShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeauthorizeDataShare{}, middleware.After)
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
	if err = addOpDeauthorizeDataShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeauthorizeDataShare(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeauthorizeDataShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DeauthorizeDataShare",
	}
}
