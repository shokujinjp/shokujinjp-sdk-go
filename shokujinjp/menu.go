package shokujinjp

import (
	"fmt"
	"time"
)

type Menu struct {
	Id          string `csv:"id" json:"id"`
	Name        string `csv:"name" json:"name"`
	Price       string `csv:"price" json:"price"`
	Category    string `csv:"category" json:"category"`
	DayStart    string `csv:"day_start" json:"day_start"`
	DayEnd      string `csv:"day_end" json:"day_end"`
	CanWeekday  string `csv:"can_weekday" json:"can_weekday"`
	Description string `csv:"description" json:"description"`
}

func (m *Menu) String() string {
	return m.Id + "," +
		m.Name + "," +
		m.Price + "," +
		m.Category + "," +
		m.DayStart + "," +
		m.DayEnd + "," +
		m.CanWeekday + "," +
		m.Description
}

func (m *Menu) MarshalStringSlice() []string {
	return []string{
		m.Id,
		m.Name,
		m.Price,
		m.Category,
		m.DayStart,
		m.DayEnd,
		m.CanWeekday,
		m.Description}
}

func UnmarshalMenuByStringSlice(menuStr []string) Menu {
	d := Menu{
		Id:          menuStr[0],
		Name:        menuStr[1],
		Price:       menuStr[2],
		Category:    menuStr[3],
		DayStart:    menuStr[4],
		DayEnd:      menuStr[5],
		CanWeekday:  menuStr[6],
		Description: menuStr[7],
	}

	return d
}

func (m Menu) toDisplayName() Menu {
	result := m

	if Category(m.Category) == SetMeal {
		result.Name = fmt.Sprintf("%s%s", m.Name, SetMeal.String())
	}

	return result
}

func checkCanOrder(m Menu, date time.Time) (bool, error) {
	// if call this func, already check "m.DayStart is not black"
	if m.DayEnd == "" {
		return true, nil
	}
	dayStart, err := time.Parse(DayFormat, m.DayStart)
	if err != nil {
		return false, err
	}

	dayEnd, err := time.Parse(DayFormat, m.DayEnd)
	if err != nil {
		return false, err
	}

	if (dayStart.Before(date)) && (dayEnd.After(date)) {
		return true, nil
	}

	return false, nil
}

func GetMenuDateData(t time.Time) ([]Menu, error) {
	var today []Menu
	wdays := [...]string{"日", "月", "火", "水", "木", "金", "土"}
	inputWdays := wdays[t.Weekday()]

	all, err := GetMenuAllData()
	if err != nil {
		return nil, err
	}

	for _, m := range all {
		m = m.toDisplayName()

		if m.DayStart != "" {
			b, err := checkCanOrder(m, t)
			if err != nil {
				return nil, err
			}
			if b == true {
				today = append(today, m)
				continue
			}

			// if can't order day, not append
			continue
		}

		if m.CanWeekday == "" {
			// if blank, can order all wdays
			today = append(today, m)
			continue
		}

		if m.CanWeekday == inputWdays {
			today = append(today, m)
			continue
		}

	}

	return today, nil
}
