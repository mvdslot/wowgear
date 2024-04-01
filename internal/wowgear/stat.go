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

func (sl *StatList) GetStat(statCode string) (*Stat, error) {
	for _, stat := range sl.Stats {
		if stat.Code == statCode {
			return stat, nil
		}
	}
	return nil, errors.New("unknown stat code " + statCode)
}
