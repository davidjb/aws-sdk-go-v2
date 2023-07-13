// Code generated by smithy-go-codegen DO NOT EDIT.

package iottwinmaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iottwinmaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a component type.
func (c *Client) CreateComponentType(ctx context.Context, params *CreateComponentTypeInput, optFns ...func(*Options)) (*CreateComponentTypeOutput, error) {
	if params == nil {
		params = &CreateComponentTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateComponentType", params, optFns, c.addOperationCreateComponentTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateComponentTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateComponentTypeInput struct {

	// The ID of the component type.
	//
	// This member is required.
	ComponentTypeId *string

	// The ID of the workspace that contains the component type.
	//
	// This member is required.
	WorkspaceId *string

	// A friendly name for the component type.
	ComponentTypeName *string

	// The description of the component type.
	Description *string

	// Specifies the parent component type to extend.
	ExtendsFrom []string

	// An object that maps strings to the functions in the component type. Each string
	// in the mapping must be unique to this object.
	Functions map[string]types.FunctionRequest

	// A Boolean value that specifies whether an entity can have more than one
	// component of this type.
	IsSingleton *bool

	// An object that maps strings to the property definitions in the component type.
	// Each string in the mapping must be unique to this object.
	PropertyDefinitions map[string]types.PropertyDefinitionRequest

	//
	PropertyGroups map[string]types.PropertyGroupRequest

	// Metadata that you can use to manage the component type.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateComponentTypeOutput struct {

	// The ARN of the component type.
	//
	// This member is required.
	Arn *string

	// The date and time when the entity was created.
	//
	// This member is required.
	CreationDateTime *time.Time

	// The current state of the component type.
	//
	// This member is required.
	State types.State

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateComponentTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateComponentType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateComponentType{}, middleware.After)
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
	if err = addEndpointPrefix_opCreateComponentTypeMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateComponentTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateComponentType(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCreateComponentTypeMiddleware struct {
}

func (*endpointPrefix_opCreateComponentTypeMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateComponentTypeMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opCreateComponentTypeMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opCreateComponentTypeMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opCreateComponentType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iottwinmaker",
		OperationName: "CreateComponentType",
	}
}
