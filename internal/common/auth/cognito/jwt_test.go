package cognito

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {

}

func Test_getJWT(t *testing.T) {
	if !(os.Getenv("AWS_COGNITO_USER_POOL_ID") != "" && os.Getenv("AWS_COGNITO_REGION") != "") {
		t.Skip("requires AWS Cognito environment variables")
	}

	auth := New(os.Getenv("AWS_COGNITO_REGION"), os.Getenv("AWS_COGNITO_USER_POOL_ID"))

	err := auth.getJWK()
	if err != nil {
		t.Error(err)
	}

	jwt := "eyJraWQiOiJlS3lvdytnb1wvXC9yWmtkbGFhRFNOM25jTTREd0xTdFhibks4TTB5b211aE09IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiJjMTcxOGY3OC00ODY5LTRmMmEtYTk2ZS1lYmEwYmJkY2RkMjEiLCJldmVudF9pZCI6IjZmYWMyZGNjLTJlMzUtMTFlOS05NDZjLTZiZDI0YmRlZjFiNiIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE1NDk5MTQyNjUsImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC51cy13ZXN0LTIuYW1hem9uYXdzLmNvbVwvdXMtd2VzdC0yX0wwVldGSEVueSIsImV4cCI6MTU0OTkxNzg2NSwiaWF0IjoxNTQ5OTE0MjY1LCJqdGkiOiIzMTg0MDdkMC0zZDNhLTQ0NDItOTMyYy1lY2I0MjQ2MzRiYjIiLCJjbGllbnRfaWQiOiI2ZjFzcGI2MzZwdG4wNzRvbjBwZGpnbms4bCIsInVzZXJuYW1lIjoiYzE3MThmNzgtNDg2OS00ZjJhLWE5NmUtZWJhMGJiZGNkZDIxIn0.rJl9mdCrw_lertWhC5RiJcfhRP-xwTYkPLPXmi_NQEO-LtIJ-kwVEvUaZsPnBXku3bWBM3V35jdJloiXclbffl4SDLVkkvU9vzXDETAMaZEzOY1gDVcg4YzNNR4H5kHnl-G-XiN5MajgaWbjohDHTvbPnqgW7e_4qNVXueZv2qfQ8hZ_VcyniNxMGaui-C0_YuR6jdH-T14Wl59Cyf-UFEyli1NZFlmpUQ8QODGMUI12PVFOZiHJIOZ3CQM_Xs-TlRy53RlKGFzf6RQfRm57rJw_zLyJHHnB8DZgbdCRfhNsqZka7ZZUUAlS9aMzdmSc3pPFSJ-hH3p8eFAgB4E71g"

	token, err := auth.Parse(jwt)
	if err != nil {
		t.Error(err)
	}

	if !token.Valid {
		t.Fail()
	}
}
