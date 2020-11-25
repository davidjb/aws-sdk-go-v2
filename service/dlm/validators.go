// Code generated by smithy-go-codegen DO NOT EDIT.

package dlm

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dlm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpCreateLifecyclePolicy struct {
}

func (*validateOpCreateLifecyclePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateLifecyclePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateLifecyclePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateLifecyclePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteLifecyclePolicy struct {
}

func (*validateOpDeleteLifecyclePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteLifecyclePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteLifecyclePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteLifecyclePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetLifecyclePolicy struct {
}

func (*validateOpGetLifecyclePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetLifecyclePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetLifecyclePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetLifecyclePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateLifecyclePolicy struct {
}

func (*validateOpUpdateLifecyclePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateLifecyclePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateLifecyclePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateLifecyclePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateLifecyclePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateLifecyclePolicy{}, middleware.After)
}

func addOpDeleteLifecyclePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteLifecyclePolicy{}, middleware.After)
}

func addOpGetLifecyclePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetLifecyclePolicy{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateLifecyclePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateLifecyclePolicy{}, middleware.After)
}

func validateCrossRegionCopyRule(v *types.CrossRegionCopyRule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CrossRegionCopyRule"}
	if v.TargetRegion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetRegion"))
	}
	if v.Encrypted == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Encrypted"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCrossRegionCopyRules(v []types.CrossRegionCopyRule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CrossRegionCopyRules"}
	for i := range v {
		if err := validateCrossRegionCopyRule(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFastRestoreRule(v *types.FastRestoreRule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "FastRestoreRule"}
	if v.AvailabilityZones == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AvailabilityZones"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePolicyDetails(v *types.PolicyDetails) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PolicyDetails"}
	if v.Schedules != nil {
		if err := validateScheduleList(v.Schedules); err != nil {
			invalidParams.AddNested("Schedules", err.(smithy.InvalidParamsError))
		}
	}
	if v.TargetTags != nil {
		if err := validateTargetTagList(v.TargetTags); err != nil {
			invalidParams.AddNested("TargetTags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSchedule(v *types.Schedule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Schedule"}
	if v.CrossRegionCopyRules != nil {
		if err := validateCrossRegionCopyRules(v.CrossRegionCopyRules); err != nil {
			invalidParams.AddNested("CrossRegionCopyRules", err.(smithy.InvalidParamsError))
		}
	}
	if v.TagsToAdd != nil {
		if err := validateTagsToAddList(v.TagsToAdd); err != nil {
			invalidParams.AddNested("TagsToAdd", err.(smithy.InvalidParamsError))
		}
	}
	if v.VariableTags != nil {
		if err := validateVariableTagsList(v.VariableTags); err != nil {
			invalidParams.AddNested("VariableTags", err.(smithy.InvalidParamsError))
		}
	}
	if v.FastRestoreRule != nil {
		if err := validateFastRestoreRule(v.FastRestoreRule); err != nil {
			invalidParams.AddNested("FastRestoreRule", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateScheduleList(v []types.Schedule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ScheduleList"}
	for i := range v {
		if err := validateSchedule(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagsToAddList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagsToAddList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTargetTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TargetTagList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateVariableTagsList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VariableTagsList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateLifecyclePolicyInput(v *CreateLifecyclePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateLifecyclePolicyInput"}
	if len(v.State) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("State"))
	}
	if v.Description == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Description"))
	}
	if v.PolicyDetails == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyDetails"))
	} else if v.PolicyDetails != nil {
		if err := validatePolicyDetails(v.PolicyDetails); err != nil {
			invalidParams.AddNested("PolicyDetails", err.(smithy.InvalidParamsError))
		}
	}
	if v.ExecutionRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExecutionRoleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteLifecyclePolicyInput(v *DeleteLifecyclePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteLifecyclePolicyInput"}
	if v.PolicyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetLifecyclePolicyInput(v *GetLifecyclePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetLifecyclePolicyInput"}
	if v.PolicyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateLifecyclePolicyInput(v *UpdateLifecyclePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateLifecyclePolicyInput"}
	if v.PolicyDetails != nil {
		if err := validatePolicyDetails(v.PolicyDetails); err != nil {
			invalidParams.AddNested("PolicyDetails", err.(smithy.InvalidParamsError))
		}
	}
	if v.PolicyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PolicyId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
