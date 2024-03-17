package exception

type BadRequestError struct {
	msg string
}

func NewBadRequestError(err string) error {
	return &BadRequestError{msg: err}
}

func (b *BadRequestError) Error() string {
	return b.msg
}
