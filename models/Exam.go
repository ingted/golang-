package models

import (
	"github.com/astaxie/beego/orm"
)

type Exam struct {
	Id int
	Name string
	Time string
}

func NewExam() *Exam  {
	return &Exam{}
}

func (this *Exam)GetExam(gradeid , classzid int) []Exam {
	var Exam  []Exam
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("exam.name,exam.time,exam.remark,,exam.type").
		From("exam").
		Where("gradeid = ?").
		And("classzid = ?").
		OrderBy("time").
		Desc()
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql,gradeid,classzid).QueryRows(&Exam)
	return Exam
}

func (this *Exam)GetAllExam() []Exam  {
	var Exam  []Exam
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.Select("exam.*").
		From("exam").
		OrderBy("time").
		Desc()
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&Exam)
	return Exam
}

func (this *Exam)AddExam(name,time string) error  {
	qb,_ := orm.NewQueryBuilder("mysql")
	qb.InsertInto("exam",name,time).
		Values()
	sql := qb.String()
	o := orm.NewOrm()
	_,error := o.Raw(sql).Exec()
	return error
}

