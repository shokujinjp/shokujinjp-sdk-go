package shokujinjp

import (
	"encoding/csv"
	"net/http"
)

const (
	dayFormat = "2006-01-02"
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

	fixedUrl := "https://raw.githubusercontent.com/shokujinjp/data/master/fixed.csv"
	fixedData, err := readCSVFromUrl(fixedUrl)
	if err != nil {
		return nil, err
	}
	for _, v := range fixedData[1:] {
		d := UnmarshalMenuByStringSlice(v)

		all = append(all, d)
	}

	weeklyUrl := "https://raw.githubusercontent.com/shokujinjp/data/master/weekly.csv"
	weeklyData, err := readCSVFromUrl(weeklyUrl)
	if err != nil {
		return nil, err
	}
	for _, v := range weeklyData[1:] {
		d := UnmarshalMenuByStringSlice(v)

		all = append(all, d)
	}

	limitedUrl := "https://raw.githubusercontent.com/shokujinjp/data/master/limited.csv"
	limitedData, err := readCSVFromUrl(limitedUrl)
	if err != nil {
		return nil, err
	}
	for _, v := range limitedData[1:] {
		d := UnmarshalMenuByStringSlice(v)

		all = append(all, d)
	}

	return all, nil
}
