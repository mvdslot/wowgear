package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"
	"wowgear/internal/wowgear"
)

func main() {

	inv := &wowgear.Inventory{}

	invFile := flag.String("inv", "", "json file containing inventory")
	statsFile := flag.String("stats", "", "json file containing stats and their weights")

	flag.Parse()

	data, err := os.ReadFile(*invFile)
	
	if err != nil {
		slog.Error("unable to read", "file", *invFile, "error", err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(data, &inv)

	if err != nil {
		slog.Error("unable to unmarshal json", "error", err.Error())
		os.Exit(1)
	}

	stats := &wowgear.StatList{}

	data, err = os.ReadFile(*statsFile)
	
	if err != nil {
		slog.Error("unable to read", "file", *statsFile, "error", err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(data, &stats)

	if err != nil {
		slog.Error("unable to unmarshal json", "error", err.Error())
		os.Exit(1)
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
	fmt.Print("\nBonuses:\n")
	for _, b := range wowgear.BestBuildFound.SetBonuses {
		fmt.Printf("%s: %d (worth %f)\n", b.Bonus.StatCode, b.Bonus.Amount, b.Value)
	}
	// fmt.Print("\n")
	// for _, i := range inv.Items {
	// 	fmt.Printf("%s is worth %f\n", i.DisplayName, i.Value)
	// }
}



