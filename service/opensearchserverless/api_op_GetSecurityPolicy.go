// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearchserverless

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearchserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about a configured OpenSearch Serverless security policy.
// For more information, see Network access for Amazon OpenSearch Serverless (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-network.html)
// and Encryption at rest for Amazon OpenSearch Serverless (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-encryption.html)
// .
func (c *Client) GetSecurityPolicy(ctx context.Context, params *GetSecurityPolicyInput, optFns ...func(*Options)) (*GetSecurityPolicyOutput, error) {
	if params == nil {
		params = &GetSecurityPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSecurityPolicy", params, optFns, c.addOperationGetSecurityPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSecurityPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSecurityPolicyInput struct {

	// The name of the security policy.
	//
	// This member is required.
	Name *string

	// The type of security policy.
	//
	// This member is required.
	Type types.SecurityPolicyType

	noSmithyDocumentSerde
}

type GetSecurityPolicyOutput struct {

	// Details about the requested security policy.
	SecurityPolicyDetail *types.SecurityPolicyDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSecurityPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetSecurityPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetSecurityPolicy{}, middleware.After)
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
	if err = addOpGetSecurityPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSecurityPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSecurityPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aoss",
		OperationName: "GetSecurityPolicy",
	}
}
