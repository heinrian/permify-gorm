package helpers

import (
	"math"
)

// NextPageCal calculate next page according to the page number and total page count.
// Param: int
// Param: int
// return int
func NextPageCal(page int, totalPage int) int {
	if page == totalPage {
		return page
	}
	return page + 1
}

// PrevPageCal calculate previous page according to page number.
// Param: int
// return int
func PrevPageCal(page int) int {
	if page > 1 {
		return page - 1
	}
	return page
}

// TotalPage calculate total page according to records count and limit.
// Param: int64
// Param: int
// return int
func TotalPage(count int64, limit int) int {
	return int(math.Ceil(float64(count) / float64(limit)))
}

// OffsetCal calculate offset according to page and limit.
// Param: int
// Param: int
// return int
func OffsetCal(page int, limit int) int {
	return (page - 1) * limit
}
