// Code generated by smithy-go-codegen DO NOT EDIT.

package macie

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/macie/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpAssociateMemberAccount struct {
}

func (*validateOpAssociateMemberAccount) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateMemberAccount) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateMemberAccountInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateMemberAccountInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAssociateS3Resources struct {
}

func (*validateOpAssociateS3Resources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateS3Resources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateS3ResourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateS3ResourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateMemberAccount struct {
}

func (*validateOpDisassociateMemberAccount) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateMemberAccount) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateMemberAccountInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateMemberAccountInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateS3Resources struct {
}

func (*validateOpDisassociateS3Resources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateS3Resources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateS3ResourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateS3ResourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateS3Resources struct {
}

func (*validateOpUpdateS3Resources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateS3Resources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateS3ResourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateS3ResourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAssociateMemberAccountValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateMemberAccount{}, middleware.After)
}

func addOpAssociateS3ResourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateS3Resources{}, middleware.After)
}

func addOpDisassociateMemberAccountValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateMemberAccount{}, middleware.After)
}

func addOpDisassociateS3ResourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateS3Resources{}, middleware.After)
}

func addOpUpdateS3ResourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateS3Resources{}, middleware.After)
}

func validateClassificationType(v *types.ClassificationType) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ClassificationType"}
	if len(v.OneTime) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("OneTime"))
	}
	if len(v.Continuous) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Continuous"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3Resource(v *types.S3Resource) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3Resource"}
	if v.BucketName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BucketName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3ResourceClassification(v *types.S3ResourceClassification) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3ResourceClassification"}
	if v.BucketName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BucketName"))
	}
	if v.ClassificationType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClassificationType"))
	} else if v.ClassificationType != nil {
		if err := validateClassificationType(v.ClassificationType); err != nil {
			invalidParams.AddNested("ClassificationType", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3ResourceClassificationUpdate(v *types.S3ResourceClassificationUpdate) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3ResourceClassificationUpdate"}
	if v.BucketName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BucketName"))
	}
	if v.ClassificationTypeUpdate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClassificationTypeUpdate"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3Resources(v []types.S3Resource) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3Resources"}
	for i := range v {
		if err := validateS3Resource(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3ResourcesClassification(v []types.S3ResourceClassification) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3ResourcesClassification"}
	for i := range v {
		if err := validateS3ResourceClassification(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3ResourcesClassificationUpdate(v []types.S3ResourceClassificationUpdate) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3ResourcesClassificationUpdate"}
	for i := range v {
		if err := validateS3ResourceClassificationUpdate(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateMemberAccountInput(v *AssociateMemberAccountInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateMemberAccountInput"}
	if v.MemberAccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberAccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateS3ResourcesInput(v *AssociateS3ResourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateS3ResourcesInput"}
	if v.S3Resources == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Resources"))
	} else if v.S3Resources != nil {
		if err := validateS3ResourcesClassification(v.S3Resources); err != nil {
			invalidParams.AddNested("S3Resources", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateMemberAccountInput(v *DisassociateMemberAccountInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateMemberAccountInput"}
	if v.MemberAccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberAccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateS3ResourcesInput(v *DisassociateS3ResourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateS3ResourcesInput"}
	if v.AssociatedS3Resources == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AssociatedS3Resources"))
	} else if v.AssociatedS3Resources != nil {
		if err := validateS3Resources(v.AssociatedS3Resources); err != nil {
			invalidParams.AddNested("AssociatedS3Resources", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateS3ResourcesInput(v *UpdateS3ResourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateS3ResourcesInput"}
	if v.S3ResourcesUpdate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3ResourcesUpdate"))
	} else if v.S3ResourcesUpdate != nil {
		if err := validateS3ResourcesClassificationUpdate(v.S3ResourcesUpdate); err != nil {
			invalidParams.AddNested("S3ResourcesUpdate", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}