// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html)
// . With the latest version, AWS WAF has a single set of endpoints for regional
// and global use. Creates a RateBasedRule . The RateBasedRule contains a RateLimit
// , which specifies the maximum number of requests that AWS WAF allows from a
// specified IP address in a five-minute period. The RateBasedRule also contains
// the IPSet objects, ByteMatchSet objects, and other predicates that identify the
// requests that you want to count or block if these requests exceed the RateLimit
// . If you add more than one predicate to a RateBasedRule , a request not only
// must exceed the RateLimit , but it also must match all the conditions to be
// counted or blocked. For example, suppose you add the following to a
// RateBasedRule :
//   - An IPSet that matches the IP address 192.0.2.44/32
//   - A ByteMatchSet that matches BadBot in the User-Agent header
//
// Further, you specify a RateLimit of 1,000. You then add the RateBasedRule to a
// WebACL and specify that you want to block requests that meet the conditions in
// the rule. For a request to be blocked, it must come from the IP address
// 192.0.2.44 and the User-Agent header in the request must contain the value
// BadBot . Further, requests that match these two conditions must be received at a
// rate of more than 1,000 requests every five minutes. If both conditions are met
// and the rate is exceeded, AWS WAF blocks the requests. If the rate drops below
// 1,000 for a five-minute period, AWS WAF no longer blocks the requests. As a
// second example, suppose you want to limit requests to a particular page on your
// site. To do this, you could add the following to a RateBasedRule :
//   - A ByteMatchSet with FieldToMatch of URI
//   - A PositionalConstraint of STARTS_WITH
//   - A TargetString of login
//
// Further, you specify a RateLimit of 1,000. By adding this RateBasedRule to a
// WebACL , you could limit requests to your login page without affecting the rest
// of your site. To create and configure a RateBasedRule , perform the following
// steps:
//   - Create and update the predicates that you want to include in the rule. For
//     more information, see CreateByteMatchSet , CreateIPSet , and
//     CreateSqlInjectionMatchSet .
//   - Use GetChangeToken to get the change token that you provide in the
//     ChangeToken parameter of a CreateRule request.
//   - Submit a CreateRateBasedRule request.
//   - Use GetChangeToken to get the change token that you provide in the
//     ChangeToken parameter of an UpdateRule request.
//   - Submit an UpdateRateBasedRule request to specify the predicates that you
//     want to include in the rule.
//   - Create and update a WebACL that contains the RateBasedRule . For more
//     information, see CreateWebACL .
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/)
// .
func (c *Client) CreateRateBasedRule(ctx context.Context, params *CreateRateBasedRuleInput, optFns ...func(*Options)) (*CreateRateBasedRuleOutput, error) {
	if params == nil {
		params = &CreateRateBasedRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRateBasedRule", params, optFns, c.addOperationCreateRateBasedRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRateBasedRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRateBasedRuleInput struct {

	// The ChangeToken that you used to submit the CreateRateBasedRule request. You
	// can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus .
	//
	// This member is required.
	ChangeToken *string

	// A friendly name or description for the metrics for this RateBasedRule . The name
	// can contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length
	// 128 and minimum length one. It can't contain whitespace or metric names reserved
	// for AWS WAF, including "All" and "Default_Action." You can't change the name of
	// the metric after you create the RateBasedRule .
	//
	// This member is required.
	MetricName *string

	// A friendly name or description of the RateBasedRule . You can't change the name
	// of a RateBasedRule after you create it.
	//
	// This member is required.
	Name *string

	// The field that AWS WAF uses to determine if requests are likely arriving from a
	// single source and thus subject to rate monitoring. The only valid value for
	// RateKey is IP . IP indicates that requests that arrive from the same IP address
	// are subject to the RateLimit that is specified in the RateBasedRule .
	//
	// This member is required.
	RateKey types.RateKey

	// The maximum number of requests, which have an identical value in the field that
	// is specified by RateKey , allowed in a five-minute period. If the number of
	// requests exceeds the RateLimit and the other predicates specified in the rule
	// are also met, AWS WAF triggers the action that is specified for this rule.
	//
	// This member is required.
	RateLimit int64

	//
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateRateBasedRuleOutput struct {

	// The ChangeToken that you used to submit the CreateRateBasedRule request. You
	// can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus .
	ChangeToken *string

	// The RateBasedRule that is returned in the CreateRateBasedRule response.
	Rule *types.RateBasedRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRateBasedRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRateBasedRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRateBasedRule{}, middleware.After)
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
	if err = addOpCreateRateBasedRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRateBasedRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRateBasedRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "CreateRateBasedRule",
	}
}
