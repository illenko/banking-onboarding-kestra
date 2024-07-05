package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/illenko/onboarding-service/pkg/http"
)

type OnboardingService struct {
}

func NewOnboardingService() *OnboardingService {
	return &OnboardingService{}
}

func (s *OnboardingService) CreateOnboarding(ctx context.Context, req *http.OnboardingRequest) (http.OnboardingStatus, error) {
	workflowId := uuid.New()

	return http.OnboardingStatus{
		ID:    workflowId,
		State: http.ProcessingState,
	}, nil
}

func (s *OnboardingService) GetOnboarding(ctx context.Context, id uuid.UUID) (http.OnboardingStatus, error) {

	return http.OnboardingStatus{
		ID:    id,
		State: http.ProcessingState,
		Data:  map[string]any{},
	}, nil
}

func (s *OnboardingService) SignAgreement(ctx context.Context, id uuid.UUID, req *http.SignatureRequest) (http.OnboardingStatus, error) {

	return http.OnboardingStatus{
		ID:    id,
		State: http.ProcessingState,
	}, nil
}
