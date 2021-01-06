package models

import "androidHomeworkApi/pkg/setting"

type Favorite struct {
	Model

	UserId int `json:"-"`
	ArticleId int 	`json:"article_id"`

	Article ArticleResume `json:"article",gorm:"foreignkey:ArticleId"`

	CreateAt int `json:"create_at"`
}



func GetUserFavorite(id int,page int) (lists []Favorite ,count int) {
	offset := pageOffset(page)
	pageSize := setting.PageSize

	db.Preload("Article","").Where("user_id=?", id).Offset(offset).Order("create_at desc").Limit(pageSize).Find(&lists)
	count = len(lists)

	return
}

func GetUserFavoriteTotal(id int) (count int) {
	db.Model(Favorite{}).Where("user_id=?",id).Count(&count)

	return
}

func AddFavorite(favorite Favorite) Favorite {
	db.Create(&favorite)

	return favorite
}

func DeleteFavorite(id int) {
	db.Where("id=?",id).Delete(Favorite{})

	return
}

func GetFavoriteById(id int) (favorite Favorite,isExist bool) {
	db.First(&favorite,id)

	if favorite.ID<1 {
		return favorite,false
	}

	return favorite,true
}

func IsFavorite(userId,articleId int) (bool,int) {
	favorite:= Favorite{}
	db.Where("user_id=?",userId).Where("article_id=?",articleId).First(&favorite)
	if favorite.ID<1 {
		return false,0
	}

	return true,favorite.ID
}