package wowgear

type Property struct {
	StatId string				`json:"statId,omitempty"`
	Amount int					`json:"amount,omitempty"`
}

type Item struct {
	DisplayName string			`json:"displayName,omitempty"`
	Properties []Property		`json:"properties,omitempty"`
	Unique bool					`json:"unique,omitempty"`
	Slotid string				`json:"slotId,omitempty"`
	IsTwoHand bool				`json:"isTwoHand,omitempty"`
}

type Inventory struct {
	Items []*Item				`json:"items,omitempty"`
}