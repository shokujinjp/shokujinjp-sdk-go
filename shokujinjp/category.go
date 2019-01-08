package shokujinjp

type CategoryType int

const (
	NoCategory CategoryType = iota
	SetMeal
	Rice
	Soup
)

var (
	categorys = []string{
		"不明",
		"定食",
		"ご飯類",
		"スープ類",
	}
)

func (c CategoryType) String() string {
	switch c {
	case SetMeal:
		return categorys[SetMeal]
	case Rice:
		return categorys[Rice]
	case Soup:
		return categorys[Soup]
	}

	return categorys[NoCategory]
}

func ToCategory(s string) int {
	switch s {
	case "定食":
		return int(SetMeal)
	case "ご飯類":
		return int(Rice)
	case "スープ類":
		return int(Soup)
	}

	return -1
}