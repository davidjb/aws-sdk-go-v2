// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the results for the Trusted Advisor check summaries for the check IDs
// that you specified. You can get the check IDs by calling the
// DescribeTrustedAdvisorChecks operation. The response contains an array of
// TrustedAdvisorCheckSummary objects.
//   - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
//     use the Amazon Web Services Support API.
//   - If you call the Amazon Web Services Support API from an account that
//     doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
//     SubscriptionRequiredException error message appears. For information about
//     changing your support plan, see Amazon Web Services Support (http://aws.amazon.com/premiumsupport/)
//     .
//
// To call the Trusted Advisor operations in the Amazon Web Services Support API,
// you must use the US East (N. Virginia) endpoint. Currently, the US West (Oregon)
// and Europe (Ireland) endpoints don't support the Trusted Advisor operations. For
// more information, see About the Amazon Web Services Support API (https://docs.aws.amazon.com/awssupport/latest/user/about-support-api.html#endpoint)
// in the Amazon Web Services Support User Guide.
func (c *Client) DescribeTrustedAdvisorCheckSummaries(ctx context.Context, params *DescribeTrustedAdvisorCheckSummariesInput, optFns ...func(*Options)) (*DescribeTrustedAdvisorCheckSummariesOutput, error) {
	if params == nil {
		params = &DescribeTrustedAdvisorCheckSummariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTrustedAdvisorCheckSummaries", params, optFns, c.addOperationDescribeTrustedAdvisorCheckSummariesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTrustedAdvisorCheckSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTrustedAdvisorCheckSummariesInput struct {

	// The IDs of the Trusted Advisor checks.
	//
	// This member is required.
	CheckIds []*string

	noSmithyDocumentSerde
}

// The summaries of the Trusted Advisor checks returned by the
// DescribeTrustedAdvisorCheckSummaries operation.
type DescribeTrustedAdvisorCheckSummariesOutput struct {

	// The summary information for the requested Trusted Advisor checks.
	//
	// This member is required.
	Summaries []types.TrustedAdvisorCheckSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTrustedAdvisorCheckSummariesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTrustedAdvisorCheckSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTrustedAdvisorCheckSummaries{}, middleware.After)
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
	if err = addOpDescribeTrustedAdvisorCheckSummariesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTrustedAdvisorCheckSummaries(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTrustedAdvisorCheckSummaries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "DescribeTrustedAdvisorCheckSummaries",
	}
}
