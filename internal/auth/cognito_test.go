package main

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/golang/mock/gomock"

	"github.com/tribefintech/microservices/internal/auth/mock/mock_cognitoidentityprovideriface"
)

const (
	clientId = "387bpblqtgrtucq2vrl65kaq16"
)

func Test_cognito_SignUp(t *testing.T) {
	type args struct {
		username  string
		password  string
		firstName string
		lastName  string
	}
	tests := []struct {
		name    string
		args    args
		want    *SignUpResponse
		wantErr bool
	}{
		{
			name: "should create an account",
			args: args{
				username:  "username",
				password:  "password",
				firstName: "firstName",
				lastName:  "lastName",
			},
			want: &SignUpResponse{
				Destination: "dk@vinova.com.sg",
				UserSub:     "135fb051-2830-4285-a280-6c862cc2ef03",
			},
			wantErr: false,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mock_cognitoidentityprovideriface.NewMockCognitoIdentityProviderAPI(ctrl)

	m.EXPECT().SignUp(gomock.Eq(&cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(clientId),
		Password: aws.String("password"),
		Username: aws.String("username"),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("given_name"),
				Value: aws.String("firstName"),
			},
			{
				Name:  aws.String("family_name"),
				Value: aws.String("lastName"),
			},
		},
		ValidationData: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String("username"),
			},
		},
	})).Return(&cognitoidentityprovider.SignUpOutput{
		CodeDeliveryDetails: &cognitoidentityprovider.CodeDeliveryDetailsType{
			Destination: aws.String("dk@vinova.com.sg"),
		},
		UserSub: aws.String("135fb051-2830-4285-a280-6c862cc2ef03"),
	}, nil)

	c := newCognito(clientId, m)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := c.SignUp(tt.args.username, tt.args.password, tt.args.firstName, tt.args.lastName)
			if (err != nil) != tt.wantErr {
				t.Errorf("cognito.SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cognito.SignUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
