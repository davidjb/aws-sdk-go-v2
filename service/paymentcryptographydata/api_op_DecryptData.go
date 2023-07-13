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

// Decrypts ciphertext data to plaintext using symmetric, asymmetric, or DUKPT
// data encryption key. For more information, see Decrypt data (https://docs.aws.amazon.com/payment-cryptography/latest/userguide/decrypt-data.html)
// in the Amazon Web Services Payment Cryptography User Guide. You can use an
// encryption key generated within Amazon Web Services Payment Cryptography, or you
// can import your own encryption key by calling ImportKey (https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html)
// . For this operation, the key must have KeyModesOfUse set to Decrypt . In
// asymmetric decryption, Amazon Web Services Payment Cryptography decrypts the
// ciphertext using the private component of the asymmetric encryption key pair.
// For data encryption outside of Amazon Web Services Payment Cryptography, you can
// export the public component of the asymmetric key pair by calling
// GetPublicCertificate (https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetPublicKeyCertificate.html)
// . For symmetric and DUKPT decryption, Amazon Web Services Payment Cryptography
// supports TDES and AES algorithms. For asymmetric decryption, Amazon Web
// Services Payment Cryptography supports RSA . When you use DUKPT, for TDES
// algorithm, the ciphertext data length must be a multiple of 16 bytes. For AES
// algorithm, the ciphertext data length must be a multiple of 32 bytes. For
// information about valid keys for this operation, see Understanding key
// attributes (https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html)
// and Key types for specific data operations (https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html)
// in the Amazon Web Services Payment Cryptography User Guide. Cross-account use:
// This operation can't be used across different Amazon Web Services accounts.
// Related operations:
//   - EncryptData
//   - GetPublicCertificate (https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetPublicKeyCertificate.html)
//   - ImportKey (https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html)
func (c *Client) DecryptData(ctx context.Context, params *DecryptDataInput, optFns ...func(*Options)) (*DecryptDataOutput, error) {
	if params == nil {
		params = &DecryptDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DecryptData", params, optFns, c.addOperationDecryptDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DecryptDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DecryptDataInput struct {

	// The ciphertext to decrypt.
	//
	// This member is required.
	CipherText *string

	// The encryption key type and attributes for ciphertext decryption.
	//
	// This member is required.
	DecryptionAttributes types.EncryptionDecryptionAttributes

	// The keyARN of the encryption key that Amazon Web Services Payment Cryptography
	// uses for ciphertext decryption.
	//
	// This member is required.
	KeyIdentifier *string

	noSmithyDocumentSerde
}

type DecryptDataOutput struct {

	// The keyARN of the encryption key that Amazon Web Services Payment Cryptography
	// uses for ciphertext decryption.
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

	// The decrypted plaintext data.
	//
	// This member is required.
	PlainText *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDecryptDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDecryptData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDecryptData{}, middleware.After)
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
	if err = addOpDecryptDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDecryptData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDecryptData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "payment-cryptography",
		OperationName: "DecryptData",
	}
}
