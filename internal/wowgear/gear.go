package wowgear

import "log/slog"

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
	Id          string
	DisplayName string
	Bonuses		[]SetBonus
}

type SetBonus struct {
	Amount int
	Bonus  Property
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
					DisplayName: "Shoulders",
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
					Type:        "bracer",
					DisplayName: "Bracers",
				},
			},
			{
				Slot: &Slot{
					Type:        "weapon_main",
					DisplayName: "Main Hand Weapon",
				},
			},
			{
				Slot: &Slot{
					Type:        "weapon_off",
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
					Type:        "hand",
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
					Type:        "leg",
					DisplayName: "Legs",
				},
			},
			{
				Slot: &Slot{
					Type:        "boot",
					DisplayName: "Boots",
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

func (b *Build) GetValue() (float64, error) {
	total := 0.0

	for _, eq := range b.Equipments {
		if eq.Item != nil {
			for _, p := range eq.Item.Properties {
				val, err := getStatValue(p.StatCode, b.StatList)
				if err != nil {
					return 0, err
				}

				total += float64(p.Amount) * val
			}
		}
	}
	// TODO: Add set bonuses

	return total, nil
}

func (b *Build) Evaluate(fromEquip int, inv *Inventory) {
	if fromEquip == len(b.Equipments) {
		value, err := b.GetValue()
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

	for _, item := range items {
		b.Equipments[fromEquip].Item = item
		next := fromEquip+1
		if item.IsTwoHand{
			next++
		}

		b.Evaluate(next, inv)
	}
}