package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"wowgear/internal/wowgear"
)

func main() {

	inv := &wowgear.Inventory{}

	data, err := os.ReadFile("./inventory.json")
	
	if err != nil {
		slog.Error("unable to read inventory.json", "error", err.Error())
	}

	err = json.Unmarshal(data, &inv)

	if err != nil {
		slog.Error("unable to unmarshal json", "error", err.Error())
	}

	stats := &wowgear.StatList{}

	data, err = os.ReadFile("./warlock.json")
	
	if err != nil {
		slog.Error("unable to read warlock.json", "error", err.Error())
	}

	err = json.Unmarshal(data, &stats)

	if err != nil {
		slog.Error("unable to unmarshal json", "error", err.Error())
	}

	build := wowgear.InitBuild()
	build.StatList = stats

	build.Evaluate(0, inv)

	for _, eq := range wowgear.BestBuildFound.Equipments {
		if eq.Item != nil {
			fmt.Printf("%s: %s\n", eq.Slot.DisplayName, eq.Item.DisplayName)
		}
	}
}



