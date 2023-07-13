// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Applies an Amazon EFS FileSystemPolicy to an Amazon EFS file system. A file
// system policy is an IAM resource-based policy and can contain multiple policy
// statements. A file system always has exactly one file system policy, which can
// be the default policy or an explicit policy set or updated using this API
// operation. EFS file system policies have a 20,000 character limit. When an
// explicit policy is set, it overrides the default policy. For more information
// about the default file system policy, see Default EFS File System Policy (https://docs.aws.amazon.com/efs/latest/ug/iam-access-control-nfs-efs.html#default-filesystempolicy)
// . EFS file system policies have a 20,000 character limit. This operation
// requires permissions for the elasticfilesystem:PutFileSystemPolicy action.
func (c *Client) PutFileSystemPolicy(ctx context.Context, params *PutFileSystemPolicyInput, optFns ...func(*Options)) (*PutFileSystemPolicyOutput, error) {
	if params == nil {
		params = &PutFileSystemPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutFileSystemPolicy", params, optFns, c.addOperationPutFileSystemPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutFileSystemPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutFileSystemPolicyInput struct {

	// The ID of the EFS file system that you want to create or update the
	// FileSystemPolicy for.
	//
	// This member is required.
	FileSystemId *string

	// The FileSystemPolicy that you're creating. Accepts a JSON formatted policy
	// definition. EFS file system policies have a 20,000 character limit. To find out
	// more about the elements that make up a file system policy, see EFS
	// Resource-based Policies (https://docs.aws.amazon.com/efs/latest/ug/access-control-overview.html#access-control-manage-access-intro-resource-policies)
	// .
	//
	// This member is required.
	Policy *string

	// (Optional) A boolean that specifies whether or not to bypass the
	// FileSystemPolicy lockout safety check. The lockout safety check determines
	// whether the policy in the request will lock out, or prevent, the IAM principal
	// that is making the request from making future PutFileSystemPolicy requests on
	// this file system. Set BypassPolicyLockoutSafetyCheck to True only when you
	// intend to prevent the IAM principal that is making the request from making
	// subsequent PutFileSystemPolicy requests on this file system. The default value
	// is False .
	BypassPolicyLockoutSafetyCheck bool

	noSmithyDocumentSerde
}

type PutFileSystemPolicyOutput struct {

	// Specifies the EFS file system to which the FileSystemPolicy applies.
	FileSystemId *string

	// The JSON formatted FileSystemPolicy for the EFS file system.
	Policy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutFileSystemPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutFileSystemPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutFileSystemPolicy{}, middleware.After)
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
	if err = addOpPutFileSystemPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutFileSystemPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutFileSystemPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "PutFileSystemPolicy",
	}
}
