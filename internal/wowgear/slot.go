package wowgear

type Slot struct {
	Id string
	DisplayName string
	Amount int
}

type Equipment struct {
	Slot *Slot
	Item *Item
}

type Build struct {
	Equipments []Equipment
}

var BuildInstance = Build{
		Equipments: []Equipment{
		{
			Slot: &Slot{
				Id: "head",
				DisplayName: "Head",
			},
		},
		{
			Slot: &Slot{
				Id: "neck",
				DisplayName: "Neck",
			},
		},
		{
			Slot: &Slot{
				Id: "shoulder",
				DisplayName: "Shoulders",
			},
		},
		{
			Slot: &Slot{
				Id: "back",
				DisplayName: "Back",
			},
		},
		{
			Slot: &Slot{
				Id: "chest",
				DisplayName: "Chest",
			},
		},
		{
			Slot: &Slot{
				Id: "bracer",
				DisplayName: "Bracers",
			},
		},
		{
			Slot: &Slot{
				Id: "weapon_1",
				DisplayName: "Main Hand Weapon",
			},
		},
		{
			Slot: &Slot{
				Id: "weapon_2",
				DisplayName: "Off Hand Weapon",
			},
		},
		{
			Slot: &Slot{
				Id: "ranged",
				DisplayName: "Ranged Weapon",
			},
		},
		{
			Slot: &Slot{
				Id: "hand",
				DisplayName: "Hands",
			},
		},
		{
			Slot: &Slot{
				Id: "belt",
				DisplayName: "Belt",
			},
		},
		{
			Slot: &Slot{
				Id: "leg",
				DisplayName: "Legs",
			},
		},
		{
			Slot: &Slot{
				Id: "boot",
				DisplayName: "Boots",
			},
		},
		{
			Slot: &Slot{
				Id: "ring_1",
				DisplayName: "Ring 1",
			},
		},
		{
			Slot: &Slot{
				Id: "ring_2",
				DisplayName: "Ring 2",
			},
		},
		{
			Slot: &Slot{
				Id: "trinket_1",
				DisplayName: "Trinket 1",
			},
		},
		{
			Slot: &Slot{
				Id: "trinket_2",
				DisplayName: "Trinket 2",
			},
		},
	},
}

func (b *Build) GetValue() (float64, error) {
	total := 0.0

	for _, eq := range b.Equipments {
		for _, p := range eq.Item.Properties {
			val, err := getStatValue(p.StatId)
			if err != nil {
				return 0, err
			}

			total += float64(p.Amount) * val
		}
	}

	return total, nil
}

