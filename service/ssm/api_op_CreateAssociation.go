// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A State Manager association defines the state that you want to maintain on your
// managed nodes. For example, an association can specify that anti-virus software
// must be installed and running on your managed nodes, or that certain ports must
// be closed. For static targets, the association specifies a schedule for when the
// configuration is reapplied. For dynamic targets, such as an Amazon Web Services
// resource group or an Amazon Web Services autoscaling group, State Manager, a
// capability of Amazon Web Services Systems Manager applies the configuration when
// new managed nodes are added to the group. The association also specifies actions
// to take when applying the configuration. For example, an association for
// anti-virus software might run once a day. If the software isn't installed, then
// State Manager installs it. If the software is installed, but the service isn't
// running, then the association might instruct State Manager to start the service.
func (c *Client) CreateAssociation(ctx context.Context, params *CreateAssociationInput, optFns ...func(*Options)) (*CreateAssociationOutput, error) {
	if params == nil {
		params = &CreateAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAssociation", params, optFns, c.addOperationCreateAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAssociationInput struct {

	// The name of the SSM Command document or Automation runbook that contains the
	// configuration information for the managed node. You can specify Amazon Web
	// Services-predefined documents, documents you created, or a document that is
	// shared with you from another Amazon Web Services account. For Systems Manager
	// documents (SSM documents) that are shared with you from other Amazon Web
	// Services accounts, you must specify the complete SSM document ARN, in the
	// following format: arn:partition:ssm:region:account-id:document/document-name
	// For example: arn:aws:ssm:us-east-2:12345678912:document/My-Shared-Document For
	// Amazon Web Services-predefined documents and SSM documents you created in your
	// account, you only need to specify the document name. For example,
	// AWS-ApplyPatchBaseline or My-Document .
	//
	// This member is required.
	Name *string

	// The details for the CloudWatch alarm you want to apply to an automation or
	// command.
	AlarmConfiguration *types.AlarmConfiguration

	// By default, when you create a new association, the system runs it immediately
	// after it is created and then according to the schedule you specified. Specify
	// this option if you don't want an association to run immediately after you create
	// it. This parameter isn't supported for rate expressions.
	ApplyOnlyAtCronInterval bool

	// Specify a descriptive name for the association.
	AssociationName *string

	// Choose the parameter that will define how your automation will branch out. This
	// target is required for associations that use an Automation runbook and target
	// resources by using rate controls. Automation is a capability of Amazon Web
	// Services Systems Manager.
	AutomationTargetParameterName *string

	// The names or Amazon Resource Names (ARNs) of the Change Calendar type documents
	// you want to gate your associations under. The associations only run when that
	// change calendar is open. For more information, see Amazon Web Services Systems
	// Manager Change Calendar (https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-change-calendar)
	// .
	CalendarNames []string

	// The severity level to assign to the association.
	ComplianceSeverity types.AssociationComplianceSeverity

	// The document version you want to associate with the target(s). Can be a
	// specific version or the default version. State Manager doesn't support running
	// associations that use a new version of a document if that document is shared
	// from another account. State Manager always runs the default version of a
	// document if shared from another account, even though the Systems Manager console
	// shows that a new version was processed. If you want to run an association using
	// a new version of a document shared form another account, you must set the
	// document version to default .
	DocumentVersion *string

	// The managed node ID. InstanceId has been deprecated. To specify a managed node
	// ID for an association, use the Targets parameter. Requests that include the
	// parameter InstanceID with Systems Manager documents (SSM documents) that use
	// schema version 2.0 or later will fail. In addition, if you use the parameter
	// InstanceId , you can't use the parameters AssociationName , DocumentVersion ,
	// MaxErrors , MaxConcurrency , OutputLocation , or ScheduleExpression . To use
	// these parameters, you must use the Targets parameter.
	InstanceId *string

	// The maximum number of targets allowed to run the association at the same time.
	// You can specify a number, for example 10, or a percentage of the target set, for
	// example 10%. The default value is 100%, which means all targets run the
	// association at the same time. If a new managed node starts and attempts to run
	// an association while Systems Manager is running MaxConcurrency associations,
	// the association is allowed to run. During the next association interval, the new
	// managed node will process its association within the limit specified for
	// MaxConcurrency .
	MaxConcurrency *string

	// The number of errors that are allowed before the system stops sending requests
	// to run the association on additional targets. You can specify either an absolute
	// number of errors, for example 10, or a percentage of the target set, for example
	// 10%. If you specify 3, for example, the system stops sending requests when the
	// fourth error is received. If you specify 0, then the system stops sending
	// requests after the first error is returned. If you run an association on 50
	// managed nodes and set MaxError to 10%, then the system stops sending the
	// request when the sixth error is received. Executions that are already running an
	// association when MaxErrors is reached are allowed to complete, but some of
	// these executions may fail as well. If you need to ensure that there won't be
	// more than max-errors failed executions, set MaxConcurrency to 1 so that
	// executions proceed one at a time.
	MaxErrors *string

	// An Amazon Simple Storage Service (Amazon S3) bucket where you want to store the
	// output details of the request.
	OutputLocation *types.InstanceAssociationOutputLocation

	// The parameters for the runtime configuration of the document.
	Parameters map[string][]string

	// A cron expression when the association will be applied to the target(s).
	ScheduleExpression *string

	// Number of days to wait after the scheduled day to run an association. For
	// example, if you specified a cron schedule of cron(0 0 ? * THU#2 *) , you could
	// specify an offset of 3 to run the association each Sunday after the second
	// Thursday of the month. For more information about cron schedules for
	// associations, see Reference: Cron and rate expressions for Systems Manager (https://docs.aws.amazon.com/systems-manager/latest/userguide/reference-cron-and-rate-expressions.html)
	// in the Amazon Web Services Systems Manager User Guide. To use offsets, you must
	// specify the ApplyOnlyAtCronInterval parameter. This option tells the system not
	// to run an association immediately after you create it.
	ScheduleOffset *int32

	// The mode for generating association compliance. You can specify AUTO or MANUAL .
	// In AUTO mode, the system uses the status of the association execution to
	// determine the compliance status. If the association execution runs successfully,
	// then the association is COMPLIANT . If the association execution doesn't run
	// successfully, the association is NON-COMPLIANT . In MANUAL mode, you must
	// specify the AssociationId as a parameter for the PutComplianceItems API
	// operation. In this case, compliance data isn't managed by State Manager. It is
	// managed by your direct call to the PutComplianceItems API operation. By
	// default, all associations use AUTO mode.
	SyncCompliance types.AssociationSyncCompliance

	// Adds or overwrites one or more tags for a State Manager association. Tags are
	// metadata that you can assign to your Amazon Web Services resources. Tags enable
	// you to categorize your resources in different ways, for example, by purpose,
	// owner, or environment. Each tag consists of a key and an optional value, both of
	// which you define.
	Tags []types.Tag

	// A location is a combination of Amazon Web Services Regions and Amazon Web
	// Services accounts where you want to run the association. Use this action to
	// create an association in multiple Regions and multiple accounts.
	TargetLocations []types.TargetLocation

	// A key-value mapping of document parameters to target resources. Both Targets
	// and TargetMaps can't be specified together.
	TargetMaps []map[string][]string

	// The targets for the association. You can target managed nodes by using tags,
	// Amazon Web Services resource groups, all managed nodes in an Amazon Web Services
	// account, or individual managed node IDs. You can target all managed nodes in an
	// Amazon Web Services account by specifying the InstanceIds key with a value of *
	// . For more information about choosing targets for an association, see Using
	// targets and rate controls with State Manager associations (https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-state-manager-targets-and-rate-controls.html)
	// in the Amazon Web Services Systems Manager User Guide.
	Targets []types.Target

	noSmithyDocumentSerde
}

type CreateAssociationOutput struct {

	// Information about the association.
	AssociationDescription *types.AssociationDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAssociation{}, middleware.After)
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
	if err = addOpCreateAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "CreateAssociation",
	}
}
