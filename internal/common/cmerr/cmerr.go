package cmerr

type Typ int

const (
	TypUnexpected Typ = iota
	TypUnAuthorization
	TypIncorrectInput
	TypDomainError
	TypRateLimited
)

type Error struct {
	error string
	slug  string
	typ   Typ
}

func (s Error) Error() string {
	return s.error
}

func (s Error) Slug() string {
	return s.slug
}

func (s Error) Type() Typ {
	return s.typ
}

func New(err interface{}, slug string, typ Typ) Error {

	if err, ok := err.(error); ok {
		return Error{
			error: err.Error(),
			slug:  slug,
			typ:   typ,
		}

	}

	if err, ok := err.(string); ok {
		return Error{
			error: err,
			slug:  slug,
			typ:   typ,
		}
	}
	return Error{
		error: "error message must be string or error type",
		slug:  InternalServerError,
		typ:   TypUnexpected,
	}

}

func NewUnexpectedError(error string, slug string) Error {
	return Error{
		error: error,
		slug:  slug,
		typ:   TypUnexpected,
	}
}

func NewAuthorizationError(error string, slug string) Error {
	return Error{
		error: error,
		slug:  slug,
		typ:   TypUnAuthorization,
	}
}

func NewIncorrectInputError(error string, slug string) Error {
	return Error{
		error: error,
		slug:  slug,
		typ:   TypIncorrectInput,
	}
}
