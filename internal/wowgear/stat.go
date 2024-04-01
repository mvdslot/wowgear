package wowgear

import "errors"

type Stat struct {
	Code        string  `yaml:"code,omitempty"`
	DisplayName string  `yaml:"displayName,omitempty"`
	Value       float64 `yaml:"value,omitempty"`
}

type StatList struct {
	Stats  []*Stat `yaml:"stats,omitempty"`
	HitCap int     `yaml:"hitCap,omitempty"`
}

func getStatValue(statCode string, statList *StatList) (float64, error) {
	for _, stat := range statList.Stats {
		if stat.Code == statCode {
			return stat.Value, nil
		}
	}
	return 0, errors.New("unknown stat code " + statCode)
}
