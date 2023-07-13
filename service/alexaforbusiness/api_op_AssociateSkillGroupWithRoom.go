// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a skill group with a given room. This enables all skills in the
// associated skill group on all devices in the room.
//
// Deprecated: Alexa For Business is no longer supported
func (c *Client) AssociateSkillGroupWithRoom(ctx context.Context, params *AssociateSkillGroupWithRoomInput, optFns ...func(*Options)) (*AssociateSkillGroupWithRoomOutput, error) {
	if params == nil {
		params = &AssociateSkillGroupWithRoomInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateSkillGroupWithRoom", params, optFns, c.addOperationAssociateSkillGroupWithRoomMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateSkillGroupWithRoomOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateSkillGroupWithRoomInput struct {

	// The ARN of the room with which to associate the skill group. Required.
	RoomArn *string

	// The ARN of the skill group to associate with a room. Required.
	SkillGroupArn *string

	noSmithyDocumentSerde
}

type AssociateSkillGroupWithRoomOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateSkillGroupWithRoomMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateSkillGroupWithRoom{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateSkillGroupWithRoom{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateSkillGroupWithRoom(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateSkillGroupWithRoom(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "AssociateSkillGroupWithRoom",
	}
}
