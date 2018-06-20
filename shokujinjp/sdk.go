package shokujinjp

import (
	"encoding/csv"
	"net/http"

	"github.com/shokujinjp/data/gen_weekly/record"
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

func GetAllMenuData() ([]record.Record, error) {
	var all []record.Record

	fixedUrl := "https://raw.githubusercontent.com/shokujinjp/data/master/fixed.csv"
	fixedData, err := readCSVFromUrl(fixedUrl)
	if err != nil {
		return nil, err
	}
	for _, v := range fixedData[1:] {
		d := record.UnmarshalRecordString(v)

		all = append(all, d)
	}

	weeklyUrl := "https://raw.githubusercontent.com/shokujinjp/data/master/weekly.csv"
	weeklyData, err := readCSVFromUrl(weeklyUrl)
	if err != nil {
		return nil, err
	}
	for _, v := range weeklyData[1:] {
		d := record.UnmarshalRecordString(v)

		all = append(all, d)
	}

	return all, nil
}
