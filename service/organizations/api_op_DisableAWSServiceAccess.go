// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disables the integration of an Amazon Web Services service (the service that is
// specified by ServicePrincipal ) with Organizations. When you disable
// integration, the specified service no longer can create a service-linked role (https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html)
// in new accounts in your organization. This means the service can't perform
// operations on your behalf on any new accounts in your organization. The service
// can still perform operations in older accounts until the service completes its
// clean-up from Organizations. We strongly recommend that you don't use this
// command to disable integration between Organizations and the specified Amazon
// Web Services service. Instead, use the console or commands that are provided by
// the specified service. This lets the trusted service perform any required
// initialization when enabling trusted access, such as creating any required
// resources and any required clean up of resources when disabling trusted access.
// For information about how to disable trusted service access to your organization
// using the trusted service, see the Learn more link under the Supports Trusted
// Access column at Amazon Web Services services that you can use with
// Organizations (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services_list.html)
// . on this page. If you disable access by using this command, it causes the
// following actions to occur:
//   - The service can no longer create a service-linked role in the accounts in
//     your organization. This means that the service can't perform operations on your
//     behalf on any new accounts in your organization. The service can still perform
//     operations in older accounts until the service completes its clean-up from
//     Organizations.
//   - The service can no longer perform tasks in the member accounts in the
//     organization, unless those operations are explicitly permitted by the IAM
//     policies that are attached to your roles. This includes any data aggregation
//     from the member accounts to the management account, or to a delegated
//     administrator account, where relevant.
//   - Some services detect this and clean up any remaining data or resources
//     related to the integration, while other services stop accessing the organization
//     but leave any historical data and configuration in place to support a possible
//     re-enabling of the integration.
//
// Using the other service's console or commands to disable the integration
// ensures that the other service is aware that it can clean up any resources that
// are required only for the integration. How the service cleans up its resources
// in the organization's accounts depends on that service. For more information,
// see the documentation for the other Amazon Web Services service. After you
// perform the DisableAWSServiceAccess operation, the specified service can no
// longer perform operations in your organization's accounts For more information
// about integrating other services with Organizations, including the list of
// services that work with Organizations, see Integrating Organizations with Other
// Amazon Web Services Services (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html)
// in the Organizations User Guide. This operation can be called only from the
// organization's management account.
func (c *Client) DisableAWSServiceAccess(ctx context.Context, params *DisableAWSServiceAccessInput, optFns ...func(*Options)) (*DisableAWSServiceAccessOutput, error) {
	if params == nil {
		params = &DisableAWSServiceAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableAWSServiceAccess", params, optFns, c.addOperationDisableAWSServiceAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableAWSServiceAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableAWSServiceAccessInput struct {

	// The service principal name of the Amazon Web Services service for which you
	// want to disable integration with your organization. This is typically in the
	// form of a URL, such as service-abbreviation.amazonaws.com .
	//
	// This member is required.
	ServicePrincipal *string

	noSmithyDocumentSerde
}

type DisableAWSServiceAccessOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisableAWSServiceAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisableAWSServiceAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisableAWSServiceAccess{}, middleware.After)
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
	if err = addOpDisableAWSServiceAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableAWSServiceAccess(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisableAWSServiceAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DisableAWSServiceAccess",
	}
}
