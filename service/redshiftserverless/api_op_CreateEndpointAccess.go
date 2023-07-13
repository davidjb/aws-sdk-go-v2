// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftserverless

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Redshift Serverless managed VPC endpoint.
func (c *Client) CreateEndpointAccess(ctx context.Context, params *CreateEndpointAccessInput, optFns ...func(*Options)) (*CreateEndpointAccessOutput, error) {
	if params == nil {
		params = &CreateEndpointAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEndpointAccess", params, optFns, c.addOperationCreateEndpointAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEndpointAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEndpointAccessInput struct {

	// The name of the VPC endpoint. An endpoint name must contain 1-30 characters.
	// Valid characters are A-Z, a-z, 0-9, and hyphen(-). The first character must be a
	// letter. The name can't contain two consecutive hyphens or end with a hyphen.
	//
	// This member is required.
	EndpointName *string

	// The unique identifers of subnets from which Amazon Redshift Serverless chooses
	// one to deploy a VPC endpoint.
	//
	// This member is required.
	SubnetIds []string

	// The name of the workgroup to associate with the VPC endpoint.
	//
	// This member is required.
	WorkgroupName *string

	// The unique identifiers of the security group that defines the ports, protocols,
	// and sources for inbound traffic that you are authorizing into your endpoint.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type CreateEndpointAccessOutput struct {

	// The created VPC endpoint.
	Endpoint *types.EndpointAccess

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEndpointAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEndpointAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEndpointAccess{}, middleware.After)
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
	if err = addOpCreateEndpointAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEndpointAccess(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEndpointAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift-serverless",
		OperationName: "CreateEndpointAccess",
	}
}
