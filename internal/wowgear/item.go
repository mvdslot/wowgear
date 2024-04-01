package wowgear

type Property struct {
	StatCode string  `yaml:"statCode,omitempty"`
	Amount   float64 `yaml:"amount,omitempty"`
}

type Item struct {
	DisplayName string     `yaml:"displayName,omitempty"`
	Properties  []Property `yaml:"properties,omitempty"`
	SlotType    string     `yaml:"slotType,omitempty"`
	IsTwoHand   bool       `yaml:"isTwoHand,omitempty"`
	IsMainHand  bool       `yaml:"isMainHand,omitempty"`
	IsOffHand   bool       `yaml:"isOffHand,omitempty"`
	SetId       string     `yaml:"setId,omitempty"`
	Value       float64    `yaml:"value,omitempty"`
}
