package auth

import (
	"context"
	"fmt"

	"github.com/isutare412/bloated/api/pkg/core/enum"
	"github.com/isutare412/bloated/api/pkg/core/model"
	"github.com/isutare412/bloated/api/pkg/core/port"
	"github.com/isutare412/bloated/api/pkg/pkgerror"
)

type Service struct {
	customJWTClient port.CustomJWTClient
	googleJWTClient port.GoogleJWTClient
	userRepo        port.UserRepository
}

func NewService(
	customJWTClient port.CustomJWTClient,
	googleJWTClient port.GoogleJWTClient,
	userRepo port.UserRepository,
) *Service {
	return &Service{
		customJWTClient: customJWTClient,
		googleJWTClient: googleJWTClient,
		userRepo:        userRepo,
	}
}

func (s *Service) IssueCustomToken(ctx context.Context, token model.CustomToken) (string, error) {
	if err := token.Validate(); err != nil {
		return "", fmt.Errorf("validating custom token: %w", err)
	}

	user, err := s.userRepo.Upsert(ctx, token.ToUser(enum.IssuerNone))
	if err != nil {
		return "", fmt.Errorf("upserting user: %w", err)
	}

	token.UserID = user.ID.String()
	signedString, err := s.customJWTClient.SignCustomToken(token)
	if err != nil {
		return "", fmt.Errorf("signing custom token: %w", err)
	}

	return signedString, nil
}

func (s *Service) IssueCustomTokenFromGoogle(ctx context.Context, tokenString string) (string, error) {
	googleToken, err := s.googleJWTClient.VerifyGoogleIDToken(tokenString)
	if err != nil {
		return "", pkgerror.Known{
			Code:   pkgerror.CodeBadRequest,
			Origin: fmt.Errorf("verifying google ID token: %w", err),
			Simple: fmt.Errorf("failed to verify Google ID token"),
		}
	}
	if err := googleToken.Validate(); err != nil {
		return "", fmt.Errorf("validating google ID token: %w", err)
	}

	customToken := googleToken.ToCustomToken()
	if err := customToken.Validate(); err != nil {
		return "", fmt.Errorf("validating custom token: %w", err)
	}

	user, err := s.userRepo.Upsert(ctx, customToken.ToUser(enum.IssuerGoogle))
	if err != nil {
		return "", fmt.Errorf("upserting user: %w", err)
	}

	customToken.UserID = user.ID.String()
	signedString, err := s.customJWTClient.SignCustomToken(customToken)
	if err != nil {
		return "", fmt.Errorf("signing custom token: %w", err)
	}

	return signedString, nil
}

func (s *Service) VerifyCustomToken(ctx context.Context, tokenString string) (model.CustomToken, error) {
	token, err := s.customJWTClient.VerifyCustomToken(tokenString)
	if err != nil {
		return model.CustomToken{}, pkgerror.Known{
			Code:   pkgerror.CodeBadRequest,
			Origin: fmt.Errorf("verifying custom token: %w", err),
			Simple: fmt.Errorf("failed to verify token"),
		}
	}

	return token, nil
}
