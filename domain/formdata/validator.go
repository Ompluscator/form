package formdata

import (
	"context"

	"flamingo.me/flamingo/v3/framework/web"
	"flamingo.me/form/domain"
)

type (
	// DefaultFormDataValidatorImpl represents implementation of default domain.FormDataValidator.
	DefaultFormDataValidatorImpl struct{}
)

var _ domain.DefaultFormDataValidator = &DefaultFormDataValidatorImpl{}

// Validate performs default form data validation, by using go-playground validator package and storing results into domain.ValidationInfo instance.
func (p *DefaultFormDataValidatorImpl) Validate(ctx context.Context, req *web.Request, validatorProvider domain.ValidatorProvider, formData interface{}) (*domain.ValidationInfo, error) {
	if _, ok := formData.(map[string]string); ok {
		return &domain.ValidationInfo{}, nil
	}
	validationInfo := validatorProvider.Validate(ctx, req, formData)
	return &validationInfo, nil
}
