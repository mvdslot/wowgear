package wowgear

import "errors"

type Stat struct {
	Code        string		`json:"code,omitempty"`
	DisplayName string		`json:"displayName,omitempty"`
	Value       float64		`json:"value,omitempty"`
}

type StatList struct{
	Stats []*Stat 			`json:"stats,omitempty"`
	HitCap int				`json:"hitCap,omitempty"`
} 

func getStatValue(statCode string, statList *StatList) (float64, error) {
	for _, stat := range statList.Stats {
		if stat.Code == statCode {
			return stat.Value, nil
		}
	}
	return 0, errors.New("unknown stat code " + statCode)
}