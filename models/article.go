package models

import "androidHomeworkApi/pkg/setting"

type Article struct {
	Model

	Title string `json:"title"`
	Content string `json:"content"`
	Cover string `json:"cover"`
	CreateTime string `json:"createTime"`
	CommentCount int `json:"commentCount"`

	IsFavotite bool `json:"is_favorite",gorm:"-"`
	Author string `json:"author"`
}

type ArticleResume struct {
	Model

	Title string `json:"title"`
	Cover string `json:"cover"`
	CreateTime string `json:"createTime"`
	CommentCount int `json:"commentCount"`
}

func (a ArticleResume) TableName() string {
	return "article"
}

func GetArticles(query interface{},page int) (lists []Article, count int) {
	offset := pageOffset(page)
	pageSize := setting.PageSize

	db.Where(query).Select("id,title,cover,create_time,comment_count,author").Offset(offset).Order("create_time desc").Limit(pageSize).Find(&lists)
	count=len(lists)

	return
}

func GetArticle(query interface{}) (article Article) {
	db.Where(query).Find(&article)

	return
}

func ArticleTotal(query interface{}) (count int) {
	db.Model(&Article{}).Where(query).Count(&count)

	return
}