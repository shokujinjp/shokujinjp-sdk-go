package shokujinjp

import (
	"sort"
)

func SortByCategory(menus []Menu) []Menu {
	sort.Slice(menus, func(i, j int) bool {
		return ToCategory(menus[i].Category) < ToCategory(menus[j].Category)
	})

	return menus
}
