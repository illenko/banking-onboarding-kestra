package http

import (
	"github.com/google/uuid"
)

type OnboardingRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	City        string `json:"city"`
	AccountType string `json:"account_type"`
	Currency    string `json:"currency"`
}

type OnboardingStatus struct {
	ID    uuid.UUID       `json:"id"`
	State OnboardingState `json:"state"`
	Data  map[string]any  `json:"data,omitempty"`
}

type SignatureRequest struct {
	Signature string `json:"signature"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type OnboardingState string

const (
	ProcessingState         OnboardingState = "processing"
	FailedState             OnboardingState = "failed"
	FraudNotPassedState     OnboardingState = "fraud_not_passed"
	SignatureNotValidSate   OnboardingState = "signature_not_valid"
	WaitingForAgreementSign OnboardingState = "waiting_for_agreement_signature"
	CompletedState          OnboardingState = "completed"
)
