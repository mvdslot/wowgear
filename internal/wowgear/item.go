package wowgear

type Property struct {
	StatCode string				`json:"statCode,omitempty"`
	Amount int					`json:"amount,omitempty"`
}

type Item struct {
	DisplayName string			`json:"displayName,omitempty"`
	Properties []Property		`json:"properties,omitempty"`
	SlotType string				`json:"slotType,omitempty"`
	IsTwoHand bool				`json:"isTwoHand,omitempty"`
	IsAnyHand bool				`json:"isAnyHand,omitempty"`
	SetId string				`json:"setId,omitempty"`
}
