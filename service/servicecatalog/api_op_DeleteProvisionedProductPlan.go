// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified plan.
func (c *Client) DeleteProvisionedProductPlan(ctx context.Context, params *DeleteProvisionedProductPlanInput, optFns ...func(*Options)) (*DeleteProvisionedProductPlanOutput, error) {
	if params == nil {
		params = &DeleteProvisionedProductPlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteProvisionedProductPlan", params, optFns, c.addOperationDeleteProvisionedProductPlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteProvisionedProductPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteProvisionedProductPlanInput struct {

	// The plan identifier.
	//
	// This member is required.
	PlanId *string

	// The language code.
	//   - jp - Japanese
	//   - zh - Chinese
	AcceptLanguage *string

	// If set to true, Service Catalog stops managing the specified provisioned
	// product even if it cannot delete the underlying resources.
	IgnoreErrors bool

	noSmithyDocumentSerde
}

type DeleteProvisionedProductPlanOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteProvisionedProductPlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteProvisionedProductPlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteProvisionedProductPlan{}, middleware.After)
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
	if err = addOpDeleteProvisionedProductPlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteProvisionedProductPlan(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteProvisionedProductPlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "DeleteProvisionedProductPlan",
	}
}
