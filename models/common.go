package models

import (
	"androidHomeworkApi/pkg/setting"
)

func pageOffset(page int) int {
	offset := 0

	if page > 0 {
		offset = (page - 1) * setting.PageSize
	}

	return offset
}