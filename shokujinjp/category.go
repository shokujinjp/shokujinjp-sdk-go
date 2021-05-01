package shokujinjp

import "strings"

type CategoryType int

const (
	NoCategory CategoryType = iota
	Limited
	SetMeal
	Rice
	Soup
	Noodle
	NewMenu
	Drink
	AlaCarte
	Snacks
	Vegetable
	Meat
	SeaFood
	ServiceSet
)

var (
	categories = []string{
		"不明",
		"期間限定",
		"定食",
		"ご飯類",
		"スープ類",
		"麺類",
		"新メニュー",
		"飲み物",
		"おつまみ",
		"点心類",
		"野菜類",
		"牛・豚・鷄・ラム",
		"海鮮類",
		"サービスセット",
	}
)

func (c CategoryType) String() string {
	input := int(c)
	if input < 0 || len(categories) < input {
		return categories[NoCategory]
	}

	return categories[c]
}

func Category(s string) CategoryType {
	for i, category := range categories {
		if strings.EqualFold(s, category) {
			return CategoryType(i)
		}
	}

	return NoCategory
}
