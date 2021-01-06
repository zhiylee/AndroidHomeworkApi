package models

import "androidHomeworkApi/pkg/setting"

type TestPaper struct {
	Model

	Title string `json:"title"`
	Type string `json:"type"`
	CreateTime int `json:"create_time"`
}

func GetTestPapers(testType string,page int) (lists []TestPaper ,count int) {
	offset := pageOffset(page)
	pageSize := setting.PageSize

	db.Where("type=?", testType).Offset(offset).Limit(pageSize).Find(&lists)
	count = len(lists)

	return
}

func GetTestPaperByType(testType string) (testPaper TestPaper,isExist bool) {
	db.Where("type=?",testType).Order("id desc").First(&testPaper)

	if testPaper.ID<1 {
		return testPaper,false
	}

	return testPaper,true
}

func GetTestPaperById(id int) (testPaper TestPaper,isExist bool) {
	db.First(&testPaper,id)

	if testPaper.ID<1 {
		return testPaper,false
	}

	return testPaper,true
}

func GetTestPapersTotal(testType string) (count int) {
	db.Model(TestPaper{}).Where("type=?",testType).Count(&count)

	return
}