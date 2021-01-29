package login

// Request is a struct for parameter of LoginService
type Request struct {
	username string
	password string
}

// NewRequest create new instance of LoginRequest
func NewRequest(u string, p string) *Request {
	return &Request{
		username: u,
		password: p,
	}
}
