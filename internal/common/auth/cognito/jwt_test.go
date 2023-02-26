package cognito

import (
	"testing"
)

const (
	AWS_COGNITO_REGION       = "ap-southeast-1"
	AWS_COGNITO_USER_POOL_ID = "ap-southeast-1_Sly3MWVC8"
)

func Test_getJWT(t *testing.T) {

	auth := New(AWS_COGNITO_REGION, AWS_COGNITO_USER_POOL_ID)

	err := auth.getJWK()
	if err != nil {
		t.Error(err)
	}

	jwt := "eyJraWQiOiJVOEdYNk82Y21ZcEV1Q2FCRGdJM3VISlFUS1wvY2VVZzJGQ2hFQjlXRDRNZz0iLCJhbGciOiJSUzI1NiJ9.eyJvcmlnaW5fanRpIjoiYjU5Y2YxNWEtY2VkOS00OTk4LWE4YzAtMGU5MWMyYTA3NzQwIiwic3ViIjoiZmU0MzU5MTgtMDkyYS00YmE4LWI5OTctNGMwYmY2Nzg1OWQwIiwiYXVkIjoiN2o4a3UyNjBiaWJvNGR0ZTJoOHZvNGM1cDgiLCJldmVudF9pZCI6IjBkZDI4ZmRmLTM0YTgtNDhlZC04YzY3LTQzZjM4ZGZjNGYwZCIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNjc3NDU0Mjc4LCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAuYXAtc291dGhlYXN0LTEuYW1hem9uYXdzLmNvbVwvYXAtc291dGhlYXN0LTFfU2x5M01XVkM4IiwiY29nbml0bzp1c2VybmFtZSI6ImRrIiwiZXhwIjoxNjc3NDU3ODc4LCJpYXQiOjE2Nzc0NTQyNzgsImp0aSI6ImY2NTYxN2NlLTkwNmEtNDA5YS05OGZlLTgzYjBmMmYxYWExNCJ9.D2SfSVPq2uPnpfHPVLpw0-jZs9Wjgiazwd06jkQV76r9TFWJeyE2crbEnMUtUIiD4NZ64KPVA4NraNsqCH7QmGmWP-iR3wo7FKKHG4x0q82c7fXyj5QhRHSkLG-MzkynWhuQR327ENNWxA-4HlO9tQhAMuFgRV6NJlXchKk5t1IbCz3a_kdCWAYlyxLILHHt0kc0WjwQ1aTYkCvlSbbud9u6mLJulsHh7EdjkNljywrretu7XMOTEbPiOHSRcvBS6UiXYxwz_kWzd6XbDqN9UAAITLpzaVhH_6L4z7Qd89I1BHrbfJ9yztRiVTwKazgSGmi5qLuQnM_BSNszAgMwvA"

	token, err := auth.Parse(jwt)
	if err != nil {
		t.Error(err)
	}

	if !token.Valid {
		t.Fail()
	}
}

func TestParse(t *testing.T) {

}
