package httperr

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/tribefintech/microservices/internal/common/cmerr"
	"github.com/tribefintech/microservices/internal/common/logs"
)

func InternalError(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "internal server error", http.StatusInternalServerError)
}

func Unauthorized(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "unauthorized", http.StatusUnauthorized)
}

func BadRequest(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "bad request", http.StatusBadRequest)
}

func DomainError(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "domain error", http.StatusOK)
}

func RateLimitedError(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "rate limited", http.StatusTooManyRequests)
}

func RespondWithSlugError(err error, w http.ResponseWriter, r *http.Request) {
	slugError, ok := err.(cmerr.Error)
	if !ok {
		InternalError("internal-server-error", err, w, r)
		return
	}

	switch slugError.Type() {
	case cmerr.TypUnAuthorization:
		Unauthorized(slugError.Slug(), slugError, w, r)
	case cmerr.TypIncorrectInput:
		BadRequest(slugError.Slug(), slugError, w, r)
	case cmerr.TypDomainError:
		DomainError(slugError.Slug(), slugError, w, r)
	case cmerr.TypRateLimited:
		RateLimitedError(slugError.Slug(), slugError, w, r)
	default:
		InternalError(slugError.Slug(), slugError, w, r)
	}
}

func httpRespondWithError(err error, slug string, w http.ResponseWriter, r *http.Request, logMSg string, status int) {
	logs.GetLogEntry(r).WithError(err).WithField("error-slug", slug).Warn(logMSg)
	resp := ErrorResponse{slug, status}

	if err := render.Render(w, r, resp); err != nil {
		panic(err)
	}
}

type ErrorResponse struct {
	Slug       string `json:"slug"`
	httpStatus int
}

func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.httpStatus)
	return nil
}
