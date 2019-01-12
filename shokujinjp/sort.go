package shokujinjp

import (
	"sort"
)

func SortByCategory(menus []Menu) []Menu {
	sort.Slice(menus, func(i, j int) bool {
		return ToCategory(menus[i].Category) < ToCategory(menus[j].Category)
	})

	sort.Slice(menus, func(i, j int) bool {
		return int(menus[i].Id) < int(menus[j].Id)
	})

	return menus
}
