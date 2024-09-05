package handler

type Servicer interface {
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}
