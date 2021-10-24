package service

type ErrRecordNotFound string

func (e ErrRecordNotFound) Error() string {
	return string(e)
}
