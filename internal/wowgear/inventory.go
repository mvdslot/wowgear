package wowgear


type Inventory struct {
	Items []*Item				`json:"items,omitempty"`
}

func (i *Inventory) GetItemsForSlotType(slotType string) []*Item {
	result := []*Item{}

	for _, item := range i.Items {
		if item.SlotType == slotType {
			result = append(result, item)
		}
	}
	return result
}