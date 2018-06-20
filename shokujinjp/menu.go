package shokujinjp

type Menu struct {
	Id          string `csv:"id" json:"id"`
	Name        string `csv:"name" json:"name"`
	Price       string `csv:"price" json:"price"`
	Category    string `csv:"category" json:"category"`
	DayStart    string `csv:"day_start" json":day_start"`
	DayEnd      string `csv:"day_end" json:"day_end"`
	CanWeekday  string `csv:"can_weekday" json:"can_weekday"`
	Description string `csv:"description" json:"description"`
}

func (m *Menu) MarshalString() string {
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
