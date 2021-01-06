package models

import "androidHomeworkApi/pkg/setting"

type Comment struct {
	Model

	ArticleId int `json:"article_id"`
	UserId int 	`json:"user_id"`
	Content	 string `json:"content"`
	CreateAt int `json:"create_at"`

	User User `json:"user",gorm:"foreignkey:UserId"`
}

func GetArticleComments(id int,page int) (lists []Comment ,count int) {
	offset := pageOffset(page)
	pageSize := setting.PageSize

	db.Preload("User").Where("article_id=?", id).Order("create_at desc").Offset(offset).Limit(pageSize).Find(&lists)
	count = len(lists)

	return
}

func GetArticleCommentsTotal(id int) (count int) {
	db.Model(Comment{}).Where("article_id=?",id).Count(&count)

	return
}

func AddComment(comment Comment) Comment {
	db.Create(&comment)

	// 评论数+1
	article := Article{}
	db.Find(&article,comment.ArticleId)
	db.Model(&article).Update("comment_count",article.CommentCount+1)

	db.Find(&comment.User,comment.UserId)

	return comment
}