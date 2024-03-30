package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"
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

	start := time.Now()

	build.Evaluate(0, inv)

	fmt.Printf("Evaluated %d combinations in %d ms, highest value found is %f\n", wowgear.Combinations, time.Since(start).Milliseconds(), wowgear.HighestValueFound)

	for _, eq := range wowgear.BestBuildFound.Equipments {
		if eq.Item != nil {
			fmt.Printf("%s: %s (worth %f)\n", eq.Slot.DisplayName, eq.Item.DisplayName, eq.Item.Value)
		}
	}
	fmt.Print("\n")
	for _, b := range wowgear.BestBuildFound.SetBonuses {
		fmt.Printf("%s: %d (worth %f)\n", b.Bonus.StatCode, b.Bonus.Amount, b.Value)
	}
	fmt.Print("\n")
	for _, i := range inv.Items {
		fmt.Printf("%s is worth %f\n", i.DisplayName, i.Value)
	}

}



