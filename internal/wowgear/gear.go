package wowgear

import (
	"log/slog"
	"os"
)

type Slot struct {
	Type        string
	DisplayName string
}

type Equipment struct {
	Slot *Slot
	Item *Item
}

type Build struct {
	Equipments []Equipment
	StatList   *StatList
}

type Set struct {
	Id          string					`json:"id,omitempty"`
	DisplayName string					`json:"displayName,omitempty"`
	Bonuses		[]SetBonus				`json:"bonuses,omitempty"`
}

type SetBonus struct {
	Amount int							`json:"amount,omitempty"`
	Bonus  Property						`json:"bonus,omitempty"`
}

var highestValueFound float64

var BestBuildFound *Build

func InitBuild() *Build {
	return &Build{
		Equipments: []Equipment{
			{
				Slot: &Slot{
					Type:        "head",
					DisplayName: "Head",
				},
			},
			{
				Slot: &Slot{
					Type:        "neck",
					DisplayName: "Neck",
				},
			},
			{
				Slot: &Slot{
					Type:        "shoulder",
					DisplayName: "Shoulder",
				},
			},
			{
				Slot: &Slot{
					Type:        "back",
					DisplayName: "Back",
				},
			},
			{
				Slot: &Slot{
					Type:        "chest",
					DisplayName: "Chest",
				},
			},
			{
				Slot: &Slot{
					Type:        "wrist",
					DisplayName: "Wrist",
				},
			},
			{
				Slot: &Slot{
					Type:        "weapon",
					DisplayName: "Main Hand Weapon",
				},
			},
			{
				Slot: &Slot{
					Type:        "weapon",
					DisplayName: "Off Hand Weapon",
				},
			},
			{
				Slot: &Slot{
					Type:        "ranged",
					DisplayName: "Ranged Weapon",
				},
			},
			{
				Slot: &Slot{
					Type:        "hands",
					DisplayName: "Hands",
				},
			},
			{
				Slot: &Slot{
					Type:        "waist",
					DisplayName: "Waist",
				},
			},
			{
				Slot: &Slot{
					Type:        "legs",
					DisplayName: "Legs",
				},
			},
			{
				Slot: &Slot{
					Type:        "feet",
					DisplayName: "Feet",
				},
			},
			{
				Slot: &Slot{
					Type:        "ring",
					DisplayName: "Ring 1",
				},
			},
			{
				Slot: &Slot{
					Type:        "ring",
					DisplayName: "Ring 2",
				},
			},
			{
				Slot: &Slot{
					Type:        "trinket",
					DisplayName: "Trinket 1",
				},
			},
			{
				Slot: &Slot{
					Type:        "trinket",
					DisplayName: "Trinket 2",
				},
			},
		},
	}
}

func (b *Build) GetValue(sets []*Set) (float64, error) {
	total := 0.0

	for _, eq := range b.Equipments {
		itemValue, err := b.getItemValue(eq.Item)
		if err != nil {
			return 0, err
		}
		total += itemValue
	}
	for _, set := range sets {
		itemsInSet := b.countItemsInSet(set.Id)
		for _, setBonus := range set.Bonuses {
			if itemsInSet >= setBonus.Amount {
				value, err := getStatValue(setBonus.Bonus.StatCode, b.StatList)
				if err != nil {
					return 0, err
				}
				total += float64(setBonus.Bonus.Amount) * value
			}
		}
	}

	return total, nil
}

func (b *Build) countItemsInSet(setId string) int {
	items := 0

	for _, eq := range b.Equipments {
		if eq.Item != nil && eq.Item.SetId == setId {
			items++
		}
	}
	return items
}

func (b *Build) getItemValue(item *Item) (float64, error) {
	total := 0.0
	if item == nil {
		return total, nil
	}
	for _, p := range item.Properties {
		val, err := getStatValue(p.StatCode, b.StatList)
		if err != nil {
			return 0, err
		}

		total += float64(p.Amount) * val
	}
	return total, nil
}

func (b *Build) Evaluate(fromEquip int, inv *Inventory) {
	if fromEquip == len(b.Equipments) {
		value, err := b.GetValue(inv.Sets)
		if err != nil {
			slog.Error(err.Error())
			return
		}

		if value > highestValueFound {
			highestValueFound = value
			BestBuildFound = &Build{
				Equipments: []Equipment{},
			}

			for _, eq := range b.Equipments{
				BestBuildFound.Equipments = append(BestBuildFound.Equipments, Equipment{
					Slot: eq.Slot,
					Item: eq.Item,
				} )
			}
		}
		return
	}

	slotType := b.Equipments[fromEquip].Slot.Type

	items := inv.GetItemsForSlotType(slotType)
	if len(items) == 0 {
		b.Evaluate(fromEquip+1, inv)
	}

	if shouldEvaluateAll(items) {
		for _, item := range items {
			b.Equipments[fromEquip].Item = item
			next := fromEquip+1

			// TODO: Main hand / off hand
			if item.IsTwoHand{
				next++
			}

			b.Evaluate(next, inv)
		}
	} else {
		var bestInSlotValue float64
		var bestInslotItem *Item
		for _, item := range items {
			itemValue, err := b.getItemValue(item)
			if err != nil {
				slog.Error(err.Error())
				os.Exit(1)
			}
			if itemValue > bestInSlotValue {
				bestInSlotValue = itemValue
				bestInslotItem = item
			}
		}
		b.Equipments[fromEquip].Item = bestInslotItem
		next := fromEquip+1

		b.Evaluate(next, inv)
	}
}

func shouldEvaluateAll(items []*Item) bool {
	for _, item := range items {
		if item.SetId != "" || item.IsTwoHand {
			return true
		}
	}
	return false
}