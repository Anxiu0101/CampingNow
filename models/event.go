package models

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	Model `gorm:"embedded"`

	Title string `json:"title"`

	MemberID int    `json:"member_id"`
	Member   Member `json:"member" gorm:"foreignKey:MemberID"`

	Type      int    `json:"type"`
	Status    int    `json:"status"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Content   string `json:"content"`
	Views     int    `json:"views"`
}

// 使用枚举定义备忘事件的类型和状态
const (
	statusDoing    = 101
	statusFinished = 102
	statusTimeOut  = 103
	statusDelay    = 104

	typeBirthday   = 201
	typeWorking    = 202
	typeAssignment = 203
	typeMeeting    = 204
)

var MsgFields = map[string]int{
	"事件进行中": statusDoing,
	"事件已完成": statusFinished,
	"事件已超时": statusTimeOut,
	"事件已延迟": statusDelay,

	"生日": typeBirthday,
	"工作": typeWorking,
	"任务": typeAssignment,
	"会议": typeMeeting,
}

// GetCode 获取备忘事件枚举字段的函数
func GetCode(msg string) int {
	code, ok := MsgFields[msg]
	if ok {
		return code
	}

	return -1
}

func ExistEventByID(id int) bool {
	var event Event
	db.Select("id").Where("id = ?", id).First(&event)

	if event.ID > 0 {
		return true
	}

	return false

}

func GetEvent(id int) (event Event) {
	db.Preload("Member").Where("id = ?", id).First(&event)

	return
}

func GetEvents(pageNum int, pageSize int, maps interface{}) (event []Event) {
	db.Preload("Member").Where(maps).Offset(pageNum).Limit(pageSize).Find(&event)

	return
}

func GetEventTotal(maps interface{}) (count int64) {
	db.Model(&Event{}).Where(maps).Count(&count)

	return
}

func GetEventByStatus(status int, pageNum int, pageSize int) (event []Event) {
	db.Model(Event{}).Where("status = ?", status).Offset(pageNum).Limit(pageSize).Find(&event)

	return
}

func GetEventByType(types int, pageNum int, pageSize int) (event []Event) {
	db.Model(Event{}).Where("status = ?", types).Offset(pageNum).Limit(pageSize).Find(&event)

	return
}

func AddEventViews(id int) {
	db.Model(Event{}).Where("id = ? ", id).Update("views", gorm.Expr("views+ ?", 1))

	return
}

func EditEvent(id int, data interface{}) bool {
	db.Model(&Event{}).Where("id = ?", id).Updates(data)

	return true
}

func AddEvent(data map[string]interface{}) bool {

	db.Create(&Event{
		Title:     data["title"].(string),
		MemberID:  data["member_id"].(int),
		Type:      GetCode(data["type"].(string)),
		Status:    GetCode(data["status"].(string)),
		StartTime: time.Now().Unix(),
		EndTime:   data["end_time"].(int64),
		Content:   data["content"].(string),
		Views:     0,
	})

	return true
}

func DeleteEvent(id int) bool {
	db.Where("id = ?", id).Delete(&Event{})

	return true
}
