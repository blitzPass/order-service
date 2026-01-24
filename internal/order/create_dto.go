package order

type CreateDto struct {
	EventId string `json:"event_id" valitate:"required, uuid4"`
	Items []ItemDto `json:"items" validate:"required, min=1, dive"`

	 
}

type ItemDto struct {
	SectionId string `json:"section_id" valite:"required, uuid4"`
	Amount uint8 `json:"amount" validate:"required, gt=0"`
}