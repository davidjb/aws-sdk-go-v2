// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Unarchive application.
func (c *Client) UnarchiveApplication(ctx context.Context, params *UnarchiveApplicationInput, optFns ...func(*Options)) (*UnarchiveApplicationOutput, error) {
	if params == nil {
		params = &UnarchiveApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UnarchiveApplication", params, optFns, c.addOperationUnarchiveApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UnarchiveApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UnarchiveApplicationInput struct {

	// Application ID.
	//
	// This member is required.
	ApplicationID *string

	// Account ID.
	AccountID *string

	noSmithyDocumentSerde
}

type UnarchiveApplicationOutput struct {

	// Application aggregated status.
	ApplicationAggregatedStatus *types.ApplicationAggregatedStatus

	// Application ID.
	ApplicationID *string

	// Application ARN.
	Arn *string

	// Application creation dateTime.
	CreationDateTime *string

	// Application description.
	Description *string

	// Application archival status.
	IsArchived *bool

	// Application last modified dateTime.
	LastModifiedDateTime *string

	// Application name.
	Name *string

	// Application tags.
	Tags map[string]string

	// Application wave ID.
	WaveID *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUnarchiveApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUnarchiveApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUnarchiveApplication{}, middleware.After)
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
	if err = addOpUnarchiveApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUnarchiveApplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUnarchiveApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgn",
		OperationName: "UnarchiveApplication",
	}
}
