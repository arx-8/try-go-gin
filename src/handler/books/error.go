package books

type APIError string

func (e APIError) Error() string {
	return string(e)
}
