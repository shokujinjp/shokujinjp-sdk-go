package shokujinjp

type CategoryType int

const (
	NoCategory CategoryType = iota
	Limited
	SetMeal
	Rice
	Soup
)

var (
	categorys = []string{
		"不明",
		"期間限定",
		"定食",
		"ご飯類",
		"スープ類",
	}
)

func (c CategoryType) String() string {
	switch c {
	case Limited:
		return categorys[Limited]
	case SetMeal:
		return categorys[SetMeal]
	case Rice:
		return categorys[Rice]
	case Soup:
		return categorys[Soup]
	}

	return categorys[NoCategory]
}

func Category(s string) CategoryType {
	switch s {
	case "期間限定":
		return Limited
	case "定食":
		return SetMeal
	case "ご飯類":
		return Rice
	case "スープ類":
		return Soup
	}

	return -1
}
