package main

import (
	"backend/log"
	"encoding/csv"
	"net/http"
)

type Utilities struct {
	logger log.Logger
}

func NewUtilities(l log.Logger) *Utilities {
	return &Utilities{logger:l}
}

func (u *Utilities) readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	//reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *Utilities) AbsInt(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
func (u *Utilities) MinInt(valueA int, valueB int) int {
	if valueA < valueB {
		return valueA
	}
	return valueB
}
