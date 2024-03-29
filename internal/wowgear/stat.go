package wowgear

import "errors"

type Stat struct {
	Id          string
	DisplayName string
	Value       float64
}

var allStats []Stat

func getStatValue(statId string) (float64, error) {
	for _, stat := range allStats {
		if stat.Id == statId {
			return stat.Value, nil
		}
	}
	return 0, errors.New("unknown stat id " + statId)
}