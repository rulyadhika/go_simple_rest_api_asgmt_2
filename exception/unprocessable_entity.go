package exception

type UnprocessableEntityError struct {
	msg string
}

func NewUnprocessableEntityError(err string) error {
	return &UnprocessableEntityError{msg: err}
}

func (n *UnprocessableEntityError) Error() string {
	return n.msg
}
