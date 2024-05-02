package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
	"wowgear/internal/wowgear"

	"gopkg.in/yaml.v2"
)

func main() {

	invFile := flag.String("inv", "", "yaml file containing inventory")
	statsFile := flag.String("stats", "", "yaml file containing stats and their weights")
	otherFile := flag.String("other", "", "optional yaml file containing other items to test")
	hitCap := flag.String("hitcap", "", "optional hitcap override")
	overrides := flag.String("overrides", "", "optional stat overrides")
	debug := flag.Int("debug", 0, "debug level")
	export := flag.String("export", "", "export file")

	flag.Parse()

	if *invFile == "" || *statsFile == "" {
		slog.Error("-inv and -stats must be provided, e.g. wowgear  -inv=warlock_inv.yaml -stats=warlock_stats.yaml")
		os.Exit(1)
	}

	data, err := os.ReadFile(*invFile)

	if err != nil {
		slog.Error("unable to read", "file", *invFile, "error", err.Error())
		os.Exit(1)
	}

	inv := &wowgear.Inventory{}
	err = yaml.Unmarshal(data, &inv)

	if err != nil {
		slog.Error("unable to unmarshal yaml", "error", err.Error())
		os.Exit(1)
	}

	data, err = os.ReadFile(*statsFile)

	if err != nil {
		slog.Error("unable to read", "file", *statsFile, "error", err.Error())
		os.Exit(1)
	}

	if *otherFile != "" {
		otherItems := []*wowgear.Item{}
		other, err := os.ReadFile(*otherFile)
		if err != nil {
			slog.Error("unable to read", "file", *otherFile, "error", err.Error())
			os.Exit(1)
		}
		err = yaml.Unmarshal(other, &otherItems)

		if err != nil {
			slog.Error("unable to unmarshal yaml", "error", err.Error())
			os.Exit(1)
		}
		inv.Items = append(inv.Items, otherItems...)
	}

	stats := &wowgear.StatList{}
	err = yaml.Unmarshal(data, &stats)

	if err != nil {
		slog.Error("unable to unmarshal yaml", "error", err.Error())
		os.Exit(1)
	}

	if *hitCap != "" {
		stats.HitCap, err = strconv.Atoi(*hitCap)
		if err != nil {
			slog.Error("error converting to int", "hitcap", hitCap, "error", err.Error())
			os.Exit(1)
		}
	}
	if *overrides != "" {
		overridesSplit := strings.Split(*overrides, "+")
		for _, override := range overridesSplit {
			overrideSplit := strings.Split(override, "=")
			if len(overrideSplit) != 2 {
				slog.Error("overrides must be in format 'stat1=value1+stat2=value2', e.g. int=1.2+rf=2.0")
				os.Exit(1)
			}

			for _, stat := range stats.Stats {
				if stat.Code == overrideSplit[0] {
					stat.Value, err = strconv.ParseFloat(overrideSplit[1], 64)
					if err != nil {
						slog.Error("error converting to float64", "value", overrideSplit[1], "error", err.Error())
						os.Exit(1)
					}
				}
			}
		}
	}

	wowgear.Debug = *debug

	build := wowgear.InitBuild()
	build.StatList = stats

	start := time.Now()

	build.Evaluate(0, inv)

	fmt.Printf("Evaluated %d combinations in %d ms, highest value found is %f\n", wowgear.Combinations, time.Since(start).Milliseconds(), wowgear.HighestValueFound)

	fmt.Println("\nEquipments:")
	for _, eq := range wowgear.BestBuildFound.Equipments {
		if eq.Item != nil {
			fmt.Printf("%s: %s (worth %f)\n", eq.Slot.DisplayName, eq.Item.DisplayName, eq.Item.Value)
		}
	}
	fmt.Println("\nSet bonuses:")
	for _, b := range wowgear.BestBuildFound.SetBonuses {
		stat, err := stats.GetStat(b.Bonus.StatCode)
		if err != nil {
			slog.Error("error getting stat", "stat", b.Bonus.StatCode, "error", err.Error())
		}
		fmt.Printf("%s: %f (worth %f)\n", stat.DisplayName, b.Bonus.Amount, b.Value)
	}
	fmt.Println("\nStats:")
	for _, s := range stats.Stats {
		if s.Value > 0 {
			total := 0.0
			for _, eq := range wowgear.BestBuildFound.Equipments {
				if eq.Item != nil {
					for _, p := range eq.Item.Properties {
						if p.StatCode == s.Code {
							total += p.Amount
						}
					}
				}
			}
			fmt.Printf("Total %s: %f\n", s.DisplayName, total)
		}
	}

	if wowgear.Debug > 0 {
		fmt.Print("\n")
		for _, i := range inv.Items {
			fmt.Printf("%s is worth %f\n", i.DisplayName, i.Value)
		}
	}

	if *export != "" {	
		bytes, err := yaml.Marshal(build)
		if err != nil {
			slog.Error("error marshaling build to yaml", "error", err.Error())
			os.Exit(1)
		}
		err = os.WriteFile(*export, bytes, 0o644)
		if err != nil {
			slog.Error("error writing yaml", "error", err.Error())
			os.Exit(1)
		}
	}
}
