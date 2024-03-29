package wowgear

type Slot struct {
	Type string
	DisplayName string
}

type Equipment struct {
	Slot *Slot
	Item *Item
}

type Build struct {
	Equipments []Equipment
	StatList *StatList
}

type Set struct {
	Id string
	DisplayName string
}

type SetBonus struct {
	SetId string
	Amount int
	Bonus Property
}

var BuildInstance = Build{
		Equipments: []Equipment{
		{
			Slot: &Slot{
				Type: "head",
				DisplayName: "Head",
			},
		},
		{
			Slot: &Slot{
				Type: "neck",
				DisplayName: "Neck",
			},
		},
		{
			Slot: &Slot{
				Type: "shoulder",
				DisplayName: "Shoulders",
			},
		},
		{
			Slot: &Slot{
				Type: "back",
				DisplayName: "Back",
			},
		},
		{
			Slot: &Slot{
				Type: "chest",
				DisplayName: "Chest",
			},
		},
		{
			Slot: &Slot{
				Type: "bracer",
				DisplayName: "Bracers",
			},
		},
		{
			Slot: &Slot{
				Type: "weapon_main",
				DisplayName: "Main Hand Weapon",
			},
		},
		{
			Slot: &Slot{
				Type: "weapon_off",
				DisplayName: "Off Hand Weapon",
			},
		},
		{
			Slot: &Slot{
				Type: "ranged",
				DisplayName: "Ranged Weapon",
			},
		},
		{
			Slot: &Slot{
				Type: "hand",
				DisplayName: "Hands",
			},
		},
		{
			Slot: &Slot{
				Type: "belt",
				DisplayName: "Belt",
			},
		},
		{
			Slot: &Slot{
				Type: "leg",
				DisplayName: "Legs",
			},
		},
		{
			Slot: &Slot{
				Type: "boot",
				DisplayName: "Boots",
			},
		},
		{
			Slot: &Slot{
				Type: "ring",
				DisplayName: "Ring 1",
			},
		},
		{
			Slot: &Slot{
				Type: "ring",
				DisplayName: "Ring 2",
			},
		},
		{
			Slot: &Slot{
				Type: "trinket",
				DisplayName: "Trinket 1",
			},
		},
		{
			Slot: &Slot{
				Type: "trinket",
				DisplayName: "Trinket 2",
			},
		},
	},
}

func (b *Build) GetValue() (float64, error) {
	total := 0.0

	for _, eq := range b.Equipments {
		for _, p := range eq.Item.Properties {
			val, err := getStatValue(p.StatCode, b.StatList)
			if err != nil {
				return 0, err
			}

			total += float64(p.Amount) * val
		}
	}
	// TODO: Add set bonuses


	return total, nil
}

