package main

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/tribefintech/microservices/internal/common/server/httperr"
)

type httpServer struct {
	ap *cognito
}

func newHTTPServer(auth *cognito) *httpServer {
	return &httpServer{ap: auth}
}

func (h *httpServer) GetAuthHealthInformation(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, nil)

}

func (h *httpServer) Login(w http.ResponseWriter, r *http.Request) {
	req := LoginRequestBody{}
	if err := render.Decode(r, &req); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	res, err := h.ap.Login(req.Email, req.Password)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.Respond(w, r, res)
}

func (h *httpServer) SignUp(w http.ResponseWriter, r *http.Request) {
	req := SignUpRequestBody{}
	if err := render.Decode(r, &req); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	res, err := h.ap.SignUp(req.Email, req.Password, req.FirstName, req.LastName)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.Respond(w, r, res)
}

func (h *httpServer) ConfirmSignUp(w http.ResponseWriter, r *http.Request) {
	req := ConfirmSignUpRequestBody{}
	if err := render.Decode(r, &req); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	err := h.ap.ConfirmSignUp(req.Email, req.Code)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.NoContent(w, r)
}

func (h *httpServer) ResendCode(w http.ResponseWriter, r *http.Request) {
	req := ResendCodeRequestBody{}
	if err := render.Decode(r, &req); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	res, err := h.ap.ResendConfirmationCode(req.Email)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.Respond(w, r, res)
}

func (h *httpServer) ChangePassword(w http.ResponseWriter, r *http.Request) {
	req := ChangePasswordRequestBody{}
	if err := render.Decode(r, &req); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	err := h.ap.ChangePassword(req.AccessToken, req.NewPassword, req.OldPassword)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.NoContent(w, r)

}

func (h *httpServer) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	req := ForgotPasswordRequestBody{}
	if err := render.Decode(r, &req); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	res, err := h.ap.ForgotPassword(req.Email)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.Respond(w, r, res)
}

func (h *httpServer) ConfirmForgotPassword(w http.ResponseWriter, r *http.Request) {
	req := ForgotPasswordRequestBody{}
	if err := render.Decode(r, &req); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	res, err := h.ap.ForgotPassword(req.Email)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.Respond(w, r, res)

}

func (h *httpServer) RefreshToken(w http.ResponseWriter, r *http.Request) {
	req := RefreshTokenRequestBody{}
	if err := render.Decode(r, &req); err != nil {
		httperr.BadRequest("invalid-request", err, w, r)
		return
	}
	res, err := h.ap.RefreshToken(req.RefreshToken)
	if err != nil {
		httperr.RespondWithSlugError(err, w, r)
		return
	}

	render.Respond(w, r, res)
}
