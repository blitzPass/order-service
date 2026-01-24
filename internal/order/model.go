package order

type Order struct {
	ID      string
	UserId  string
	EventId string
	Status  string
	Items   []Item
}


type Item struct {
	SectionId string
	Amount    uint8
	Price float32
}

