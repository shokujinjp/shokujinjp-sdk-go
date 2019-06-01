package shokujinjp

import (
	"encoding/csv"
	"net/http"
)

const (
	DayFormat  = "2006-01-02"
	BaseURL    = "https://raw.githubusercontent.com/shokujinjp/data/master"
	FixedURL   = BaseURL + "/fixed.csv"
	WeeklyURL  = BaseURL + "/weekly.csv"
	LimitedURL = BaseURL + "/limited.csv"
)

func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetMenuAllData() ([]Menu, error) {
	var all []Menu

	fixedData, err := readCSVFromUrl(FixedURL)
	if err != nil {
		return nil, err
	}
	for _, v := range fixedData[1:] {
		d := UnmarshalMenuByStringSlice(v)

		all = append(all, d)
	}

	weeklyData, err := readCSVFromUrl(WeeklyURL)
	if err != nil {
		return nil, err
	}
	for _, v := range weeklyData[1:] {
		d := UnmarshalMenuByStringSlice(v)

		all = append(all, d)
	}

	limitedData, err := readCSVFromUrl(LimitedURL)
	if err != nil {
		return nil, err
	}
	for _, v := range limitedData[1:] {
		d := UnmarshalMenuByStringSlice(v)

		all = append(all, d)
	}

	return all, nil
}
