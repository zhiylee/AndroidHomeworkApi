package models

import "encoding/json"

type Question struct {
	Model

	Title string `json:"title"`
	Options interface{} `json:"options"`
	Answer string `json:"answer"`
	TestPaperId int `json:"-"`
}

func GetQuestions(testPaperId int) (lists []Question ,count int) {
	db.Where("test_paper_id=?", testPaperId).Find(&lists)
	count = len(lists)

	return
}

func AddQuestion(question Question)  {
	db.Create(&question)
}

func (q *Question) BeforeSave()  {
	q.Options,_ = json.Marshal(q.Options)
}

func (q *Question) AfterFind()  {
	options:=make(map[string]string)
	json.Unmarshal( []byte( string( q.Options.([]uint8) ) ),&options)

	q.Options = options
}