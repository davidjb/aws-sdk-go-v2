// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/endpoints/awsrulesfn"
	internalendpoints "github.com/aws/aws-sdk-go-v2/service/eventbridge/internal/endpoints"
	smithy "github.com/aws/smithy-go"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/endpoints/private/rulesfn"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"net/http"
	"net/url"
	"strings"
)

// EndpointResolverOptions is the service endpoint resolver options
type EndpointResolverOptions = internalendpoints.Options

// EndpointResolver interface for resolving service endpoints.
type EndpointResolver interface {
	ResolveEndpoint(region string, options EndpointResolverOptions) (aws.Endpoint, error)
}

var _ EndpointResolver = &internalendpoints.Resolver{}

// NewDefaultEndpointResolver constructs a new service endpoint resolver
func NewDefaultEndpointResolver() *internalendpoints.Resolver {
	return internalendpoints.New()
}

// EndpointResolverFunc is a helper utility that wraps a function so it satisfies
// the EndpointResolver interface. This is useful when you want to add additional
// endpoint resolving logic, or stub out specific endpoints with custom values.
type EndpointResolverFunc func(region string, options EndpointResolverOptions) (aws.Endpoint, error)

func (fn EndpointResolverFunc) ResolveEndpoint(region string, options EndpointResolverOptions) (endpoint aws.Endpoint, err error) {
	return fn(region, options)
}

func resolveDefaultEndpointConfiguration(o *Options) {
	if o.EndpointResolver != nil {
		return
	}
	o.EndpointResolver = &compatibleEndpointResolver{
		EndpointResolverV2: NewDefaultEndpointResolverV2(),
	}
}

// EndpointResolverFromURL returns an EndpointResolver configured using the
// provided endpoint url. By default, the resolved endpoint resolver uses the
// client region as signing region, and the endpoint source is set to
// EndpointSourceCustom.You can provide functional options to configure endpoint
// values for the resolved endpoint.
func EndpointResolverFromURL(url string, optFns ...func(*aws.Endpoint)) EndpointResolver {
	e := aws.Endpoint{URL: url, Source: aws.EndpointSourceCustom}
	for _, fn := range optFns {
		fn(&e)
	}

	return EndpointResolverFunc(
		func(region string, options EndpointResolverOptions) (aws.Endpoint, error) {
			if len(e.SigningRegion) == 0 {
				e.SigningRegion = region
			}
			return e, nil
		},
	)
}

type ResolveEndpoint struct {
	Resolver EndpointResolver
	Options  EndpointResolverOptions
}

func (*ResolveEndpoint) ID() string {
	return "ResolveEndpoint"
}

func (m *ResolveEndpoint) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.Resolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	eo := m.Options
	eo.Logger = middleware.GetLogger(ctx)

	var endpoint aws.Endpoint
	endpoint, err = m.Resolver.ResolveEndpoint(awsmiddleware.GetRegion(ctx), eo)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL, err = url.Parse(endpoint.URL)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to parse endpoint URL: %w", err)
	}

	if len(awsmiddleware.GetSigningName(ctx)) == 0 {
		signingName := endpoint.SigningName
		if len(signingName) == 0 {
			signingName = "events"
		}
		ctx = awsmiddleware.SetSigningName(ctx, signingName)
	}
	ctx = awsmiddleware.SetEndpointSource(ctx, endpoint.Source)
	ctx = smithyhttp.SetHostnameImmutable(ctx, endpoint.HostnameImmutable)
	ctx = awsmiddleware.SetSigningRegion(ctx, endpoint.SigningRegion)
	ctx = awsmiddleware.SetPartitionID(ctx, endpoint.PartitionID)
	return next.HandleSerialize(ctx, in)
}
func addResolveEndpointMiddleware(stack *middleware.Stack, o Options) error {
	return stack.Serialize.Insert(&ResolveEndpoint{
		Resolver: o.EndpointResolver,
		Options:  o.EndpointOptions,
	}, "OperationSerializer", middleware.Before)
}

func removeResolveEndpointMiddleware(stack *middleware.Stack) error {
	_, err := stack.Serialize.Remove((&ResolveEndpoint{}).ID())
	return err
}

type wrappedEndpointResolver struct {
	awsResolver aws.EndpointResolverWithOptions
	resolver    EndpointResolver
}

func (w *wrappedEndpointResolver) ResolveEndpoint(region string, options EndpointResolverOptions) (endpoint aws.Endpoint, err error) {
	if w.awsResolver == nil {
		goto fallback
	}
	endpoint, err = w.awsResolver.ResolveEndpoint(ServiceID, region, options)
	if err == nil {
		return endpoint, nil
	}

	if nf := (&aws.EndpointNotFoundError{}); !errors.As(err, &nf) {
		return endpoint, err
	}

fallback:
	if w.resolver == nil {
		return endpoint, fmt.Errorf("default endpoint resolver provided was nil")
	}
	return w.resolver.ResolveEndpoint(region, options)
}

type awsEndpointResolverAdaptor func(service, region string) (aws.Endpoint, error)

func (a awsEndpointResolverAdaptor) ResolveEndpoint(service, region string, options ...interface{}) (aws.Endpoint, error) {
	return a(service, region)
}

var _ aws.EndpointResolverWithOptions = awsEndpointResolverAdaptor(nil)

// withEndpointResolver returns an EndpointResolver that first delegates endpoint resolution to the awsResolver.
// If awsResolver returns aws.EndpointNotFoundError error, the resolver will use the the provided
// fallbackResolver for resolution.
//
// fallbackResolver must not be nil
func withEndpointResolver(awsResolver aws.EndpointResolver, awsResolverWithOptions aws.EndpointResolverWithOptions, fallbackResolver EndpointResolver) EndpointResolver {
	var resolver aws.EndpointResolverWithOptions

	if awsResolverWithOptions != nil {
		resolver = awsResolverWithOptions
	} else if awsResolver != nil {
		resolver = awsEndpointResolverAdaptor(awsResolver.ResolveEndpoint)
	}

	return &wrappedEndpointResolver{
		awsResolver: resolver,
		resolver:    fallbackResolver,
	}
}

func finalizeClientEndpointResolverOptions(options *Options) {
	options.EndpointOptions.LogDeprecated = options.ClientLogMode.IsDeprecatedUsage()

	if len(options.EndpointOptions.ResolvedRegion) == 0 {
		const fipsInfix = "-fips-"
		const fipsPrefix = "fips-"
		const fipsSuffix = "-fips"

		if strings.Contains(options.Region, fipsInfix) ||
			strings.Contains(options.Region, fipsPrefix) ||
			strings.Contains(options.Region, fipsSuffix) {
			options.EndpointOptions.ResolvedRegion = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(
				options.Region, fipsInfix, "-"), fipsPrefix, ""), fipsSuffix, "")
			options.EndpointOptions.UseFIPSEndpoint = aws.FIPSEndpointStateEnabled
		}
	}

}

type legacyEndpointResolverAdapter struct {
	legacyResolver EndpointResolver
	resolver       EndpointResolverV2
}

func (l *legacyEndpointResolverAdapter) ResolveEndpoint(ctx context.Context, params EndpointParameters) (endpoint smithyendpoints.Endpoint, err error) {
	var fips aws.FIPSEndpointState
	var dualStack aws.DualStackEndpointState

	if aws.ToBool(params.UseFIPS) {
		fips = aws.FIPSEndpointStateEnabled
	}
	if aws.ToBool(params.UseDualStack) {
		dualStack = aws.DualStackEndpointStateEnabled
	}

	resolveEndpoint, err := l.legacyResolver.ResolveEndpoint(aws.ToString(params.Region), EndpointResolverOptions{
		ResolvedRegion:       aws.ToString(params.Region),
		UseFIPSEndpoint:      fips,
		UseDualStackEndpoint: dualStack,
	})
	if err != nil {
		return endpoint, err
	}

	if resolveEndpoint.HostnameImmutable {
		uriString := resolveEndpoint.URL
		uri, err := url.Parse(uriString)
		if err != nil {
			return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
		}

		return smithyendpoints.Endpoint{
			URI: *uri,
		}, nil
	}

	if resolveEndpoint.Source == aws.EndpointSourceServiceMetadata {
		return l.resolver.ResolveEndpoint(ctx, params)
	}

	params = params.WithDefaults()
	params.Endpoint = &resolveEndpoint.URL

	return l.resolver.ResolveEndpoint(ctx, params)
}

type isDefaultProvidedImplementation interface {
	isDefaultProvidedImplementation()
}

type compatibleEndpointResolver struct {
	EndpointResolverV2 EndpointResolverV2
}

func (n *compatibleEndpointResolver) isDefaultProvidedImplementation() {}

func (n *compatibleEndpointResolver) ResolveEndpoint(region string, options EndpointResolverOptions) (endpoint aws.Endpoint, err error) {
	reg := region
	fips := options.UseFIPSEndpoint
	if len(options.ResolvedRegion) > 0 {
		reg = options.ResolvedRegion
	} else {
		// EndpointResolverV2 needs to support pseudo-regions to maintain backwards-compatibility
		// with the legacy EndpointResolver
		reg, fips = mapPseudoRegion(region)
	}
	ctx := context.Background()
	resolved, err := n.EndpointResolverV2.ResolveEndpoint(ctx, EndpointParameters{
		Region:       &reg,
		UseFIPS:      aws.Bool(fips == aws.FIPSEndpointStateEnabled),
		UseDualStack: aws.Bool(options.UseDualStackEndpoint == aws.DualStackEndpointStateEnabled),
	})
	if err != nil {
		return endpoint, err
	}

	endpoint = aws.Endpoint{
		URL:               resolved.URI.String(),
		HostnameImmutable: false,
		Source:            aws.EndpointSourceServiceMetadata,
	}

	return endpoint, nil
}

// Utility function to aid with translating pseudo-regions to classical regions
// with the appropriate setting indicated by the pseudo-region
func mapPseudoRegion(pr string) (region string, fips aws.FIPSEndpointState) {
	const fipsInfix = "-fips-"
	const fipsPrefix = "fips-"
	const fipsSuffix = "-fips"

	if strings.Contains(pr, fipsInfix) ||
		strings.Contains(pr, fipsPrefix) ||
		strings.Contains(pr, fipsSuffix) {
		region = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(
			pr, fipsInfix, "-"), fipsPrefix, ""), fipsSuffix, "")
		fips = aws.FIPSEndpointStateEnabled
	} else {
		region = pr
	}

	return region, fips
}

func finalizeEndpointResolverV2(options *Options) {
	// Check if the EndpointResolver was not user provided
	// but is the SDK's default provided version.
	_, ok := options.EndpointResolver.(isDefaultProvidedImplementation)
	if options.EndpointResolverV2 == nil {
		options.EndpointResolverV2 = NewDefaultEndpointResolverV2()
	}
	if ok {
		// Nothing further to do
		return
	}

	options.EndpointResolverV2 = &legacyEndpointResolverAdapter{
		legacyResolver: options.EndpointResolver,
		resolver:       NewDefaultEndpointResolverV2(),
	}
}

// BuiltInParameterResolver is the interface responsible for resolving BuiltIn
// values during the sourcing of EndpointParameters
type BuiltInParameterResolver interface {
	ResolveBuiltIns(*EndpointParameters) error
}

// BuiltInResolver resolves modeled BuiltIn values using only the members defined
// below.
type BuiltInResolver struct {
	// The AWS region used to dispatch the request.
	Region string

	// Sourced BuiltIn value in a historical enabled or disabled state.
	UseDualStack aws.DualStackEndpointState

	// Sourced BuiltIn value in a historical enabled or disabled state.
	UseFIPS aws.FIPSEndpointState

	// Base endpoint that can potentially be modified during Endpoint resolution.
	Endpoint *string
}

// Invoked at runtime to resolve BuiltIn Values. Only resolution code specific to
// each BuiltIn value is generated.
func (b *BuiltInResolver) ResolveBuiltIns(params *EndpointParameters) error {

	region, _ := mapPseudoRegion(b.Region)
	if len(region) == 0 {
		return fmt.Errorf("Could not resolve AWS::Region")
	} else {
		params.Region = aws.String(region)
	}
	if b.UseDualStack == aws.DualStackEndpointStateEnabled {
		params.UseDualStack = aws.Bool(true)
	} else {
		params.UseDualStack = aws.Bool(false)
	}
	if b.UseFIPS == aws.FIPSEndpointStateEnabled {
		params.UseFIPS = aws.Bool(true)
	} else {
		params.UseFIPS = aws.Bool(false)
	}
	params.Endpoint = b.Endpoint
	return nil
}

// EndpointParameters provides the parameters that influence how endpoints are
// resolved.
type EndpointParameters struct {
	// The AWS region used to dispatch the request.
	//
	// Parameter is
	// required.
	//
	// AWS::Region
	Region *string

	// When true, use the dual-stack endpoint. If the configured endpoint does not
	// support dual-stack, dispatching the request MAY return an error.
	//
	// Defaults to
	// false if no value is provided.
	//
	// AWS::UseDualStack
	UseDualStack *bool

	// When true, send this request to the FIPS-compliant regional endpoint. If the
	// configured endpoint does not have a FIPS compliant endpoint, dispatching the
	// request will return an error.
	//
	// Defaults to false if no value is
	// provided.
	//
	// AWS::UseFIPS
	UseFIPS *bool

	// Override the endpoint used to send this request
	//
	// Parameter is
	// required.
	//
	// SDK::Endpoint
	Endpoint *string

	// Operation parameter for EndpointId
	//
	// Parameter is required.
	EndpointId *string
}

// ValidateRequired validates required parameters are set.
func (p EndpointParameters) ValidateRequired() error {
	if p.UseDualStack == nil {
		return fmt.Errorf("parameter UseDualStack is required")
	}

	if p.UseFIPS == nil {
		return fmt.Errorf("parameter UseFIPS is required")
	}

	return nil
}

// WithDefaults returns a shallow copy of EndpointParameterswith default values
// applied to members where applicable.
func (p EndpointParameters) WithDefaults() EndpointParameters {
	if p.UseDualStack == nil {
		p.UseDualStack = ptr.Bool(false)
	}

	if p.UseFIPS == nil {
		p.UseFIPS = ptr.Bool(false)
	}
	return p
}

// EndpointResolverV2 provides the interface for resolving service endpoints.
type EndpointResolverV2 interface {
	// ResolveEndpoint attempts to resolve the endpoint with the provided options,
	// returning the endpoint if found. Otherwise an error is returned.
	ResolveEndpoint(ctx context.Context, params EndpointParameters) (
		smithyendpoints.Endpoint, error,
	)
}

// resolver provides the implementation for resolving endpoints.
type resolver struct{}

func NewDefaultEndpointResolverV2() EndpointResolverV2 {
	return &resolver{}
}

// ResolveEndpoint attempts to resolve the endpoint with the provided options,
// returning the endpoint if found. Otherwise an error is returned.
func (r *resolver) ResolveEndpoint(
	ctx context.Context, params EndpointParameters,
) (
	endpoint smithyendpoints.Endpoint, err error,
) {
	params = params.WithDefaults()
	if err = params.ValidateRequired(); err != nil {
		return endpoint, fmt.Errorf("endpoint parameters are not valid, %w", err)
	}
	_UseDualStack := *params.UseDualStack
	_UseFIPS := *params.UseFIPS

	if exprVal := params.EndpointId; exprVal != nil {
		_EndpointId := *exprVal
		_ = _EndpointId
		if exprVal := params.Region; exprVal != nil {
			_Region := *exprVal
			_ = _Region
			if exprVal := awsrulesfn.GetPartition(_Region); exprVal != nil {
				_PartitionResult := *exprVal
				_ = _PartitionResult
				if rulesfn.IsValidHostLabel(_EndpointId, true) {
					if _UseFIPS == false {
						if exprVal := params.Endpoint; exprVal != nil {
							_Endpoint := *exprVal
							_ = _Endpoint
							uriString := _Endpoint

							uri, err := url.Parse(uriString)
							if err != nil {
								return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
							}

							return smithyendpoints.Endpoint{
								URI:     *uri,
								Headers: http.Header{},
								Properties: func() smithy.Properties {
									var out smithy.Properties
									out.Set("authSchemes", []interface{}{
										map[string]interface{}{
											"name":        "sigv4a",
											"signingName": "events",
											"signingRegionSet": []interface{}{
												"*",
											},
										},
									})
									return out
								}(),
							}, nil
						}
						if _UseDualStack == true {
							if true == _PartitionResult.SupportsDualStack {
								uriString := func() string {
									var out strings.Builder
									out.WriteString("https://")
									out.WriteString(_EndpointId)
									out.WriteString(".endpoint.events.")
									out.WriteString(_PartitionResult.DualStackDnsSuffix)
									return out.String()
								}()

								uri, err := url.Parse(uriString)
								if err != nil {
									return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
								}

								return smithyendpoints.Endpoint{
									URI:     *uri,
									Headers: http.Header{},
									Properties: func() smithy.Properties {
										var out smithy.Properties
										out.Set("authSchemes", []interface{}{
											map[string]interface{}{
												"name":        "sigv4a",
												"signingName": "events",
												"signingRegionSet": []interface{}{
													"*",
												},
											},
										})
										return out
									}(),
								}, nil
							}
							return endpoint, fmt.Errorf("endpoint rule error, %s", "DualStack is enabled but this partition does not support DualStack")
						}
						uriString := func() string {
							var out strings.Builder
							out.WriteString("https://")
							out.WriteString(_EndpointId)
							out.WriteString(".endpoint.events.")
							out.WriteString(_PartitionResult.DnsSuffix)
							return out.String()
						}()

						uri, err := url.Parse(uriString)
						if err != nil {
							return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
						}

						return smithyendpoints.Endpoint{
							URI:     *uri,
							Headers: http.Header{},
							Properties: func() smithy.Properties {
								var out smithy.Properties
								out.Set("authSchemes", []interface{}{
									map[string]interface{}{
										"name":        "sigv4a",
										"signingName": "events",
										"signingRegionSet": []interface{}{
											"*",
										},
									},
								})
								return out
							}(),
						}, nil
					}
					return endpoint, fmt.Errorf("endpoint rule error, %s", "Invalid Configuration: FIPS is not supported with EventBridge multi-region endpoints.")
				}
				return endpoint, fmt.Errorf("endpoint rule error, %s", "EndpointId must be a valid host label.")
			}
		}
	}
	if exprVal := params.Endpoint; exprVal != nil {
		_Endpoint := *exprVal
		_ = _Endpoint
		if _UseFIPS == true {
			return endpoint, fmt.Errorf("endpoint rule error, %s", "Invalid Configuration: FIPS and custom endpoint are not supported")
		}
		if _UseDualStack == true {
			return endpoint, fmt.Errorf("endpoint rule error, %s", "Invalid Configuration: Dualstack and custom endpoint are not supported")
		}
		uriString := _Endpoint

		uri, err := url.Parse(uriString)
		if err != nil {
			return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
		}

		return smithyendpoints.Endpoint{
			URI:     *uri,
			Headers: http.Header{},
		}, nil
	}
	if exprVal := params.Region; exprVal != nil {
		_Region := *exprVal
		_ = _Region
		if exprVal := awsrulesfn.GetPartition(_Region); exprVal != nil {
			_PartitionResult := *exprVal
			_ = _PartitionResult
			if _UseFIPS == true {
				if _UseDualStack == true {
					if true == _PartitionResult.SupportsFIPS {
						if true == _PartitionResult.SupportsDualStack {
							uriString := func() string {
								var out strings.Builder
								out.WriteString("https://events-fips.")
								out.WriteString(_Region)
								out.WriteString(".")
								out.WriteString(_PartitionResult.DualStackDnsSuffix)
								return out.String()
							}()

							uri, err := url.Parse(uriString)
							if err != nil {
								return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
							}

							return smithyendpoints.Endpoint{
								URI:     *uri,
								Headers: http.Header{},
							}, nil
						}
					}
					return endpoint, fmt.Errorf("endpoint rule error, %s", "FIPS and DualStack are enabled, but this partition does not support one or both")
				}
			}
			if _UseFIPS == true {
				if true == _PartitionResult.SupportsFIPS {
					if _Region == "us-gov-east-1" {
						uriString := "https://events.us-gov-east-1.amazonaws.com"

						uri, err := url.Parse(uriString)
						if err != nil {
							return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
						}

						return smithyendpoints.Endpoint{
							URI:     *uri,
							Headers: http.Header{},
						}, nil
					}
					if _Region == "us-gov-west-1" {
						uriString := "https://events.us-gov-west-1.amazonaws.com"

						uri, err := url.Parse(uriString)
						if err != nil {
							return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
						}

						return smithyendpoints.Endpoint{
							URI:     *uri,
							Headers: http.Header{},
						}, nil
					}
					uriString := func() string {
						var out strings.Builder
						out.WriteString("https://events-fips.")
						out.WriteString(_Region)
						out.WriteString(".")
						out.WriteString(_PartitionResult.DnsSuffix)
						return out.String()
					}()

					uri, err := url.Parse(uriString)
					if err != nil {
						return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
					}

					return smithyendpoints.Endpoint{
						URI:     *uri,
						Headers: http.Header{},
					}, nil
				}
				return endpoint, fmt.Errorf("endpoint rule error, %s", "FIPS is enabled but this partition does not support FIPS")
			}
			if _UseDualStack == true {
				if true == _PartitionResult.SupportsDualStack {
					uriString := func() string {
						var out strings.Builder
						out.WriteString("https://events.")
						out.WriteString(_Region)
						out.WriteString(".")
						out.WriteString(_PartitionResult.DualStackDnsSuffix)
						return out.String()
					}()

					uri, err := url.Parse(uriString)
					if err != nil {
						return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
					}

					return smithyendpoints.Endpoint{
						URI:     *uri,
						Headers: http.Header{},
					}, nil
				}
				return endpoint, fmt.Errorf("endpoint rule error, %s", "DualStack is enabled but this partition does not support DualStack")
			}
			uriString := func() string {
				var out strings.Builder
				out.WriteString("https://events.")
				out.WriteString(_Region)
				out.WriteString(".")
				out.WriteString(_PartitionResult.DnsSuffix)
				return out.String()
			}()

			uri, err := url.Parse(uriString)
			if err != nil {
				return endpoint, fmt.Errorf("Failed to parse uri: %s", uriString)
			}

			return smithyendpoints.Endpoint{
				URI:     *uri,
				Headers: http.Header{},
			}, nil
		}
		return endpoint, fmt.Errorf("no rules matched these parameters. This is a bug, %#v", params)
	}
	return endpoint, fmt.Errorf("endpoint rule error, %s", "Invalid Configuration: Missing Region")
}
