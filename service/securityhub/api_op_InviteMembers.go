// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Invites other Amazon Web Services accounts to become member accounts for the
// Security Hub administrator account that the invitation is sent from. This
// operation is only used to invite accounts that do not belong to an organization.
// Organization accounts do not receive invitations. Before you can use this action
// to invite a member, you must first use the CreateMembers action to create the
// member account in Security Hub. When the account owner enables Security Hub and
// accepts the invitation to become a member account, the administrator account can
// view the findings generated from the member account.
func (c *Client) InviteMembers(ctx context.Context, params *InviteMembersInput, optFns ...func(*Options)) (*InviteMembersOutput, error) {
	if params == nil {
		params = &InviteMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InviteMembers", params, optFns, c.addOperationInviteMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InviteMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InviteMembersInput struct {

	// The list of account IDs of the Amazon Web Services accounts to invite to
	// Security Hub as members.
	//
	// This member is required.
	AccountIds []string

	noSmithyDocumentSerde
}

type InviteMembersOutput struct {

	// The list of Amazon Web Services accounts that could not be processed. For each
	// account, the list includes the account ID and the email address.
	UnprocessedAccounts []types.Result

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationInviteMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInviteMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInviteMembers{}, middleware.After)
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
	if err = addOpInviteMembersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInviteMembers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInviteMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "InviteMembers",
	}
}
