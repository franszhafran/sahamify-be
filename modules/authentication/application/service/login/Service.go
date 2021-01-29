package login

// Service is login service struct
type Service struct{}

// Execute is to run this service
func (l *Service) Execute(r *Request) interface{} {
	r.password = "tes"
	return true
}
