package exception

type NotFoundError struct {
	msg string
}

func NewNotFoundError(err string) error {
	return &NotFoundError{msg: err}
}

func (n *NotFoundError) Error() string {
	return n.msg
}
