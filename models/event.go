package models

import "time"

type Event struct {
	Model `gorm:"embedded"`

	Title string `json:"title"`

	MemberID int    `json:"member_id"`
	Member   Member `json:"member" gorm:"foreignKey:MemberID"`

	Type      int       `json:"type"`
	Status    int       `json:"status"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Content   string    `json:"content"`
	Views     int       `json:"views"`
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

var MsgFields = map[int]string{
	statusDoing:    "事件进行中",
	statusFinished: "事件已完成",
	statusTimeOut:  "事件已超时",
	statusDelay:    "事件已延迟",

	typeBirthday:   "生日",
	typeWorking:    "工作",
	typeAssignment: "任务",
	typeMeeting:    "会议",
}

// GetMsg 获取备忘事件枚举字段的函数
func GetMsg(code int) string {
	msg, ok := MsgFields[code]
	if ok {
		return msg
	}

	return "Not Exist Event Type Or Status"
}

func ExistEventByID(id int) bool {
	var event Event
	db.Select("id").Where("id = ?", id).First(&event)

	if event.ID > 0 {
		return true
	}

	return false

}

func GetEvent(maps interface{}) (id int) {
	db.Model(&Event{}).Where("id = ?", id)

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

func EditEvent(id int, data interface{}) bool {
	db.Model(&Event{}).Where("id = ?", id).Updates(data)

	return true
}

func AddEvent(data map[string]interface{}) bool {
	db.Create(&Event{
		Title:     data["title"].(string),
		Member:    data["member"].(Member),
		Type:      data["type"].(int),
		Status:    data["status"].(int),
		StartTime: data["start_time"].(time.Time),
		EndTime:   data["end_time"].(time.Time),
		Content:   data["content"].(string),
		Views:     data["views"].(int),
	})

	return true
}

func DeleteEvent(id int) bool {
	db.Where("id = ?", id).Delete(Event{})

	return true
}
