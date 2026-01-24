package order

type Service interface {
	Create(dto CreateDto, userId string) *Order
}


type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo }
}

func(s *service) Create(dto CreateDto, userId string) *Order {

	return nil
}


