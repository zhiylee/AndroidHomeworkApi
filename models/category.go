package models

type Category struct {
	Model

	Name string `json:"name"`
}

func GetCategories() (lists []Category, count int) {
	db.Find(&lists)
	count=len(lists)

	return
}