package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
	"github.com/tribefintech/microservices/internal/common/cmerr"
)

const (
	ErrInvalidPassword       = "invalid-password"
	ErrUserNotConfirmed      = "user-not-confirmed"
	ErrUsernameExisted       = "username-existed"
	ErrInvalidParameter      = "invalid-parameter"
	ErrCodeDeliveryFailure   = "code-delivery-failure"
	ErrUserConfirmed         = "user-confirmed"
	ErrCodeMismatch          = "code-mismatch"
	ErrCodeExpired           = "code-expired"
	ErrTooManyFailedAttempts = "too-many-failed-attempts"
	ErrUserNotFound          = "user-not-found"
	ErrRateLimited           = "rate-limited"
	ErrInvalidRefreshToken   = "invalid-refresh-token"
)

type cognito struct {
	provider cognitoidentityprovideriface.CognitoIdentityProviderAPI
	clientId string
}

func newCognito(clientId string, provider cognitoidentityprovideriface.CognitoIdentityProviderAPI) *cognito {
	return &cognito{
		provider: provider,
		clientId: clientId,
	}
}

func (c *cognito) SignUp(username, password, firstName, lastName string) (*SignUpResponse, error) {
	userInput := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(c.clientId),
		Password: aws.String(password),
		Username: aws.String(username),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("given_name"),
				Value: aws.String(firstName),
			},
			{
				Name:  aws.String("family_name"),
				Value: aws.String(lastName),
			},
		},
		ValidationData: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(username),
			},
		},
	}

	res, err := c.provider.SignUp(userInput)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case cognitoidentityprovider.ErrCodeUsernameExistsException:
				return nil, cmerr.New(
					fmt.Errorf("cognito.SignUp: %w", awsErr),
					ErrUsernameExisted,
					cmerr.TypDomainError,
				)
			case cognitoidentityprovider.ErrCodeInvalidParameterException:
				return nil, cmerr.New(
					fmt.Errorf("cognito.SignUp: %w", awsErr),
					ErrInvalidPassword,
					cmerr.TypDomainError,
				)

			default:
				return nil, cmerr.New(
					fmt.Errorf("cognito.SignUp: %w", awsErr),
					cmerr.InternalServerError,
					cmerr.TypUnexpected,
				)
			}
		}

		return nil, cmerr.New(
			fmt.Errorf("cognito.Login: %w", err),
			cmerr.InternalServerError,
			cmerr.TypUnexpected,
		)
	}

	return &SignUpResponse{
		UserSub:     *res.UserSub,
		Destination: *res.CodeDeliveryDetails.Destination,
	}, nil
}

func (c *cognito) ConfirmSignUp(username, code string) error {

	confirmInput := &cognitoidentityprovider.ConfirmSignUpInput{
		ConfirmationCode: aws.String(code),
		ClientId:         aws.String(c.clientId),
		Username:         aws.String(username),
	}

	_, err := c.provider.ConfirmSignUp(confirmInput)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case cognitoidentityprovider.ErrCodeNotAuthorizedException:
				return cmerr.New(
					fmt.Errorf("cognito.ConfirmSignUp: %w", awsErr),
					ErrUserConfirmed,
					cmerr.TypDomainError,
				)
			case cognitoidentityprovider.ErrCodeCodeMismatchException:
				return cmerr.New(
					fmt.Errorf("cognito.ConfirmSignUp: %w", awsErr),
					ErrCodeMismatch,
					cmerr.TypDomainError,
				)
			case cognitoidentityprovider.ErrCodeExpiredCodeException:
				return cmerr.New(
					fmt.Errorf("cognito.ConfirmSignUp: %w", awsErr),
					ErrCodeExpired,
					cmerr.TypDomainError,
				)
			case cognitoidentityprovider.ErrCodeUserNotFoundException:
				return cmerr.New(
					fmt.Errorf("cognito.ConfirmSignUp: %w", awsErr),
					ErrUserNotFound,
					cmerr.TypDomainError,
				)

			default:
				return cmerr.New(
					fmt.Errorf("cognito.ConfirmSignUp: %w", awsErr),
					cmerr.InternalServerError,
					cmerr.TypUnexpected,
				)
			}
		}

		return cmerr.New(
			fmt.Errorf("cognito.ConfirmSignUp: %w", err),
			cmerr.InternalServerError,
			cmerr.TypUnexpected,
		)
	}

	return nil
}

func (c *cognito) ResendConfirmationCode(username string) (*ResendCodeResponse, error) {
	input := &cognitoidentityprovider.ResendConfirmationCodeInput{
		ClientId: aws.String(c.clientId),
		Username: aws.String(username),
	}

	res, err := c.provider.ResendConfirmationCode(input)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case cognitoidentityprovider.ErrCodeUserNotFoundException:
				return nil, cmerr.New(
					fmt.Errorf("cognito.ResendConfirmationCode: %w", awsErr),
					ErrUserNotFound,
					cmerr.TypDomainError,
				)

			default:
				return nil, cmerr.New(
					fmt.Errorf("cognito.ResendConfirmationCode: %w", awsErr),
					cmerr.InternalServerError,
					cmerr.TypUnexpected,
				)
			}
		}

		return nil, cmerr.New(
			fmt.Errorf("cognito.ResendConfirmationCode: %w", err),
			cmerr.InternalServerError,
			cmerr.TypUnexpected,
		)
	}

	return &ResendCodeResponse{
		DeliveryMedium: res.CodeDeliveryDetails.DeliveryMedium,
		Destination:    res.CodeDeliveryDetails.Destination,
	}, nil
}

func (c *cognito) Login(username, password string) (*LoginResponse, error) {
	authInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String(cognitoidentityprovider.AuthFlowTypeUserPasswordAuth),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(username),
			"PASSWORD": aws.String(password),
		},
		ClientId: aws.String(c.clientId),
	}

	// Call the InitiateAuth method to initiate the authentication process
	res, err := c.provider.InitiateAuth(authInput)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case cognitoidentityprovider.ErrCodeUserNotFoundException:
				return nil, cmerr.New(
					fmt.Errorf("cognito.Login: %w", awsErr),
					ErrUserNotFound,
					cmerr.TypDomainError,
				)

			case cognitoidentityprovider.ErrCodeNotAuthorizedException:
				return nil, cmerr.New(
					fmt.Errorf("cognito.Login: %w", awsErr),
					ErrInvalidPassword,
					cmerr.TypDomainError,
				)

			case cognitoidentityprovider.ErrCodeUserNotConfirmedException:
				return nil, cmerr.New(
					fmt.Errorf("cognito.Login: %w", awsErr),
					ErrUserNotConfirmed,
					cmerr.TypDomainError,
				)

			default:
				return nil, cmerr.New(
					fmt.Errorf("cognito.Login: %w", awsErr),
					cmerr.InternalServerError,
					cmerr.TypUnexpected,
				)
			}
		}

		return nil, cmerr.New(
			fmt.Errorf("cognito.Login: %w", err),
			cmerr.InternalServerError,
			cmerr.TypUnexpected,
		)
	}

	return &LoginResponse{
		AccessToken:  *res.AuthenticationResult.AccessToken,
		IdToken:      *res.AuthenticationResult.IdToken,
		ExpireIn:     *res.AuthenticationResult.ExpiresIn,
		TokenType:    *res.AuthenticationResult.TokenType,
		RefreshToken: *res.AuthenticationResult.RefreshToken,
	}, nil

}

func (c *cognito) RefreshToken(refreshToken string) (*RefreshTokenResponse, error) {
	authInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String(cognitoidentityprovider.AuthFlowTypeRefreshTokenAuth),
		AuthParameters: map[string]*string{
			"REFRESH_TOKEN": aws.String(refreshToken),
		},
		ClientId: aws.String(c.clientId),
	}

	// Call the InitiateAuth method to initiate the authentication process
	res, err := c.provider.InitiateAuth(authInput)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case cognitoidentityprovider.ErrCodeNotAuthorizedException:
				return nil, cmerr.New(
					fmt.Errorf("cognito.RefreshToken: %w", awsErr),
					ErrInvalidRefreshToken,
					cmerr.TypIncorrectInput,
				)

			default:
				return nil, cmerr.New(
					fmt.Errorf("cognito.RefreshToken: %w", awsErr),
					cmerr.InternalServerError,
					cmerr.TypUnexpected,
				)
			}
		}

		return nil, cmerr.New(
			fmt.Errorf("cognito.RefreshToken: %w", err),
			cmerr.InternalServerError,
			cmerr.TypUnexpected,
		)
	}

	return &RefreshTokenResponse{
		AccessToken: *res.AuthenticationResult.AccessToken,
		IdToken:     *res.AuthenticationResult.IdToken,
		ExpireIn:    *res.AuthenticationResult.ExpiresIn,
		TokenType:   *res.AuthenticationResult.TokenType,
	}, nil
}

func (c *cognito) ForgotPassword(username string) (*ForgotPasswordResponse, error) {
	res, err := c.provider.ForgotPassword(&cognitoidentityprovider.ForgotPasswordInput{
		Username: &username,
		ClientId: aws.String(c.clientId),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case cognitoidentityprovider.ErrCodeUserNotFoundException:
				return nil, cmerr.New(
					fmt.Errorf("cognito.ForgotPassword: %w", awsErr),
					ErrUserNotFound,
					cmerr.TypDomainError,
				)

			default:
				return nil, cmerr.New(
					fmt.Errorf("cognito.Login: %w", awsErr),
					cmerr.InternalServerError,
					cmerr.TypUnexpected,
				)
			}
		}

		return nil, cmerr.New(
			fmt.Errorf("cognito.Login: %w", err),
			cmerr.InternalServerError,
			cmerr.TypUnexpected,
		)
	}

	return &ForgotPasswordResponse{
		DeliveryMedium: res.CodeDeliveryDetails.DeliveryMedium,
		Destination:    res.CodeDeliveryDetails.Destination,
	}, nil
}

func (c *cognito) ConfirmForgotPassword(username, code, password string) error {

	confirmInput := &cognitoidentityprovider.ConfirmForgotPasswordInput{
		ConfirmationCode: aws.String(code),
		ClientId:         aws.String(c.clientId),
		Username:         aws.String(username),
		Password:         aws.String(password),
	}

	_, err := c.provider.ConfirmForgotPassword(confirmInput)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case cognitoidentityprovider.ErrCodeNotAuthorizedException:
				return cmerr.New(
					fmt.Errorf("cognito.ConfirmForgotPassword: %w", awsErr),
					ErrUserConfirmed,
					cmerr.TypDomainError,
				)
			case cognitoidentityprovider.ErrCodeCodeMismatchException:
				return cmerr.New(
					fmt.Errorf("cognito.ConfirmForgotPassword: %w", awsErr),
					ErrCodeMismatch,
					cmerr.TypDomainError,
				)
			case cognitoidentityprovider.ErrCodeExpiredCodeException:
				return cmerr.New(
					fmt.Errorf("cognito.ConfirmForgotPassword: %w", awsErr),
					ErrCodeExpired,
					cmerr.TypDomainError,
				)

			case cognitoidentityprovider.ErrCodeInvalidParameterException:
				return cmerr.New(
					fmt.Errorf("cognito.ConfirmForgotPassword: %w", awsErr),
					ErrInvalidParameter,
					cmerr.TypDomainError,
				)
			case cognitoidentityprovider.ErrCodeUserNotFoundException:
				return cmerr.New(
					fmt.Errorf("cognito.ConfirmForgotPassword: %w", awsErr),
					ErrUserNotFound,
					cmerr.TypDomainError,
				)

			default:
				return cmerr.New(
					fmt.Errorf("cognito.ConfirmForgotPassword: %w", awsErr),
					cmerr.InternalServerError,
					cmerr.TypUnexpected,
				)
			}
		}

		return cmerr.New(
			fmt.Errorf("cognito.ConfirmForgotPassword: %w", err),
			cmerr.InternalServerError,
			cmerr.TypUnexpected,
		)
	}

	return nil
}

func (c *cognito) ChangePassword(token, newp, oldp string) error {
	_, err := c.provider.ChangePassword(&cognitoidentityprovider.ChangePasswordInput{
		AccessToken:      &token,
		PreviousPassword: &oldp,
		ProposedPassword: &newp,
	})

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case cognitoidentityprovider.ErrCodeNotAuthorizedException:
				return cmerr.New(
					fmt.Errorf("cognito.SignUp: %w", awsErr),
					ErrInvalidPassword,
					cmerr.TypDomainError,
				)

			default:
				return cmerr.New(
					fmt.Errorf("cognito.ConfirmForgotPassword: %w", awsErr),
					cmerr.InternalServerError,
					cmerr.TypUnexpected,
				)
			}
		}

		return cmerr.New(
			fmt.Errorf("cognito.ConfirmForgotPassword: %w", err),
			cmerr.InternalServerError,
			cmerr.TypUnexpected,
		)
	}

	return nil
}
