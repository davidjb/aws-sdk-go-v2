// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates Audit Manager settings for the current account.
func (c *Client) UpdateSettings(ctx context.Context, params *UpdateSettingsInput, optFns ...func(*Options)) (*UpdateSettingsOutput, error) {
	if params == nil {
		params = &UpdateSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSettings", params, optFns, c.addOperationUpdateSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSettingsInput struct {

	// The default S3 destination bucket for storing assessment reports.
	DefaultAssessmentReportsDestination *types.AssessmentReportsDestination

	// The default S3 destination bucket for storing evidence finder exports.
	DefaultExportDestination *types.DefaultExportDestination

	// A list of the default audit owners.
	DefaultProcessOwners []types.Role

	// The deregistration policy for your Audit Manager data. You can use this
	// attribute to determine how your data is handled when you deregister Audit
	// Manager.
	DeregistrationPolicy *types.DeregistrationPolicy

	// Specifies whether the evidence finder feature is enabled. Change this attribute
	// to enable or disable evidence finder. When you use this attribute to disable
	// evidence finder, Audit Manager deletes the event data store that’s used to query
	// your evidence data. As a result, you can’t re-enable evidence finder and use the
	// feature again. Your only alternative is to deregister (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeregisterAccount.html)
	// and then re-register (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_RegisterAccount.html)
	// Audit Manager.
	EvidenceFinderEnabled *bool

	// The KMS key details.
	KmsKey *string

	// The Amazon Simple Notification Service (Amazon SNS) topic that Audit Manager
	// sends notifications to.
	SnsTopic *string

	noSmithyDocumentSerde
}

type UpdateSettingsOutput struct {

	// The current list of settings.
	Settings *types.Settings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSettings{}, middleware.After)
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
	if err = addOpUpdateSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "UpdateSettings",
	}
}
