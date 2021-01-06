package models

import "androidHomeworkApi/pkg/setting"

type TestRecord struct {
	Model

	TestPaperId int `json:"test_paper_id"`
	//TestPaperTitle string `json:"test_paper_title"`
	//TestPaperType string `josn:"test_paper_type"`
	TestPaper TestPaper `json:"test_paper",gorm:"foreignkey:TestPaperId"`

	UserId int `json:"-"`
	CreateAt int `json:"create_at"`
	Score int `json:"score"`
}

func GetTestRecords(userId int,page int) (lists []TestRecord ,count int) {
	offset := pageOffset(page)
	pageSize := setting.PageSize

	db.Preload("TestPaper").Where("user_id=?", userId).Offset(offset).Order("create_at desc").Limit(pageSize).Find(&lists)
	count = len(lists)

	return
}

func TestRecordsTotal(userId int) (count int) {
	db.Model(&TestRecord{}).Where("user_id=?",userId).Count(&count)

	return
}

func AddTestRecord(record TestRecord) TestRecord {
	db.Create(&record)

	return record
}