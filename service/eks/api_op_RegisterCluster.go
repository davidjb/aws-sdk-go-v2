// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Connects a Kubernetes cluster to the Amazon EKS control plane. Any Kubernetes
// cluster can be connected to the Amazon EKS control plane to view current
// information about the cluster and its nodes. Cluster connection requires two
// steps. First, send a RegisterClusterRequest to add it to the Amazon EKS control
// plane. Second, a Manifest (https://amazon-eks.s3.us-west-2.amazonaws.com/eks-connector/manifests/eks-connector/latest/eks-connector.yaml)
// containing the activationID and activationCode must be applied to the
// Kubernetes cluster through it's native provider to provide visibility. After the
// Manifest is updated and applied, then the connected cluster is visible to the
// Amazon EKS control plane. If the Manifest is not applied within three days, then
// the connected cluster will no longer be visible and must be deregistered. See
// DeregisterCluster .
func (c *Client) RegisterCluster(ctx context.Context, params *RegisterClusterInput, optFns ...func(*Options)) (*RegisterClusterOutput, error) {
	if params == nil {
		params = &RegisterClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterCluster", params, optFns, c.addOperationRegisterClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterClusterInput struct {

	// The configuration settings required to connect the Kubernetes cluster to the
	// Amazon EKS control plane.
	//
	// This member is required.
	ConnectorConfig *types.ConnectorConfigRequest

	// Define a unique name for this cluster for your Region.
	//
	// This member is required.
	Name *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// The metadata that you apply to the cluster to assist with categorization and
	// organization. Each tag consists of a key and an optional value, both of which
	// you define. Cluster tags do not propagate to any other resources associated with
	// the cluster.
	Tags map[string]string

	noSmithyDocumentSerde
}

type RegisterClusterOutput struct {

	// An object representing an Amazon EKS cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterCluster{}, middleware.After)
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
	if err = addIdempotencyToken_opRegisterClusterMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpRegisterClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterCluster(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpRegisterCluster struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpRegisterCluster) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpRegisterCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*RegisterClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *RegisterClusterInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opRegisterClusterMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpRegisterCluster{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opRegisterCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "RegisterCluster",
	}
}
