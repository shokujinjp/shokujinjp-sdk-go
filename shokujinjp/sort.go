package shokujinjp

import (
	"sort"
	"strconv"
)

func SortByCategory(menus []Menu) []Menu {
	sort.SliceStable(menus, func(i, j int) bool {
		ii, _ := strconv.Atoi(menus[i].Id)
		jj, _ := strconv.Atoi(menus[j].Id)

		return ii < jj
	})

	sort.SliceStable(menus, func(i, j int) bool {
		return ToCategory(menus[i].Category) < ToCategory(menus[j].Category)
	})

	return menus
}
