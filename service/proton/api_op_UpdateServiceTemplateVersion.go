// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update a major or minor version of a service template.
func (c *Client) UpdateServiceTemplateVersion(ctx context.Context, params *UpdateServiceTemplateVersionInput, optFns ...func(*Options)) (*UpdateServiceTemplateVersionOutput, error) {
	if params == nil {
		params = &UpdateServiceTemplateVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateServiceTemplateVersion", params, optFns, c.addOperationUpdateServiceTemplateVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateServiceTemplateVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServiceTemplateVersionInput struct {

	// To update a major version of a service template, include major Version.
	//
	// This member is required.
	MajorVersion *string

	// To update a minor version of a service template, include minorVersion.
	//
	// This member is required.
	MinorVersion *string

	// The name of the service template.
	//
	// This member is required.
	TemplateName *string

	// An array of environment template objects that are compatible with this service
	// template version. A service instance based on this service template version can
	// run in environments based on compatible templates.
	CompatibleEnvironmentTemplates []types.CompatibleEnvironmentTemplateInput

	// A description of a service template version to update.
	Description *string

	// The status of the service template minor version to update.
	Status types.TemplateVersionStatus

	// An array of supported component sources. Components with supported sources can
	// be attached to service instances based on this service template version. A
	// change to supportedComponentSources doesn't impact existing component
	// attachments to instances based on this template version. A change only affects
	// later associations. For more information about components, see Proton components
	// (https://docs.aws.amazon.com/proton/latest/adminguide/ag-components.html) in the
	// Proton Administrator Guide.
	SupportedComponentSources []types.ServiceTemplateSupportedComponentSourceType

	noSmithyDocumentSerde
}

type UpdateServiceTemplateVersionOutput struct {

	// The service template version detail data that's returned by Proton.
	//
	// This member is required.
	ServiceTemplateVersion *types.ServiceTemplateVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateServiceTemplateVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateServiceTemplateVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateServiceTemplateVersion{}, middleware.After)
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateServiceTemplateVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateServiceTemplateVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateServiceTemplateVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "UpdateServiceTemplateVersion",
	}
}
