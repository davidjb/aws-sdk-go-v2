// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the Amazon Inspector deep inspection custom paths for your
// organization. You must be an Amazon Inspector delegated administrator to use
// this API.
func (c *Client) UpdateOrgEc2DeepInspectionConfiguration(ctx context.Context, params *UpdateOrgEc2DeepInspectionConfigurationInput, optFns ...func(*Options)) (*UpdateOrgEc2DeepInspectionConfigurationOutput, error) {
	if params == nil {
		params = &UpdateOrgEc2DeepInspectionConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateOrgEc2DeepInspectionConfiguration", params, optFns, c.addOperationUpdateOrgEc2DeepInspectionConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateOrgEc2DeepInspectionConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateOrgEc2DeepInspectionConfigurationInput struct {

	// The Amazon Inspector deep inspection custom paths you are adding for your
	// organization.
	//
	// This member is required.
	OrgPackagePaths []string

	noSmithyDocumentSerde
}

type UpdateOrgEc2DeepInspectionConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateOrgEc2DeepInspectionConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateOrgEc2DeepInspectionConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateOrgEc2DeepInspectionConfiguration{}, middleware.After)
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
	if err = addOpUpdateOrgEc2DeepInspectionConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateOrgEc2DeepInspectionConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateOrgEc2DeepInspectionConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector2",
		OperationName: "UpdateOrgEc2DeepInspectionConfiguration",
	}
}
