// Code generated by smithy-go-codegen DO NOT EDIT.

package paymentcryptographydata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates a Message Authentication Code (MAC) cryptogram within Amazon Web
// Services Payment Cryptography. You can use this operation when keys won't be
// shared but mutual data is present on both ends for validation. In this case,
// known data values are used to generate a MAC on both ends for comparision
// without sending or receiving data in ciphertext or plaintext. You can use this
// operation to generate a DUPKT, HMAC or EMV MAC by setting generation attributes
// and algorithm to the associated values. The MAC generation encryption key must
// have valid values for KeyUsage such as TR31_M7_HMAC_KEY for HMAC generation,
// and they key must have KeyModesOfUse set to Generate and Verify . For
// information about valid keys for this operation, see Understanding key
// attributes (https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html)
// and Key types for specific data operations (https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html)
// in the Amazon Web Services Payment Cryptography User Guide. Cross-account use:
// This operation can't be used across different Amazon Web Services accounts.
// Related operations:
//   - VerifyMac
func (c *Client) GenerateMac(ctx context.Context, params *GenerateMacInput, optFns ...func(*Options)) (*GenerateMacOutput, error) {
	if params == nil {
		params = &GenerateMacInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GenerateMac", params, optFns, c.addOperationGenerateMacMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GenerateMacOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateMacInput struct {

	// The attributes and data values to use for MAC generation within Amazon Web
	// Services Payment Cryptography.
	//
	// This member is required.
	GenerationAttributes types.MacAttributes

	// The keyARN of the MAC generation encryption key.
	//
	// This member is required.
	KeyIdentifier *string

	// The data for which a MAC is under generation.
	//
	// This member is required.
	MessageData *string

	// The length of a MAC under generation.
	MacLength *int32

	noSmithyDocumentSerde
}

type GenerateMacOutput struct {

	// The keyARN of the encryption key that Amazon Web Services Payment Cryptography
	// uses for MAC generation.
	//
	// This member is required.
	KeyArn *string

	// The key check value (KCV) of the encryption key. The KCV is used to check if
	// all parties holding a given key have the same key or to detect that a key has
	// changed. Amazon Web Services Payment Cryptography calculates the KCV by using
	// standard algorithms, typically by encrypting 8 or 16 bytes or "00" or "01" and
	// then truncating the result to the first 3 bytes, or 6 hex digits, of the
	// resulting cryptogram.
	//
	// This member is required.
	KeyCheckValue *string

	// The MAC cryptogram generated within Amazon Web Services Payment Cryptography.
	//
	// This member is required.
	Mac *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGenerateMacMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGenerateMac{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGenerateMac{}, middleware.After)
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
	if err = addOpGenerateMacValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateMac(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGenerateMac(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "payment-cryptography",
		OperationName: "GenerateMac",
	}
}
