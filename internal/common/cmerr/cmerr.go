package cmerr

type Typ int

const (
	TypUnexpected Typ = iota
	TypAuthorization
	TypIncorrectInput
	TypDomainError
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

func (s Error) ErrorType() Typ {
	return s.typ
}

func New(error string, slug string, typ Typ) Error {
	return Error{
		error: error,
		slug:  slug,
		typ:   typ,
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
		typ:   TypAuthorization,
	}
}

func NewIncorrectInputError(error string, slug string) Error {
	return Error{
		error: error,
		slug:  slug,
		typ:   TypIncorrectInput,
	}
}
