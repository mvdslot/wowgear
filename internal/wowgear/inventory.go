package wowgear

type Inventory struct {
	Items []*Item `json:"items,omitempty"`
	Sets  []*Set  `json:"sets,omitempty"`
}

func (i *Inventory) getItemsForSlotType(slotType string, previous *Item) []*Item {
	result := []*Item{}

	for _, item := range i.Items {
		if item.SlotType == slotType && (previous == nil || previous != item) {
			if item.SlotType == "weapon" && previous.SlotType == "weapon" && item.IsMainHand {
				continue
			}
			if  item.SlotType == "weapon" && previous.SlotType != "weapon" && item.IsOffHand {
				continue
			}

			result = append(result, item)
		}
	}
	return result
}
